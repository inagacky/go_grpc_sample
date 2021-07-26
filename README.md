# go_grpc_sample

Generate Protocol Buffer File
```
$ protoc -I ./proto cat.proto --go_out=plugins=grpc:./pb --go_opt=module=github.com/inagacky/go_grpc_sample/pb
$ protoc -I ./proto helloworld.proto --go_out=plugins=grpc:./pb --go_opt=module=github.com/inagacky/go_grpc_sample/pb
```
