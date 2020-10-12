package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	sv "github.com/y-harashima/grpc-sample/server"
	pb "github.com/y-harashima/grpc-sample/proto"
)

func main() {
	
	port, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	zap, _ := zap.NewProduction()
	zap_opt := grpc_zap.WithLevels(
		func(c codes.Code) zapcore.Level {
			var l zapcore.Level
			switch c {
			case codes.OK:
				l = zapcore.InfoLevel
				
			case codes.Internal:
				l = zapcore.ErrorLevel

			default:
				l = zapcore.DebugLevel
			}
			return l
		},
	)
	grpc := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap, zap_opt),
		),
	)
	

	server := &sv.Server{}
	pb.RegisterStudentServer(grpc, server)
	reflection.Register(grpc)
	log.Println("Server process starting...")
	if err := grpc.Serve(port); err != nil {
		log.Fatal(err)
	}
}