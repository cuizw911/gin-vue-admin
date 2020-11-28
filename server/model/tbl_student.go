// 自动生成模板TblStudents
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type TblStudents struct {
	global.GVA_MODEL
	Name          string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(50);size:50;"`
	Age           int    `json:"age" form:"age" gorm:"column:age;comment:;type:int;size:4;"`
	Gender        string `json:"gender" form:"gender" gorm:"column:gender;comment:性别;type:char(1);size:1;"`
	Birthday      string `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日;type:varchar(20);size:20;"`
	ParentName    string `json:"parentName" form:"parentName" gorm:"column:parent_name;comment:家长名字;type:varchar(20);size:20;"`
	ParentPhone   string `json:"parentPhone" form:"parentPhone" gorm:"column:parent_phone;comment:家长电话;type:varchar(20);size:20;"`
	BelongClass   string `json:"belongClass" form:"belongClass" gorm:"column:belong_class;comment:所属班级;type:varchar(20);size:20;"`
	AdmissionDate string `json:"admissionDate" form:"admissionDate" gorm:"column:admission_date;comment:入学时间;type:varchar(20);size:20;"`
	Remark        string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;type:varchar(255);size:255;"`
}

func (TblStudents) TableName() string {
	return "tbl_students"
}
