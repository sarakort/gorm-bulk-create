
Benchmark
=========
```
go test -bench=. -benchmem  
goos: darwin
goarch: amd64
pkg: github.com/sarakort/gorm-bulk-create
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

BenchmarkOrmCreate-12                                332           3400990 ns/op            8035 B/op        101 allocs/op
BenchmarkOrmCreateAll-12                            6734           2034451 ns/op         5153891 B/op         46 allocs/op
BenchmarkOrmCreateAllBatchSize100-12               39424             27299 ns/op            3589 B/op         24 allocs/op
BenchmarkOrmCreateAllBatchSize500-12               48614             22209 ns/op            3703 B/op         25 allocs/op
BenchmarkOrmCreateAllBatchSize1000-12              56954             21611 ns/op            3927 B/op         25 allocs/op
BenchmarkBulkCreate-12                             60891             21411 ns/op            2562 B/op         14 allocs/op
BenchmarkBulkCreateSize1-12                         1069            986167 ns/op            2805 B/op         40 allocs/op
BenchmarkBulkCreateSize100-12                      43088             25499 ns/op            3173 B/op         19 allocs/op
BenchmarkBulkCreateSize500-12                      58972             20976 ns/op            3309 B/op         20 allocs/op
BenchmarkBulkCreateSize1000-12                     63199             19782 ns/op            3560 B/op         20 allocs/op
PASS
ok      github.com/sarakort/gorm-bulk-create    30.112s
```