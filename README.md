# `struct-bingen`

**`bingen` automatically generates Go structs bindings to C (and some C++) libraries.**

For example, given the C header `testdata/bpf/cpu.h`:

```c
typedef struct {
  u64 low_pc;
  u64 high_pc;
  u64 shard_index;
  u64 low_index;
  u64 high_index;
} chunk_info_t;
```

`struct-bingen` produces Go struct code allowing you to call into the `cpu` library's
functions and use its types:

```go
type ChunkInfoT struct {
	LowPc      uint64
	HighPc     uint64
	ShardIndex uint64
	LowIndex   uint64
	HighIndex  uint64
}
```

## Usage

````shell
    ### Add your gcc/gpp search path in config.yaml
    ### Add predefined files according yours arch
    go run main.go --source testdata/bpf/cpu.h --target testdata/cpu/cpu.go
    cd testdata/cpu
    go test
````
Although it varies according to operating systems and your configuration settings, you can usually run the command below to find the search path.

````shell
    echo | gcc -E -Wp,-v -
````
Example yaml configuration. You can check your gcc search path and cpu arch defined in predefined for preprocessor.
````yaml
include_paths:
  - "/usr/local/include"
  - "/Library/Developer/CommandLineTools/usr/lib/clang/14.0.3/include"
  - "/Users/YOURNAME/struct-bingen/exclude/c/header"
sys_include_paths:
  - "/usr/local/include"
  - "/Library/Developer/CommandLineTools/usr/lib/clang/14.0.3/include"
  - "/Users/YOURNAME/struct-bingen/exclude/c/header"
predefine:
  - "#define __x86_64__ 1"
````
## Notice

only tested on macos darwin and ubuntu x86-64. Necessary test coverage for all architectures and os will be released in the next realases.

## Memory Allignment

Unfortunately, the packaged builds do not give the correct result in size, as there is no build compression in Go. I hope to fix this in future versions.

