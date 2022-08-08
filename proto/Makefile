gen-go:
	rm -rf grpc
	rm -rf temp
	mkdir temp

	protoc --proto_path=proto proto/**/*.proto --go_out=temp
	protoc --proto_path=proto proto/**/*.proto  --go-grpc_out=./temp

	ls temp/proto/grpc/*/*.pb.go | xargs -n1 -Ix bash -c 'sed s/,omitempty// x > x.tmp && mv x{.tmp,}'

	mv temp/proto/grpc .
	rm -rf temp
.PHONY: gen-go
