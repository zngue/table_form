package service

import (
	"github.com/zngue/form_common/app/model"
	"github.com/zngue/go_helper/pkg"
)

type ICategory interface {
	Add(request *CategoryRequest) error
	Save(request *CategoryRequest) error
	List(request *CategoryRequest) (*[]model.Category, error)
	Detail(request *CategoryRequest) (*model.Category, error)
	Delete(request *CategoryRequest) error
	Count(request *CategoryRequest) (int64, error)
	GetModel() interface{}
}
type Category struct {
}
type CategoryRequest struct {
	pkg.CommonRequest
	ID    int    `form:"id" field:"id" where:"eq" default:"0"`
	Name  string `form:"name" field:"name" where:"like" default:""`
	IdArr string `form:"idArr" field:"id" where:"in" default:""`
}

func NewCategory() ICategory {
	return new(Category)
}
func (c *Category) GetModel() interface{} {
	return model.NewCategory()
}

// Add 添加
func (c *Category) Add(request *CategoryRequest) error {
	request.ReturnType = 3
	return pkg.MysqlConn.Model(c.GetModel()).Create(request.Data).Error
}

// Save 修改
func (c *Category) Save(request *CategoryRequest) error {
	request.ReturnType = 3
	return request.Init(pkg.MysqlConn.Model(c.GetModel()), *request).Updates(request.Data).Error
}
func (c *Category) List(request *CategoryRequest) (*[]model.Category, error) {
	var list []model.Category
	err := request.Init(pkg.MysqlConn.Model(c.GetModel()), *request).Find(&list).Error
	return &list, err
}

// Detail 详情
func (c *Category) Detail(request *CategoryRequest) (*model.Category, error) {
	var detail model.Category
	request.ReturnType = 3
	err := request.Init(pkg.MysqlConn.Model(c.GetModel()), *request).First(&detail).Error
	return &detail, err
}

// Delete 删除
func (c *Category) Delete(request *CategoryRequest) error {
	request.ReturnType = 3
	return request.Init(pkg.MysqlConn, *request).Delete(c.GetModel()).Error
}

// Count  数量
func (c *Category) Count(request *CategoryRequest) (int64, error) {
	request.ReturnType = 3
	var num int64
	err := request.Init(pkg.MysqlConn, *request).Model(c.GetModel()).Count(&num).Error
	return num, err
}
