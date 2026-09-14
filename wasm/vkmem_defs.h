/* Force-included into every Valkey translation unit (build.sh -include). */
#ifndef VKMEM_DEFS_H
#define VKMEM_DEFS_H
/*
 * valkeymodule.h declares the module API function pointers with
 * __attribute__((__common__)); the wasm backend of clang 20 crashes on
 * common symbols (AsmPrinter::emitGlobalVariable). Weak definitions give
 * the same "many TUs, one symbol" effect and link fine with wasm-ld.
 */
#define VALKEYMODULE_ATTR __attribute__((weak))
#endif
