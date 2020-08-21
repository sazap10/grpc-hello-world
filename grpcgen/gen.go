package grpcgen

// Add go generate lines to the bottom of this file to auto generate the corresponding grpc go code.
//
// Prerequistes
// 1. Protobuf tools installed     ( brew install protobuf )
// 2. Go generator plugin installed  ( go get -u github.com/golang/protobuf/protoc-gen-go )
//
// Adding a grpc service to the app
// 1, Create a folder (or use an existing folder) under grpcgen with the name of the package
// 2. Copy the proto file and paste it in the folder, renaming the file to include the version e.g. v1.2.3 of test.proto becomes test.v1-2-3.proto
// 3. Add a go generate comment for generating the command ending with the path to the file
//    i.e. //go:generate protoc --go_out=plugins=grpc:. ./<package_name>/<service_name>.<version>.proto
// 4. Run go generate grpcgen/gen.go

//go:generate protoc --go_out=plugins=grpc:. ./chat-service/chat.v1-0-0.proto
