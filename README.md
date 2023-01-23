Example git repository to exporting and importing an Value array from Go memory to a file without any encoder. It might usefull who need very fast encoder for init time stuffs. However it is using unsafe package which is "unsafe". Also exported and imported binaries are might be not compatible with other operation systems or architectures.

```go test -bench=. -benchmem -count 3```

```
goos: linux
goarch: amd64
pkg: cot
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkArrayEncode-8           	1000000000	         0.4930 ns/op	       0 B/op	       0 allocs/op
BenchmarkArrayEncode-8           	1000000000	         0.4886 ns/op	       0 B/op	       0 allocs/op
BenchmarkArrayEncode-8           	1000000000	         0.4868 ns/op	       0 B/op	       0 allocs/op
BenchmarkArrayDecode-8           	390091548	         2.932 ns/op	       0 B/op	       0 allocs/op
BenchmarkArrayDecode-8           	389081352	         2.956 ns/op	       0 B/op	       0 allocs/op
BenchmarkArrayDecode-8           	393874976	         2.901 ns/op	       0 B/op	       0 allocs/op
BenchmarkPointerArraypDecode-8   	   33084	     36243 ns/op	   81920 B/op	       1 allocs/op
BenchmarkPointerArraypDecode-8   	   32616	     36361 ns/op	   81920 B/op	       1 allocs/op
BenchmarkPointerArraypDecode-8   	   32690	     36658 ns/op	   81920 B/op	       1 allocs/op
BenchmarkEncode-8                	1000000000	         0.4796 ns/op	       0 B/op	       0 allocs/op
BenchmarkEncode-8                	1000000000	         0.4726 ns/op	       0 B/op	       0 allocs/op
BenchmarkEncode-8                	1000000000	         0.4713 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecode-8                	1000000000	         0.3229 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecode-8                	1000000000	         0.3321 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecode-8                	1000000000	         0.3271 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	cot	13.480s
```

I wrote my memory into `memout_littleendian_x64` file. You can try to execute my exported memory with

```bash
cot --path memout_littleendian_x64 read
```


```bash
cot --path myCustomMem --size 100_000 write
```