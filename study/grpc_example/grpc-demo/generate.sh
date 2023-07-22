# go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
# go

protoc --go_out=. --go-grpc_out=. example.proto

# cpp
# brew install grpc
# vcpkg install grpc

protoc --cpp_out=./generated --grpc_out=./generated --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` example.proto


# nodejs
# npm i -g grpc-tools
# npm i -g protoc-gen-js
protoc --js_out=import_style=commonjs,binary:./generated --grpc_out=./generated --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` example.proto

# rust

# use cloudwego/volo
# or use thonic
