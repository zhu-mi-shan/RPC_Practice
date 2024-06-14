package main

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"strings"
	gateway "student/kitex_gen/gateway"
	"student/module"
)

// BizServiceImpl implements the last service interface defined in the IDL.
type BizServiceImpl struct{ db *gorm.DB }

// Register implements the BizServiceImpl interface.
func (s *BizServiceImpl) Register(ctx context.Context, req *gateway.BizRequest) (resp *gateway.BizResponse, err error) {
	// TODO: Your code here...
	log.Println(req)
	result := s.db.Table("students").Create(student2Model(req.Student))
	if result.Error != nil {
		return nil, err
	}
	resp = &gateway.BizResponse{}
	successValue := true
	resp.Success = &successValue

	// 设置 Message 字段的值
	messageValue := "This is a success message"
	resp.Message = &messageValue
	log.Println("register success")
	return
}

// Query implements the BizServiceImpl interface.
func (s *BizServiceImpl) Query(ctx context.Context, req *gateway.BizRequest) (resp *gateway.BizResponse, err error) {
	// TODO: Your code here...
	log.Println(req.ItemId)
	var stuRes module.Student
	//log.Println(req.Id)
	log.Println("query student")
	log.Println(req)
	//log.Println(&req.Id)
	err = s.db.Table("students").First(&stuRes, req.ItemId).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp = &gateway.BizResponse{}
	successValue := true
	resp.Success = &successValue

	// 设置 Message 字段的值
	messageValue := "This is a success message"
	resp.Message = &messageValue
	resp.Student = model2Student(&stuRes)
	log.Println("query success")
	return
}

// InitDB eg: 初始化db，注意服务启动时初始化
func (s *BizServiceImpl) InitDB() {
	db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&module.Student{})
	if err != nil {
		panic(err)
		return
	}
	s.db = db
}

func student2Model(student *gateway.Student) *module.Student {
	return &module.Student{
		Id:             student.Id,
		Name:           student.Name,
		Email:          strings.Join(student.Email, ","),
		CollegeName:    student.College.Name,
		CollegeAddress: student.College.Address,
	}
}

func model2Student(student *module.Student) *gateway.Student {
	return &gateway.Student{
		Id:      student.Id,
		Name:    student.Name,
		Email:   strings.Split(student.Email, ","),
		College: &gateway.College{Name: student.CollegeName, Address: student.CollegeAddress},
	}
}
