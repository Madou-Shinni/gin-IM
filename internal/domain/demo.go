package domain

import (
	"gin-IM/pkg/model"
	"gin-IM/pkg/request"
)

type Demo struct {
	model.Model
	F *string
}

type PageDemoSearch struct {
	Demo
	request.PageSearch
}

func (Demo) TableName() string {
	return "demo"
}
