# protobuf-tutorial

## json formula

```
curl http://localhost:8080 -d '{"op":"*","args":[1,2,3,6]}'
```

## protobuf compile

```
go get -u github.com/golang/protobuf/protoc-gen-go
sudo apt  install protobuf-compiler
protoc -I=api --go_out=api api/formula.proto
```

## test

```
gotest ./... -v 
```
