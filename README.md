# GWSync

Used to sync Gearworks playerdata and more

## proto

For Go use the following to generate files

```bash
protoc --go_out=./proto/ --go_opt=paths=source_relative --go-grpc_out=./proto/ --go-grpc_opt=paths=source_relative gearworks.proto
```

For Java use the following to generate files

```bash
protoc -I=. --java_out=. gearworks.proto
```