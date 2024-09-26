protoc:
	cd proto && protoc --go_out=../protogen/output --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/output --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../protogen/output --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	./**/*.proto
