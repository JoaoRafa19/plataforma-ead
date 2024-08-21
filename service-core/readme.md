## Dev dependencies

Protobuf compiler
```
brew install protobuf
```

AIR (nodemon do go/docker)
```
go install github.com/air-verse/air@latest
```


## Scripts

Compila os protobuffers
```
protoc --go_out=./service-core --go-grpc_out=./service-core ./service-core/protos/*.proto
```
