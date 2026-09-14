/*
 * vkmem_sha1.inc: SHA-1 on the host (Go crypto/sha1). Included by sha1.c
 * under __VKMEM__ in place of the C implementation; the context keeps the
 * host handle in state[0]. Every caller follows Init/Update/Final.
 */
#include <stdint.h>
#include "sha1.h"

#define VKMEM_IMPORT(name) __attribute__((import_module("env"), import_name(name)))
VKMEM_IMPORT("vkmem_hash_create") extern int vkmem_hash_create(int type);
VKMEM_IMPORT("vkmem_hash_update") extern void vkmem_hash_update(int h, const void *data, int len);
VKMEM_IMPORT("vkmem_hash_final") extern void vkmem_hash_final(int h, void *out, int max);

void SHA1Transform(uint32_t state[5], const unsigned char buffer[64])
{
	(void) state; (void) buffer; /* not used outside sha1.c */
}

void SHA1Init(SHA1_CTX *context)
{
	context->state[0] = (uint32_t) vkmem_hash_create(1);
}

void SHA1Update(SHA1_CTX *context, const unsigned char *data, uint32_t len)
{
	vkmem_hash_update((int) context->state[0], data, (int) len);
}

void SHA1Final(unsigned char digest[20], SHA1_CTX *context)
{
	vkmem_hash_final((int) context->state[0], digest, 20);
	context->state[0] = 0;
}
