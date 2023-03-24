package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Student struct {
	Name string
	Age int
}

// model层需要的东西
// 指定data层需要实现的方法
type StudentRepo interface {
	Save(ctx context.Context, student *Student) (*Student, error)
}

type StudentUsercase struct {
	repo StudentRepo
	log *log.Helper
}

func NewStudentUsecase(repo StudentRepo, logger log.Logger) *StudentUsercase {
	return &StudentUsercase{repo: repo, log: log.NewHelper(logger)}
}

//func (s StudentUsercase) Save() {
//	s.Repo.Save()
//}