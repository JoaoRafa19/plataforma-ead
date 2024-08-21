# Plataforma EAD

### Comandos principais
***Rodando o container***
```shell
make up
```
___(ou com logs)___
```shell
MODE=l make up
```

***Logs***
>Com o container já de pé ele vai acoplar o terminal ao terminal de logs do docker.

```shell
make log
```

Restart
> Reinicia o container
```shell
make restart
```

Parar
>Encerra a execução da aplicação
```shell
make down
```


# Service Core

## Dev dependencies

Protobuf compiler
```
brew install protobuf
```

AIR (live reload do go)
```
go install github.com/air-verse/air@latest 
```
### Comandos do projeto
Compilar os arquivo do protobuffer
```shell
make gen-grpc
```

## Testes

### Service Core
Executar testes service core
```shell
make test-service-core
```


## Scripts

Compila os protobuffers
```
protoc --go_out=./service-core --go-grpc_out=./service-core ./service-core/protos/*.proto
```
