package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s gRPCServerStruct) AddCourse(ctx context.Context, req *CourseRequest) (*CourseResponse, error) {
	course, err := s.CourseApp.AddCourse(ctx, req.GetName(), req.GetDescription())
	if err != nil {
		return &CourseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &CourseResponse{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CreatedDate: timestamppb.New(course.CreatedAt),
	}, nil
}
