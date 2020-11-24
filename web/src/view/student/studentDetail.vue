<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
               <el-form-item label="ID:"><el-input v-model.number="formData.id" clearable placeholder="请输入"></el-input>
               </el-form-item>

               <el-form-item label="名字:">
                   <el-input v-model="formData.studentName" clearable placeholder="请输入" ></el-input>
               </el-form-item>

               <el-form-item label="年龄:"><el-input v-model.number="formData.age" clearable placeholder="请输入"></el-input>
               </el-form-item>

               <el-form-item label="性别:">
                   <el-input v-model="formData.gender" clearable placeholder="请输入" ></el-input>
               </el-form-item>

               <el-form-item label="生日:">
                   <el-input v-model="formData.birthday" clearable placeholder="请输入" ></el-input>
               </el-form-item>

               <el-form-item label="家长名字:">
                   <el-input v-model="formData.parentName" clearable placeholder="请输入" ></el-input>
               </el-form-item>

               <el-form-item label="家长电话:">
                   <el-input v-model="formData.parentPhone" clearable placeholder="请输入" ></el-input>
               </el-form-item>

               <el-form-item label="所属班级:">
                   <el-input v-model="formData.belongClass" clearable placeholder="请输入" ></el-input>
               </el-form-item>

               <el-form-item label="入学时间:">
                   <el-input v-model="formData.admissionDate" clearable placeholder="请输入" ></el-input>
               </el-form-item>

               <el-form-item label="备注:">
                   <el-input v-model="formData.remark" clearable placeholder="请输入" ></el-input>
               </el-form-item>
               <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createTblStudents,
    updateTblStudents,
    findTblStudents
} from "@/api/students";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "TblStudents",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            id:0,
            studentName:"",
            age:0,
            gender:"",
            birthday:"",
            parentName:"",
            parentPhone:"",
            belongClass:"",
            admissionDate:"",
            remark:"",

      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createTblStudents(this.formData);
          break;
        case "update":
          res = await updateTblStudents(this.formData);
          break;
        default:
          res = await createTblStudents(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findTblStudents({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.restudents
       this.type == "update"
     }
    }else{
     this.type == "create"
   }

}
};
</script>

<style>
</style>
