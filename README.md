```
protoc --go_out=plugins=grpc:advcalc advcalc.proto 
go run server.go
go run client.go
```
