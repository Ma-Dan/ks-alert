
PROTO_FILES=$(sort $(wildcard ./*.proto))

all: $(PROTO_FILES) Makefile
	make executor

executor:
	# generate the gRPC code
	protoc --go_out=plugins=grpc:../pb/ ${PROTO_FILES}
