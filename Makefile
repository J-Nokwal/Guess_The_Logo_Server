gen:
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=. --go-grpc_out=. proto/*.proto
clean:
	rm pb/*.go

pd:
	pwd