/*
 * vkmem_shim.c
 *   Host bridge for valkeymem: entry point helpers the Go host calls.
 *   Compiled without any -D overrides, plain libc names.
 */
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

#define VKMEM_EXPORT __attribute__((used, visibility("default")))

extern int main(int argc, char **argv);

/* args: NUL-separated argv strings, len bytes total (including trailing NUL). */
VKMEM_EXPORT int vkmem_main(char *args, int len)
{
	int argc = 0, i;
	char **argv;
	char *p = args;
	for (i = 0; i < len; i++)
		if (args[i] == 0)
			argc++;
	argv = (char **) malloc(sizeof(char *) * (argc + 1));
	for (i = 0; i < argc; i++) {
		argv[i] = p;
		p += strlen(p) + 1;
	}
	argv[argc] = NULL;
	return main(argc, argv);
}

/*
 * Emscripten's libc ships weak "unsupported syscall" stubs for these; the
 * strong definitions here win at link time. Socket options are accepted
 * and ignored (the host owns the real sockets), and resource limits are
 * reported as unlimited so Valkey does not try to raise them.
 */
int __syscall_setsockopt(int sockfd, int level, int optname, intptr_t optval, size_t optlen, int dummy)
{
	(void) sockfd; (void) level; (void) optname; (void) optval; (void) optlen; (void) dummy;
	return 0;
}

int __syscall_prlimit64(int pid, int resource, intptr_t new_limit, intptr_t old_limit)
{
	(void) pid; (void) resource; (void) new_limit;
	if (old_limit) {
		uint64_t *lim = (uint64_t *) old_limit;
		lim[0] = ~(uint64_t) 0; /* rlim_cur = RLIM_INFINITY */
		lim[1] = ~(uint64_t) 0; /* rlim_max */
	}
	return 0;
}

int __syscall_setrlimit(int resource, intptr_t limit)
{
	(void) resource; (void) limit;
	return 0;
}

int __syscall_getrusage(int who, intptr_t usage)
{
	(void) who;
	if (usage)
		memset((void *) usage, 0, 152); /* sizeof(struct rusage) on wasm32 */
	return 0;
}

/*
 * dlopen/dlsym replacement (Valkey sources are compiled with
 * -Ddlopen=vkmem_dlopen ...): the build has no dynamic linking, and the only
 * use is module.c resolving the statically linked Lua engine through
 * dlopen(NULL) + dlsym("ValkeyModule_OnLoad_lua").
 */
extern int ValkeyModule_OnLoad_lua(void *ctx, void **argv, int argc);
extern int ValkeyModule_OnUnload_lua(void *ctx);

static const struct { const char *name; void *fn; } vkmem_static_symbols[] = {
	{"ValkeyModule_OnLoad_lua", (void *) ValkeyModule_OnLoad_lua},
	{"ValkeyModule_OnUnload_lua", (void *) ValkeyModule_OnUnload_lua},
};
static int vkmem_self_handle;
static const char *vkmem_dl_error;

VKMEM_EXPORT void *vkmem_dlopen(const char *file, int mode)
{
	(void) mode;
	if (file == NULL)
		return &vkmem_self_handle;
	vkmem_dl_error = "valkeymem: loadable modules are not supported";
	return NULL;
}

VKMEM_EXPORT void *vkmem_dlsym(void *handle, const char *name)
{
	size_t i;
	if (handle == &vkmem_self_handle) {
		for (i = 0; i < sizeof(vkmem_static_symbols) / sizeof(vkmem_static_symbols[0]); i++)
			if (strcmp(vkmem_static_symbols[i].name, name) == 0)
				return vkmem_static_symbols[i].fn;
	}
	vkmem_dl_error = "valkeymem: symbol not found";
	return NULL;
}

VKMEM_EXPORT int vkmem_dlclose(void *handle)
{
	(void) handle;
	return 0;
}

VKMEM_EXPORT char *vkmem_dlerror(void)
{
	const char *e = vkmem_dl_error;
	vkmem_dl_error = NULL;
	return (char *) e;
}

/* Called by the host to implement the env.__call_sighandler import. */
VKMEM_EXPORT void vkmem_call_sighandler(int fp, int sig)
{
	((void (*)(int)) fp)(sig);
}
