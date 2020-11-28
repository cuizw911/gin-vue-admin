package request

import "gin-vue-admin/model"

type TblCourseInfoSearch struct{
    model.TblCourseInfo
    PageInfo
}