package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"course/internal/repository/pg/entity"
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

func (s gRPCServerStruct) GetAll(ctx context.Context, req *AllCourseRequest) (*AllCourseResponse, error) {
	courses, err := s.CourseApp.AllCourse(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		return &AllCourseResponse{}, status.Error(codes.Internal, err.Error())
	}

	pbCourses := transform(courses, func(course *entity.Course) *CourseResponse {
		return &CourseResponse{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CreatedDate: timestamppb.New(course.CreatedAt),
		}
	})

	return &AllCourseResponse{
		Courses: pbCourses,
	}, nil
}

func transform[T1, T2 any](items []T1, f func(item T1) T2) []T2 {
	if len(items) == 0 {
		return []T2{}
	}

	result := make([]T2, 0, len(items))

	for _, item := range items {
		result = append(result, f(item))
	}

	return result
}
