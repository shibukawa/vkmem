/*
 * vkmem_sha256.inc: SHA-256 on the host (Go crypto/sha256). Included by
 * sha256.c under __VKMEM__; the host handle lives in ctx->state[0].
 */
#include <stdint.h>
#include "sha256.h"

#define VKMEM_IMPORT(name) __attribute__((import_module("env"), import_name(name)))
VKMEM_IMPORT("vkmem_hash_create") extern int vkmem_hash_create(int type);
VKMEM_IMPORT("vkmem_hash_update") extern void vkmem_hash_update(int h, const void *data, int len);
VKMEM_IMPORT("vkmem_hash_final") extern void vkmem_hash_final(int h, void *out, int max);

void sha256_transform(SHA256_CTX *ctx, const BYTE data[])
{
	(void) ctx; (void) data;
}

void sha256_init(SHA256_CTX *ctx)
{
	ctx->state[0] = (WORD) vkmem_hash_create(2);
}

void sha256_update(SHA256_CTX *ctx, const BYTE data[], size_t len)
{
	vkmem_hash_update((int) ctx->state[0], data, (int) len);
}

void sha256_final(SHA256_CTX *ctx, BYTE hash[])
{
	vkmem_hash_final((int) ctx->state[0], hash, 32);
	ctx->state[0] = 0;
}
