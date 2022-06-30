# 2someone-proto

## Buf
Generate code  
```sh
$ buf generate
```

## Golang 
```sh
$ go get github.com/leaper-one/2someone-proto/gen/golang
```

## 需要安装的插件
1. protoc  
2. protoc-gen-go  
3. `go install go-micro.dev/v4/cmd/protoc-gen-micro@v4`  
## buf.gen.yaml
自动编译配置文件
name：--插件名称
out：--编译后.go文件路径


## buf.yaml
识别.proto文件用
roots：--.proto目录
