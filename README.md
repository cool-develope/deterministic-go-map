# Deterministic Map

Go-map is not deterministic, we need to match the order of map range in test of state machine.
Here, it suggests simple deterministic map based on the arrival order.

## Benchmark

```
BenchmarkDMapSet-4      1000000000               0.7382 ns/op          0 B/op          0 allocs/op
BenchmarkMapSet-4       1000000000               0.5826 ns/op          0 B/op          0 allocs/op
```