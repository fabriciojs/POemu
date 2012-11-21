go tool pprof bin/searchsm profiling/mem --pdf > profiling/mem.pdf;
go tool pprof bin/searchsm profiling/x64 --pdf > profiling/cpu.pdf;