package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	pb "github.com/y-harashima/grpc-sample/proto"
)

type Server struct{}

func (s *Server) Get(ctx context.Context, req *pb.StudentRequest) (*pb.StudentResponse, error) {
	if req.Id == 1 {
		res := &pb.StudentResponse{
			Id: 1,
			Name: "Taro",
			Age: 11,
			School: &pb.School{
				Id: 1,
				Name: "ABC school",
				Grade: "5th",
			},
		}
		log := map[string]interface{}{
			"name": res.Name, 
			"age": res.Age, 
			"school_name": res.School.Name,
			"school_grade": res.School.Grade,
		}
		grpc_ctxtags.Extract(ctx).Set("data", log)
		return res, nil
	} else {
		grpc_ctxtags.Extract(ctx).Set("request_id", req.Id)
		return nil, status.Errorf(codes.Internal, "No students found.")
	}
}