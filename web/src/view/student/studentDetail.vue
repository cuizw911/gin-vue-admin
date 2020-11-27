<template>
    <div class="student-detail">
        <el-row :gutter="20">
            <el-card>
                <div slot="header" class="clearfix">
                    <span class="title">基本信息</span>
                </div>
                <el-col :span="6">
                    <div class="term">名字</div>
                    <div class="detail">{{studentDetail.name}}</div>
                </el-col>

                <el-col :span="6">
                    <div class="term">年龄</div>
                    <div class="detail">{{studentDetail.age}}</div>
                </el-col>

                <el-col :span="6">
                    <div class="term">性别</div>
                    <div class="detail">{{studentDetail.gender|formatGender}}</div>
                </el-col>

                <el-col :span="6">
                    <div class="term">生日</div>
                    <div class="detail">{{studentDetail.birthday}}</div>
                </el-col>

                <el-col :span="6">
                    <div class="term">家长名字</div>
                    <div class="detail">{{studentDetail.parentName}}</div>
                </el-col>

                <el-col :span="6">
                    <div class="term">联系电话</div>
                    <div class="detail">{{studentDetail.parentPhone}}</div>
                </el-col>
            </el-card>

            <el-card style="margin-top: 15px">
                <div slot="header" class="clearfix">
                    <span class="title">教学信息</span>
                </div>
                <el-col :span="6">
                    <div class="term">所属班级</div>
                    <div class="detail">{{ classFmt(studentDetail.belongClass) }}</div>
                </el-col>

                <el-col :span="6">
                    <div class="term">班主任</div>
                    <div class="detail">李老师</div>
                </el-col>

                <el-col :span="6">
                    <div class="term">入学时间</div>
                    <div class="detail">{{studentDetail.admissionDate}}</div>
                </el-col>
            </el-card>
        </el-row>
    </div>
</template>

<script>
import {
  findTblStudents
} from '@/api/students'
import { findSysDictionary } from '@/api/sysDictionary'
// import infoList from '@/mixins/infoList'

export default {
  name: 'StudentsDetail',
  // mixins: [infoList],
  data () {
    return {
      type: '',
      studentDetail: {},
      classList: [],
    }
  },
  filters: {
    formatGender: function (gender) {
      if (gender != null && gender != '') {
        return gender === '1' ? '男生' : '女生'
      } else {
        return ''
      }
    }
  },
  methods: {
    classFmt: function (c) {
      let className = '';
      this.classList.some(v => {
        if (c === v.value.toString()) {
          className = v.label
          return true
        }
      })
      return className
    },
    async getClassList () {
      const res = await findSysDictionary({ type: 'class' })
      if (res.code === 0) {
        const dicDetails = res.data.resysDictionary.sysDictionaryDetails
        let details = []
        dicDetails.forEach((item, index) => {
          details[index] = {
            value: item.value,
            label: item.label,
          }
        })
        this.classList = details
      }
    }
  },
  async created () {
    await this.getClassList()
    if (this.$route.params.id) {
      const res = await findTblStudents({ ID: this.$route.params.id })
      if (res.code == 0) {
        this.studentDetail = res.data.restudents
      }
    }
  }
}
</script>

<style lang="less" scoped>
    .student-detail {
        .title {
            font-weight: 700;
            font-size: 16px;
            line-height: 1.5;
            /*margin-bottom: 20px;*/
            color: rgba(0, 0, 0, 0.85);
        }

        .term {
            width: 100px;
            color: rgba(0, 0, 0, .85);
            font-weight: 400;
            font-size: 14px;
            line-height: 22px;
            padding-bottom: 16px;
            margin-right: 8px;
            white-space: nowrap;
            display: table-cell;

            &:after {
                content: ":";
                margin: 0 8px 0 2px;
                position: relative;
                top: -0.5px;
            }
        }

        .detail {
            font-size: 14px;
            line-height: 1.5;
            width: 100%;
            padding-bottom: 16px;
            color: rgba(0, 0, 0, 0.65);
            display: table-cell;
        }
    }
</style>
