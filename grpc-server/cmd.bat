cd proto && protoc --go_out=plugins=grpc:../services Prot.proto
protoc --go_out=plugins=grpc:../services Orders.proto
protoc --go_out=plugins=grpc:../services User.proto

protoc --go_out=plugins=grpc:../services --validate_out=lang=go:../services Model.proto

protoc --grpc-gateway_out=logtostderr=true:../services Prot.proto
protoc --grpc-gateway_out=logtostderr=true:../services Orders.proto
protoc --grpc-gateway_out=logtostderr=true:../services User.proto

cd ..