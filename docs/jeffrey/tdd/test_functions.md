`func AllocsPerRun(runs int, f func()) (avg float64)`

- description: `AllocsPerRun()` 는 float 로 할당한 평균 수를 반환합니다. 비록 float64 의 타입을 가지고는 있지만 항상
                정수 형태의 수를 반환합니다. 할당 수를 계산하기 위해 기능은 먼저 워밍업으로 한 번 실행됩니다.
                그러면 지정된 실행 횟수에 대한 평균 할당 수가 측정되어 반환됩니다.

- usage: ?

- implementation:

```go
package testing
import "runtime"

func AllocsPerRun(runs int, f func()) (avg float64) {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(1))

	// Warm up the function
	f()

	// Measure the starting statistics
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)
	mallocs := 0 - memstats.Mallocs

	// Run the function the specified number of times
	for i := 0; i < runs; i++ {
		f()
	}

	// Read the final statistics
	runtime.ReadMemStats(&memstats)
	mallocs += memstats.Mallocs

	// Average the mallocs over the runs (not counting the warm-up).
	// We are forced to return a float64 because the API is silly, but do
	// the division as integers so we can ask if AllocsPerRun()==1
	// instead of AllocsPerRun()<2.
	return float64(mallocs / uint64(runs))
}

```





`func CoverMode() string`

- description:
- usage:
- implementation
- example

`func Coverage() float64`

- description:
- usage:
- implementation
- example

func Init()
func Main(matchString func(pat, str string) (bool, error), tests []InternalTest, ...)
func RegisterCover(c Cover)
func RunBenchmarks(matchString func(pat, str string) (bool, error), ...)
func RunExamples(matchString func(pat, str string) (bool, error), examples []InternalExample) (ok bool)
func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)
func Short() bool
func Verbose() bool
type B
func (c *B) Cleanup(f func())
func (c *B) Error(args ...any)
func (c *B) Errorf(format string, args ...any)
func (c *B) Fail()
func (c *B) FailNow()
func (c *B) Failed() bool
func (c *B) Fatal(args ...any)
func (c *B) Fatalf(format string, args ...any)
func (c *B) Helper()
func (c *B) Log(args ...any)
func (c *B) Logf(format string, args ...any)
func (c *B) Name() string
func (b *B) ReportAllocs()
func (b *B) ReportMetric(n float64, unit string)
func (b *B) ResetTimer()
func (b *B) Run(name string, f func(b *B)) bool
func (b *B) RunParallel(body func(*PB))
func (b *B) SetBytes(n int64)
func (b *B) SetParallelism(p int)
func (c *B) Setenv(key, value string)
func (c *B) Skip(args ...any)
func (c *B) SkipNow()
func (c *B) Skipf(format string, args ...any)
func (c *B) Skipped() bool
func (b *B) StartTimer()
func (b *B) StopTimer()
func (c *B) TempDir() string
type BenchmarkResult
func Benchmark(f func(b *B)) BenchmarkResult
func (r BenchmarkResult) AllocedBytesPerOp() int64
func (r BenchmarkResult) AllocsPerOp() int64
func (r BenchmarkResult) MemString() string
func (r BenchmarkResult) NsPerOp() int64
func (r BenchmarkResult) String() string
type Cover
type CoverBlock
type F
func (f *F) Add(args ...any)
func (c *F) Cleanup(f func())
func (c *F) Error(args ...any)
func (c *F) Errorf(format string, args ...any)
func (f *F) Fail()
func (c *F) FailNow()
func (c *F) Failed() bool
func (c *F) Fatal(args ...any)
func (c *F) Fatalf(format string, args ...any)
func (f *F) Fuzz(ff any)
func (f *F) Helper()
func (c *F) Log(args ...any)
func (c *F) Logf(format string, args ...any)
func (c *F) Name() string
func (c *F) Setenv(key, value string)
func (c *F) Skip(args ...any)
func (c *F) SkipNow()
func (c *F) Skipf(format string, args ...any)
func (f *F) Skipped() bool
func (c *F) TempDir() string
type InternalBenchmark
type InternalExample
type InternalFuzzTarget
type InternalTest
type M
func MainStart(deps testDeps, tests []InternalTest, benchmarks []InternalBenchmark, ...) *M
func (m *M) Run() (code int)
type PB
func (pb *PB) Next() bool
type T
func (c *T) Cleanup(f func())
func (t *T) Deadline() (deadline time.Time, ok bool)
func (c *T) Error(args ...any)
func (c *T) Errorf(format string, args ...any)
func (c *T) Fail()
func (c *T) FailNow()
func (c *T) Failed() bool
func (c *T) Fatal(args ...any)
func (c *T) Fatalf(format string, args ...any)
func (c *T) Helper()
func (c *T) Log(args ...any)
func (c *T) Logf(format string, args ...any)
func (c *T) Name() string
func (t *T) Parallel()
func (t *T) Run(name string, f func(t *T)) bool
func (t *T) Setenv(key, value string)
func (c *T) Skip(args ...any)
func (c *T) SkipNow()
func (c *T) Skipf(format string, args ...any)
func (c *T) Skipped() bool
func (c *T) TempDir() string
type TB
Examples ¶
B.ReportMetric
B.RunParallel
