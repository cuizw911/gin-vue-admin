// 自动生成模板TblTeacherInfo
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type TblTeacherInfo struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(20);size:20;"`
      Course  string `json:"course" form:"course" gorm:"column:course;comment:课程;type:char(2);size:2;"`
      Grade  string `json:"grade" form:"grade" gorm:"column:grade;comment:年级;type:char(2);size:2;"`
      EntryDate  string `json:"entryDate" form:"entryDate" gorm:"column:entry_date;comment:入职时间;type:varchar(10);size:10;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:电话;type:varchar(20);size:20;"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:家庭住址;type:varchar(80);size:80;"`
      Condition  string `json:"condition" form:"condition" gorm:"column:condition;comment:健康状态;type:char(2);size:2;"`
      Level  string `json:"level" form:"level" gorm:"column:level;comment:星级;type:char(2);size:2;"`
      Category  string `json:"category" form:"category" gorm:"column:category;comment:类型（全职、兼职、外校）;type:char(2);size:2;"`
}


func (TblTeacherInfo) TableName() string {
  return "tbl_teacher_info"
}
