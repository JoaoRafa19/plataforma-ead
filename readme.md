## Dev dependencies

Protobuf compiler
```
brew install protobuf
```


## Scripts

Compila os protobuffers
```
protoc --go_out=./service-core --go-grpc_out=./service-core ./service-core/protos/*.proto
```

