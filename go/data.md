Array size: 5000

goos: linux
goarch: amd64
pkg: bsearch_benchmark
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkOneComp-2      10856928                95.36 ns/op
BenchmarkTwoComp-2      14542315                81.79 ns/op
PASS
ok      bsearch_benchmark       2.430s
user@work ~/g/e/b/go (main *)> go test -bench=.
goos: linux
goarch: amd64
pkg: bsearch_benchmark
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkOneComp-2      10535823                95.49 ns/op
BenchmarkTwoComp-2      14488572                82.35 ns/op
PASS
ok      bsearch_benchmark       2.407s
user@work ~/g/e/b/go (main *)> go test -bench=.
goos: linux
goarch: amd64
pkg: bsearch_benchmark
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkOneComp-2      11862303                95.07 ns/op
BenchmarkTwoComp-2      14745253                81.72 ns/op
PASS
ok      bsearch_benchmark       2.528s
user@work ~/g/e/b/go (main *)> go test -bench=.
goos: linux
goarch: amd64
pkg: bsearch_benchmark
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkOneComp-2      11421444                93.94 ns/op
BenchmarkTwoComp-2      14655752                80.36 ns/op
PASS
ok      bsearch_benchmark       2.450s
user@work ~/g/e/b/go (main *)> go test -bench=.
goos: linux
goarch: amd64
pkg: bsearch_benchmark
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkOneComp-2      11875916                94.01 ns/op
BenchmarkTwoComp-2      14531046                80.43 ns/op
PASS
ok      bsearch_benchmark       2.482s

Array size: 100000

goos: linux
goarch: amd64
pkg: bsearch_benchmark
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkOneComp-2   	 5919397	       200.2 ns/op
BenchmarkTwoComp-2   	 5987878	       179.3 ns/op
PASS
ok  	bsearch_benchmark	2.738s
user@work ~/g/e/b/go (main *)> go test -bench=.
goos: linux
goarch: amd64
pkg: bsearch_benchmark
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkOneComp-2   	 6063636	       194.7 ns/op
BenchmarkTwoComp-2   	 6690996	       175.4 ns/op
PASS
ok  	bsearch_benchmark	2.808s
user@work ~/g/e/b/go (main *)> go test -bench=.
goos: linux
goarch: amd64
pkg: bsearch_benchmark
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkOneComp-2   	 5835162	       195.9 ns/op
BenchmarkTwoComp-2   	 6792552	       174.2 ns/op
PASS
ok  	bsearch_benchmark	2.776s
user@work ~/g/e/b/go (main *)> go test -bench=.
goos: linux
goarch: amd64
pkg: bsearch_benchmark
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkOneComp-2   	 6108228	       195.2 ns/op
BenchmarkTwoComp-2   	 6799393	       174.6 ns/op
PASS
ok  	bsearch_benchmark	2.819s
user@work ~/g/e/b/go (main *)> go test -bench=.
goos: linux
goarch: amd64
pkg: bsearch_benchmark
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkOneComp-2   	 6022304	       203.6 ns/op
BenchmarkTwoComp-2   	 6622332	       190.0 ns/op
PASS
ok  	bsearch_benchmark	2.929s

