
PROTO_FILES=$(sort $(wildcard ./*.proto))

all: $(PROTO_FILES) Makefile
	make executor

executor:
	# generate the gRPC code
	protoc3.6  --go_out=plugins=grpc:../pb/  alertconfig.proto alerthistory.proto alertrule.proto ent.proto error.proto product.proto receiver.proto repeatsend.proto resource.proto resourcetype.proto resourceuritmpl.proto silence.proto suggestion.proto severity.proto executor.proto
	# protoc3.6 --go_out=plugins=grpc:../pb/ ${PROTO_FILES}
