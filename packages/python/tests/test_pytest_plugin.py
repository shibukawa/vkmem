def test_vkmem_fixture_starts_an_isolated_fork(vkmem_fork, vkmem_dsn):
    assert vkmem_fork.dsn == vkmem_dsn
    assert vkmem_fork.port > 0
