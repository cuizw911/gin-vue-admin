<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="教师名字">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>      
        <el-form-item label="年级">
          <el-input placeholder="搜索条件" v-model="searchInfo.grade"></el-input>
        </el-form-item>                
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增教师信息</el-button>
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
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="教师名字" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="课程" prop="course" width="120"></el-table-column> 
    
    <el-table-column label="年级" prop="grade" width="120"></el-table-column> 
    
    <el-table-column label="入职时间" prop="entryDate" width="120"></el-table-column> 
    
    <el-table-column label="电话" prop="phone" width="120"></el-table-column> 
    
    <el-table-column label="家庭住址" prop="address" width="120"></el-table-column> 
    
    <el-table-column label="健康状况" prop="condition" width="120"></el-table-column> 
    
    <el-table-column label="星级" prop="level" width="120"></el-table-column> 
    
    <el-table-column label="类别" prop="category" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateTblTeacherInfo(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteTblTeacherInfo(scope.row)">确定</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="教师名字:">
            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="课程:">
            <el-input v-model="formData.course" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="年级:">
            <el-input v-model="formData.grade" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="入职时间:">
            <el-input v-model="formData.entryDate" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="电话:">
            <el-input v-model="formData.phone" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="家庭住址:">
            <el-input v-model="formData.address" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="健康状况:">
            <el-input v-model="formData.condition" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="星级:">
            <el-input v-model="formData.level" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="类别:">
            <el-input v-model="formData.category" clearable placeholder="请输入" ></el-input>
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
    createTblTeacherInfo,
    deleteTblTeacherInfo,
    deleteTblTeacherInfoByIds,
    updateTblTeacherInfo,
    findTblTeacherInfo,
    getTblTeacherInfoList
} from "@/api/teacher";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "TblTeacherInfo",
  mixins: [infoList],
  data() {
    return {
      listApi: getTblTeacherInfoList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            name:"",
            course:"",
            grade:"",
            entryDate:"",
            phone:"",
            address:"",
            condition:"",
            level:"",
            category:"",
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10             
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
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
        const res = await deleteTblTeacherInfoByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateTblTeacherInfo(row) {
      const res = await findTblTeacherInfo({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.retblTeacherInfo;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          name:"",
          course:"",
          grade:"",
          entryDate:"",
          phone:"",
          address:"",
          condition:"",
          level:"",
          category:"",
          
      };
    },
    async deleteTblTeacherInfo(row) {
      this.visible = false;
      const res = await deleteTblTeacherInfo({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createTblTeacherInfo(this.formData);
          break;
        case "update":
          res = await updateTblTeacherInfo(this.formData);
          break;
        default:
          res = await createTblTeacherInfo(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>