
Benchmark
=========
```
BenchmarkOrmCreate-12                                332           3400990 ns/op            8035 B/op        101 allocs/op
BenchmarkOrmCreateAllBatchSize100-12               35943             28943 ns/op            3642 B/op         24 allocs/op
BenchmarkOrmCreateAllBatchSize500-12               49316             22280 ns/op            3697 B/op         25 allocs/op
BenchmarkOrmCreateAllBatchSize1000-12              56979             22524 ns/op            3925 B/op         25 allocs/op
BenchmarkBulkCreate-12                             61497             19028 ns/op            2562 B/op         14 allocs/op
BenchmarkBulkCreateSize1-12                         1072            972102 ns/op            2802 B/op         40 allocs/op
BenchmarkBulkCreateSize100-12                      43797             27199 ns/op            3173 B/op         19 allocs/op
BenchmarkBulkCreateSize500-12                      52965             21378 ns/op            3304 B/op         20 allocs/op
BenchmarkBulkCreateSize1000-12                     55068             20212 ns/op            3559 B/op         20 allocs/op
```