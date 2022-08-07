gen-go:
	rm -rf grpc
	rm -rf grpc2
	mkdir grpc2

	protoc --proto_path=proto proto/**/*.proto --go_out=grpc2
	protoc --proto_path=proto proto/**/*.proto  --go-grpc_out=./grpc2

	ls grpc2/proto/grpc/*/*.pb.go | xargs -n1 -Ix bash -c 'sed s/,omitempty// x > x.tmp && mv x{.tmp,}'

	mv grpc2/proto/grpc .
	rm -rf grpc2
.PHONY: gen-go
