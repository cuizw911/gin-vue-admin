package request

import "gin-vue-admin/model"

type TblStudentsSearch struct{
    model.TblStudents
    PageInfo
}