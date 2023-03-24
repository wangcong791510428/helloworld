package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
)

type studentRepo struct {
	data *Data
	log *log.Helper
}

func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *studentRepo) Save(ctx context.Context, g *biz.Student) (*biz.Student, error) {
	fmt.Println("111111111111111111111111111111")
	return g, nil
}