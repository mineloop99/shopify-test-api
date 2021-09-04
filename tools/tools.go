package tools

import (
	_ "github.com/alta/insecure"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/lucas-clemente/quic-go"
	_ "github.com/lucas-clemente/quic-go/http3"
	_ "github.com/lucas-clemente/quic-go/logging"
	_ "github.com/lucas-clemente/quic-go/qlog"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
