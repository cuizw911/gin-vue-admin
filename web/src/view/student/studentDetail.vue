<template>
<div>
    <DescriptionList title="退款申请" col="6" :content="data">
        <Description term="取货单号">1000000000</Description>
        <Description term="状态">已取货</Description>
        <Description term="销售单号">1234123421</Description>
        <Description term="子订单">3214321432</Description>
    </DescriptionList>
</div>
</template>

<script>
import {
    createTblStudents,
    updateTblStudents,
    findTblStudents
} from "@/api/students";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
import DescriptionList from "@/components/description"; // 引入组件
export default {
  name: "StudentsDetail",
  mixins: [infoList],
  components: {DescriptionList},
  data() {
    return {
      type: "",
      formData: {
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
  async created () {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findTblStudents({ ID: this.$route.query.id })
      console.log(res)
      if (res.code == 0) {
        this.formData = res.data.restudents
      }
    }
  }
};
</script>

<style>
</style>
