# protobuf-tutorial

## json formula

start server

```
go run main.go
```

server app outputs length of binary request

### calc handler

```
curl http://localhost:8080/calc -d '{"op":"*","args":[1,2,3,6]}'
# or
curl http://localhost:8080/calc -d @sample.json
```

### calc ProtoBuf handler

```
go run test/protobuf_requestor.go
# or 
curl http://localhost:8080/calcpb --data-binary @dat
```

## protobuf compile

```
go get -u github.com/golang/protobuf/protoc-gen-go
sudo apt  install protobuf-compiler
protoc -I=api --go_out=api api/*.proto
```

## test

```
gotest ./... -v 
```
