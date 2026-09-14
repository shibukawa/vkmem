#!/usr/bin/env python3
"""Idempotent vkmem source patches for the Valkey checkout.

usage: patches.py <valkey dir>

Everything is guarded by __VKMEM__ so the tree still builds unpatched.
"""
import sys, os

src = sys.argv[1]

def patch(rel, old, new, marker):
    p = os.path.join(src, rel)
    s = open(p).read()
    if marker in s:
        print(f'{rel}: already patched')
        return
    if old not in s:
        sys.exit(f'{rel}: anchor not found')
    open(p, 'w').write(s.replace(old, new, 1))
    print(f'{rel}: patched')

# bio.c: the wasm build has no threads. Background jobs (lazy free, file
# close/fsync) run synchronously at submit time instead of on worker
# threads, and bioInit spawns nothing.
patch('src/bio.c',
'''    for (bio_worker_data *bwd = bio_workers; bwd != bio_worker_end; ++bwd) {
        int err = pthread_create(&bwd->bio_thread_id, &attr, bioProcessBackgroundJobs, (void *)bwd);''',
'''#ifdef __VKMEM__
    if (1) { pthread_attr_destroy(&attr); return; }
#endif
    for (bio_worker_data *bwd = bio_workers; bwd != bio_worker_end; ++bwd) {
        int err = pthread_create(&bwd->bio_thread_id, &attr, bioProcessBackgroundJobs, (void *)bwd);''',
'__VKMEM__')

patch('src/bio.c',
'''void bioSubmitJob(int type, bio_job *job) {
    job->header.type = type;
''',
'''static void bioExecuteJob(bio_job *job);

void bioSubmitJob(int type, bio_job *job) {
    job->header.type = type;
#ifdef __VKMEM__
    atomic_fetch_add(&bio_jobs_counter[type], 1);
    bioExecuteJob(job);
    return;
#endif
''',
'bioExecuteJob(job);\n    return;')

patch('src/bio.c',
'''    while (1) {
        bio_job *job = mutexQueuePop(bwd->bio_jobs, true);

        /* Process the job accordingly to its type. */
        int job_type = job->header.type;
''',
'''    while (1) {
        bio_job *job = mutexQueuePop(bwd->bio_jobs, true);
        bioExecuteJob(job);
    }
}

/* Process one job accordingly to its type and free it. */
static void bioExecuteJob(bio_job *job) {
    {
        int job_type = job->header.type;
''',
'static void bioExecuteJob(bio_job *job) {')

# util.c: on non-Linux getTimeZone() reads the struct timezone that
# gettimeofday() is supposed to fill; Emscripten's gettimeofday ignores it,
# so the log timestamps were shifted by stack garbage. Use libc's timezone.
patch('src/util.c',
'''long getTimeZone(void) {
#if defined(__linux__) || defined(__sun)''',
'''long getTimeZone(void) {
#if defined(__linux__) || defined(__sun) || defined(__VKMEM__)''',
'defined(__sun) || defined(__VKMEM__)')

# Hashes on the host: sha1.c / sha256.c keep their C bodies for builds
# without __VKMEM__; with it the contexts hold a handle to a Go hash (see
# wasm/vkmem_sha1.inc, vkmem_sha256.inc and internal/host/crypto.go).
for rel, inc in (('src/sha1.c', 'vkmem_sha1.inc'), ('src/sha256.c', 'vkmem_sha256.inc')):
    p = os.path.join(src, rel)
    s = open(p).read()
    if inc in s:
        print(f'{rel}: already patched')
    else:
        open(p, 'w').write(f'#ifdef __VKMEM__\n#include "{inc}"\n#else\n' + s + '\n#endif /* __VKMEM__ */\n')
        print(f'{rel}: patched')

# crc64(): RDB/DUMP checksums go to Go's hash/crc64 (same Jones CRC).
patch('src/crc64.c',
'''uint64_t crc64(uint64_t crc, const unsigned char *s, uint64_t l) {
    return crcspeed64native(crc64_table, crc, (void *) s, l);
}''',
'''#ifdef __VKMEM__
extern uint64_t vkmem_crc64(uint64_t crc, const unsigned char *s, uint64_t l)
    __attribute__((import_module("env"), import_name("vkmem_crc64")));
#endif
uint64_t crc64(uint64_t crc, const unsigned char *s, uint64_t l) {
#ifdef __VKMEM__
    return vkmem_crc64(crc, s, l);
#else
    return crcspeed64native(crc64_table, crc, (void *) s, l);
#endif
}''',
'vkmem_crc64')
