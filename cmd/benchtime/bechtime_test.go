package benchtime_test

import (
	"sync"
	"testing"
	"time"
)

func BenchmarkTimeAfterMicro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		<-time.After(time.Microsecond)
	}
}

func BenchmarkTimeSleepMicro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Microsecond)
	}
}

func BenchmarkTimeAfterMili(b *testing.B) {
	for i := 0; i < b.N; i++ {
		<-time.After(time.Millisecond)
	}
}

func BenchmarkTimeSleepMili(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond)
	}
}

func BenchmarkTimeAfterSec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		<-time.After(time.Second)
	}
}

func BenchmarkTimeSleepSec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Second)
	}
}

func BenchmarkTimeAfterSecConcurrent(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		func() {
			<-time.After(time.Second / 2)
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkTimeSleepSecConcurrent(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		func() {
			time.Sleep(time.Second / 2)
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/cmd/benchtime
cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
BenchmarkTimeAfterMicro
BenchmarkTimeAfterMicro-8           	   17702	     64802 ns/op	     200 B/op	       3 allocs/op
BenchmarkTimeSleepMicro
BenchmarkTimeSleepMicro-8           	   70459	     16928 ns/op	       0 B/op	       0 allocs/op
BenchmarkTimeAfterMili
BenchmarkTimeAfterMili-8            	    1066	   1142579 ns/op	     200 B/op	       3 allocs/op
BenchmarkTimeSleepMili
BenchmarkTimeSleepMili-8            	    1069	   1140064 ns/op	       0 B/op	       0 allocs/op
BenchmarkTimeAfterSec
BenchmarkTimeAfterSec-8             	       1	1000166030 ns/op	     200 B/op	       3 allocs/op
BenchmarkTimeSleepSec
BenchmarkTimeSleepSec-8             	       1	1000256042 ns/op	      80 B/op	       1 allocs/op
BenchmarkTimeAfterSecConcurrent
BenchmarkTimeAfterSecConcurrent-8   	       1	1000188685 ns/op	     216 B/op	       4 allocs/op
BenchmarkTimeSleepSecConcurrent
BenchmarkTimeSleepSecConcurrent-8   	       1	1000279153 ns/op	      96 B/op	       2 allocs/op
PASS
coverage: [no statements]
ok  	github.com/thanhpp/gopher/cmd/benchtime	11.516s
*/
