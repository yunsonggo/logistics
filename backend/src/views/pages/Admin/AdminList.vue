<template>
  <div style="width: 100%">
    <el-card> 工作人员列表 </el-card>
    <el-card style="margin-top: 20px">
      <el-table
        :data="
          adminData.slice((currentPage - 1) * pageSize, currentPage * pageSize)
        "
        row-key="id"
        border
        style="width: 1024px"
      >
        <el-table-column prop="id" label="ID" width="60"></el-table-column>
        <el-table-column
          prop="username"
          label="注册名称"
          width="150"
        ></el-table-column>
        <el-table-column
          prop="usernick"
          label="昵称"
          width="150"
        ></el-table-column>
        <el-table-column prop="user_icon" label="头像" width="150">
            <template slot-scope="scope">
                <el-avatar v-if="scope.row.user_icon" shape="square" :size="58"  :src="static_url + scope.row.user_icon"></el-avatar>
            </template>
        </el-table-column>
        <el-table-column
          prop="user_gender"
          label="性别"
          width="50"
        ></el-table-column>
        <el-table-column
          prop="phone"
          label="联系电话"
          width="150"
        ></el-table-column>
        <el-table-column prop="power" label="权限" width="150">
          <template slot-scope="scope">
            <span v-if="scope.row.power == '7'">管理员</span>
            <span v-else-if="scope.row.power == '77'">系统管理员</span>
            <span v-else>超级管理员</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
              >编辑
            </el-button>
            <el-button size="mini" type="danger" @click="handleDelete(scope.row)"
              >删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        background
        :current-page="currentPage"
        :page-sizes="[2, 5, 10, 15]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="pageTotal"
      >
      </el-pagination>
    </el-card>
  </div>
</template>

<script>
import { reqAdminList,reqDelAdminById } from "@/api/index";
import global from '@/components/global'

export default {
  name: "adminlist",
  data() {
    return {
      static_url: global.static_url,
      adminData: [],
      currentPage: 1,
      pageSize: 5,
      pageTotal: 0,
    };
  },
  mounted() {
    this.getAdminList();
  },
  methods: {
    // 获取 人员数据
    async getAdminList() {
      let manager_token = localStorage.getItem("manager_token");
      const result = await reqAdminList(manager_token);
      if (result.code === 0) {
        this.adminData = result.data;
        this.pageTotal = result.data.length;
        //console.log(this.adminData);
      } else {
        this.adminData = [];
        this.$message.error(result.msg);
      }
      return;
    },
    // 编辑
    handleEdit(index, row) {
      this.$router.push({name:'adminedit',params:{adminInfo:row}})
    },
    handleDelete(row) {
      if (row.power == "777") {
        this.$message.error('无法删除超级管理员');
        return
      } else {
        this.$confirm("此操作将永久删除该工作账号, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          //调用 删除方法
          this.$delete(row,'create_time')
          this.removeAdminById(row).then(res => {
            console.log(res);
            if (res.code === 0) {
              this.$message({
                type: "success",
                message: "删除管理员" + row.username + "成功",
              });
              this.$router.push('/yunsong/home/index')
              return
            } else {
              this.$message.error('删除管理员失败');
              return
            }
          })
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
      }
    },
    // 实现删除方法
    async removeAdminById(row) {
      let manager_token = localStorage.getItem("manager_token");
      const result = await reqDelAdminById(manager_token,row)
      return result
    },
    handleSizeChange(size) {
      this.pageSize = size;
    },
    handleCurrentChange(currentPage) {
      this.currentPage = currentPage;
    },
  },
};
</script>