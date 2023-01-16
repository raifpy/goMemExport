Example git repository to exporting and importing an Value array from Go memory to a file without any encoder. It might usefull who need very fast encoder for init time stuffs. However it is using unsafe package which is "unsafe". Also exported and imported binaries are might be not compatible with other operation systems or architectures.

```go
cot --path filemem read
```


```go
cot --path filemem --size 100_000 write
```