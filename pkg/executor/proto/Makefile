
PROTO_FILES=$(sort $(wildcard ./*.proto))

all: $(PROTO_FILES) Makefile
	make executor

executor:
	# generate the gRPC code
	protoc3.6.exe --go_out=plugins=grpc:../pb/ ${PROTO_FILES}
	protoc3.6.exe --go_out=plugins=grpc:../pb/ executor.proto
