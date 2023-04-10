# struct-bingen

## Test

```shell
    cp config-test.yaml config.yaml
    ### Add your gcc/gpp search path in config.yaml
    ### Add predefined files according yours arch
```

````shell
    go run main.go --source testdata/bpf/cpu.h --target testdata/cpu/cpu.go
    cd testdata/cpu
    go test
````