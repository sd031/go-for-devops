### Run Benchmark
go test -bench=.

### CPU Profiling
go test -bench=. -cpuprofile=cpu.prof

### CPU Profiling
go test -bench=. -memprofile=mem.prof

### Analyze Profile
go tool pprof cpu.prof

* For better UI Visualisation try: https://www.speedscope.app/ or below commad:
go tool pprof -http=:8080 cpu.prof  
(Install graphviz for this : https://graphviz.org/download/)


