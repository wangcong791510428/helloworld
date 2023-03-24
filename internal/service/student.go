package service

import (
	"context"
	"fmt"
	"helloworld/internal/biz"

	pb "helloworld/api/student/v1"
)

type StudentService struct {
	pb.UnimplementedStudentServer

	uc *biz.StudentUsercase
}

func NewStudentService(uc *biz.StudentUsercase) *StudentService {
	return &StudentService{
		uc: uc,
	}
}

// 可以引用biz层中的方法
func (s *StudentService) CallStudent(ctx context.Context, req *pb.StudentRequest) (*pb.StudentReply, error) {
	msg := fmt.Sprintf("姓名: %s, 年龄: %d", req.Name, req.Age)
	fmt.Println("222222222222222")

	return &pb.StudentReply{
		Message: msg,
	}, nil
}
