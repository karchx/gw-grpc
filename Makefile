protoc:
	cd proto && protoc --go_out=../protogen/output --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/output --go-grpc_opt=paths=source_relative \
	./**/*.proto
