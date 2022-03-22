package model

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID          int    `gorm:"primaryKey;column:id;" json:"id"`
	Name        string `gorm:"column:name;type:varchar(100);default:" json:"name" form:"name"`
	Description string `gorm:"column:description;type:varchar(200);default:" json:"description" form:"description"` //描述
	Color       string `gorm:"column:color;type:varchar(100);default:" json:"color" form:"color"`
	Icon        string `gorm:"column:icon;type:varchar(200);default:" json:"icon" form:"icon"`
	TbName      string `gorm:"column:tb_name;type:varchar(20);default:" json:"tb_name" form:"tb_name"`                    //操作表名
	ListUrl     string `gorm:"column:list_url;type:varchar(200);default:" json:"list_url" form:"list_url"`                //列表地址
	ContentUrl  string `gorm:"column:content_url;type:varchar(200);default:" json:"content_url" form:"content_url"`       //详情地址
	SaveFormUrl string `gorm:"column:save_form_url;type:varchar(200);default:" json:"save_form_url" form:"save_form_url"` //编辑地址
	DelUrl      string `gorm:"column:del_url;type:varchar(200);default:" json:"del_url" form:"del_url"`                   //删除地址
	CreatedAt   time.Time
	AddTime     int32 `gorm:"column:add_time;type:int(10);default:0" json:"addTime"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (i *Category) TableName() string {
	return "category"
}
func NewCategory() *Category {
	return new(Category)
}
