API_COMMON_PROTOS?=$(shell go env GOPATH)/src/github.com/googleapis/api-common-protos
GOOGLEAPIS_PROTOS?=$(shell go env GOPATH)/src/github.com/googleapis/googleapis

protos:
	rm -rf bqconn/*
	protoc -I ${API_COMMON_PROTOS} -I ${GOOGLEAPIS_PROTOS} --go_gapic_out .//bqconn --go_gapic_opt 'go-gapic-package=github.com/tmc/mkbigqueryconn/bqconn;bqconn' google/cloud/bigquery/connection/v1beta1/connection.proto
	mv bqconn/github.com/tmc/mkbigqueryconn/bqconn/*.go ./bqconn/
	protoc -I ${GOOGLEAPIS_PROTOS} --go_cli_out=cmd/bqconn --go_cli_opt "root=bqconn" --go_cli_opt "gapic=github.com/tmc/mkbigqueryconn/bqconn" google/cloud/bigquery/connection/v1beta1/connection.proto
