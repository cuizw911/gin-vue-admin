package request

import "gin-vue-admin/model"

type TblTeacherInfoSearch struct{
    model.TblTeacherInfo
    PageInfo
}