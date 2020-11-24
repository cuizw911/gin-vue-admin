<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item>
                    <el-button @click="onSubmit" type="primary">查询</el-button>
                </el-form-item>
                <el-form-item>
                    <el-button @click="openDialog" type="primary">新增学生信息表</el-button>
                </el-form-item>
                <el-form-item>
                    <el-popover placement="top" v-model="deleteVisible" width="160">
                        <p>确定要删除吗？</p>
                        <div style="text-align: right; margin: 0">
                            <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                            <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
                        </div>
                        <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
                    </el-popover>
                </el-form-item>
            </el-form>
        </div>
        <el-table
                :data="tableData"
                @selection-change="handleSelectionChange"
                border
                ref="multipleTable"
                stripe
                style="width: 100%"
                tooltip-effect="dark"
        >
            <el-table-column type="selection" width="55"></el-table-column>

            <el-table-column label="ID" prop="ID" width="60"></el-table-column>

            <el-table-column label="名字" prop="name" width="120"></el-table-column>

            <el-table-column label="年龄" prop="age" width="120"></el-table-column>

            <el-table-column label="性别" prop="gender" width="120" :formatter="genderFmt"></el-table-column>

            <el-table-column label="生日" prop="birthday" width="120"></el-table-column>

            <el-table-column label="家长名字" prop="parentName" width="120"></el-table-column>

            <el-table-column label="家长电话" prop="parentPhone" width="120"></el-table-column>

            <el-table-column label="所属班级" prop="belongClass" width="120" :formatter="classFmt"></el-table-column>

            <el-table-column label="入学时间" prop="admissionDate" width="120"></el-table-column>

<!--            <el-table-column label="日期" width="180">-->
<!--                <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>-->
<!--            </el-table-column>-->

<!--            <el-table-column label="备注" prop="remark" width="120"></el-table-column>-->

            <el-table-column label="按钮组">
                <template slot-scope="scope">
                    <el-button class="table-button" @click="updateTblStudents(scope.row)" size="small" type="primary"
                               icon="el-icon-edit">变更
                    </el-button>
                    <el-popover placement="top" width="160" v-model="scope.row.visible">
                        <p>确定要删除吗？</p>
                        <div style="text-align: right; margin: 0">
                            <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                            <el-button type="primary" size="mini" @click="deleteTblStudents(scope.row)">确定</el-button>
                        </div>
                        <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
                    </el-popover>
                </template>
            </el-table-column>
        </el-table>

        <el-pagination
                :current-page="page"
                :page-size="pageSize"
                :page-sizes="[10, 30, 50, 100]"
                :style="{float:'right',padding:'20px'}"
                :total="total"
                @current-change="handleCurrentChange"
                @size-change="handleSizeChange"
                layout="total, sizes, prev, pager, next, jumper"
        ></el-pagination>

        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增学生">
            <el-form :model="formData" :rules="rules" label-position="right" label-width="80px">
                <el-form-item label="学生名字:">
                    <el-col :span="8">
                        <el-input v-model="formData.name" clearable placeholder="请输入"></el-input>
                    </el-col>
                </el-form-item>

                <el-form-item label="学生年龄:">
                    <el-col :span="8">
                        <el-input v-model.number="formData.age" clearable placeholder="请输入"></el-input>
                    </el-col>
                </el-form-item>

                <el-form-item label="学生性别:">
                    <el-radio v-model="formData.gender" label="1">男生</el-radio>
                    <el-radio v-model="formData.gender" label="2">女生</el-radio>
                </el-form-item>

                <el-form-item label="学生生日:">
                    <div class="block">
                        <el-date-picker
                                v-model="formData.birthday"
                                type="date"
                                value-format="yyyy-MM-dd"
                                placeholder="选择日期">
                        </el-date-picker>
                    </div>
                </el-form-item>

                <el-form-item label="家长名字:">
                    <el-col :span="8">
                        <el-input v-model="formData.parentName" clearable placeholder="请输入家长名字"></el-input>
                    </el-col>
                </el-form-item>

                <el-form-item label="家长电话:">
                    <el-col :span="8">
                        <el-input v-model="formData.parentPhone" clearable placeholder="请输入电话" maxlength="11"></el-input>
                    </el-col>
                </el-form-item>

                <el-form-item label="所属班级:">
                    <el-select v-model="formData.belongClass" placeholder="请选择">
                        <el-option
                                v-for="item in classList"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>

                <el-form-item label="入学时间:">
                    <div class="block">
                        <el-date-picker
                                v-model="formData.admissionDate"
                                type="date"
                                value-format="yyyy-MM-dd"
                                placeholder="选择日期">
                        </el-date-picker>
                    </div>
                </el-form-item>

                <el-form-item label="备注:">
                    <el-input type="textarea" :rows="3" placeholder="请输入内容" v-model="formData.remark" maxlength="250" show-word-limit></el-input>
                </el-form-item>
            </el-form>
            <div class="dialog-footer" slot="footer">
                <el-button @click="closeDialog">取 消</el-button>
                <el-button @click="enterDialog" type="primary">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
import {
  createTblStudents,
  deleteTblStudents,
  deleteTblStudentsByIds,
  updateTblStudents,
  findTblStudents,
  getTblStudentsList
} from '@/api/students'  //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'

export default {
  name: 'TblStudents',
  mixins: [infoList],
  data () {
    return {
      listApi: getTblStudentsList,
      dialogFormVisible: false,
      visible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        id: 0,
        name: '',
        age: 0,
        gender: '1',
        birthday: '',
        parentName: '',
        parentPhone: '',
        belongClass: '',
        admissionDate: '',
        remark: '',
      },
      classList: [{
        value: "1",
        label: "小一班",
      }, {
        value: "2",
        label: "小二班",
      }, {
        value: "3",
        label: "小三班",
      }],
      rules: {
        'formData.name': [{ require: true, message: "请输入学生名字", trigger: "blur" }],
        parentName: [{ require: true, message: "请输入家长名字", trigger: "blur" }],
        parentPhone: [{ require: true, pattern: /^1[34578]\d{9}$/, message: "请输入正确的手机号", trigger: "change" }],
      }
    }
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    }
  },
  methods: {
    //条件搜索前端看此方法
    onSubmit () {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    handleSelectionChange (val) {
      this.multipleSelection = val
    },
    genderFmt(row) {
      if (row.gender === '1') {
        return "男生"
      } else {
        return "女生"
      }
    },
    classFmt(row) {
      let className
      this.classList.some(v => {
        if (row.belongClass === v.value) {
          className = v.label
          return true
        }
      })
      return className
    },
    async onDelete () {
      const ids = []
      if (this.multipleSelection.length == 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
      this.multipleSelection.map(item => {
        ids.push(item.ID)
      })
      const res = await deleteTblStudentsByIds({ ids })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateTblStudents (row) {
      const res = await findTblStudents({ ID: row.ID })
      this.type = 'update'
      if (res.code == 0) {
        this.formData = res.data.restudents
        this.dialogFormVisible = true
      }
    },
    closeDialog () {
      this.dialogFormVisible = false
      this.formData = {
        id: 0,
        name: '',
        age: 0,
        gender: '',
        birthday: '',
        parentName: '',
        parentPhone: '',
        belongClass: '',
        admissionDate: '',
        remark: '',

      }
    },
    async deleteTblStudents (row) {
      this.visible = false
      const res = await deleteTblStudents({ ID: row.ID })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        this.getTableData()
      }
    },
    async enterDialog () {
      let res
      // this.$refs.formData.validate(async (v) => {
      //
      //
      //
      // }
      //
      switch (this.type) {
        case 'create':
          res = await createTblStudents(this.formData)
          break
        case 'update':
          res = await updateTblStudents(this.formData)
          break
        default:
          res = await createTblStudents(this.formData)
          break
      }
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog () {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
  async created () {
    await this.getTableData()

  }
}
</script>

<style>
</style>
