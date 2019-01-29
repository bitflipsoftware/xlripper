rm -f bench.cpu.out && \
GOCACHE=off go test -cpuprofile bench_small.cpu.out -bench=BenchSmall && \
go tool pprof --pdf bench_small.cpu.out > bench_small.cpu.pdf && \
open bench_small.cpu.pdf

rm -f bench_qv.cpu.out && \
GOCACHE=off go test -cpuprofile bench_qv.cpu.out -bench=BenchQv && \
go tool pprof --pdf bench_qv.cpu.out > bench_qv.cpu.pdf && \
open bench_qv.cpu.pdf

rm -f bench_large.cpu.out && \
GOCACHE=off go test -cpuprofile bench_large.cpu.out -bench -test.v && \
go tool pprof --pdf bench_large.cpu.out > bench_large.cpu.pdf && \
open bench_large.cpu.pdf