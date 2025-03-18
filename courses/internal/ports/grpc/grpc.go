package grpc

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"course/internal/app"
)

type IServer interface {
	CourseServiceServer
}

type gRPCServerStruct struct {
	UnimplementedCourseServiceServer // Вызывает панику, если метод не реализован
	CourseApp                        app.ICourseApp
}

func (s gRPCServerStruct) Ping(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	log.Println("ping course server")

	return &emptypb.Empty{}, nil
}

func NewService(app app.ICourseApp) IServer {
	return &gRPCServerStruct{
		CourseApp: app,
	}
}
