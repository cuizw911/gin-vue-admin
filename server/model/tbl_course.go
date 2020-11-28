// 自动生成模板TblCourseInfo
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type TblCourseInfo struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(50);size:50;"`
      Category  string `json:"category" form:"category" gorm:"column:category;comment:类型;type:char(2);size:2;"`
      LessonHour  int `json:"lessonHour" form:"lessonHour" gorm:"column:lesson_hour;comment:课时;type:int;size:11;"`
      CourseFee  float64 `json:"courseFee" form:"courseFee" gorm:"column:course_fee;comment:课程费用;type:double;size:11;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:;type:varchar(255);size:255;"`
}


func (TblCourseInfo) TableName() string {
  return "tbl_course_info"
}
