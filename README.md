# Golang Stack Learning

## Protobuf
* Install Protoc in Windows using Choco : run `choco install protoc --pre `
* Upgrade Protoc in Windows using Choco : run `choco upgrade protoc --pre `
* Uninstall Protoc in Windows using Choco : run `choco uninstall protoc --pre `
* Convert `.proto` to `.go` : run `protoc --go_out=. *.proto` or `protoc --go_out=plugins=grpc:. *.proto` to handle gRPC support (service tag)

# Running
## GRPC
1. Run Garage Service with `cd  grpc/services/service-garage/`, then run `go run main.go`
2. Run Member Service with `cd  grpc/services/service-member/`, then run `go run main.go`
3. Run Client Service with `cd grpc/client/`, then run `go run main.go`

## REST API
* In root directory, run `go run main.go`

## Websocket - Chatting
* Run chatting app with `cd websocket/`, then run `go run main.go`, the app using port `:8050`. Access `localhost:8050` to open the app. 

## Reference
* https://chocolatey.org/packages/protoc
* https://dasarpemrogramangolang.novalagung.com/C-29-golang-protobuf-implementation.html
* https://dasarpemrogramangolang.novalagung.com/C-30-golang-grpc-protobuf.html
* https://dasarpemrogramangolang.novalagung.com/C-28-golang-web-socket.html