<template>
  <div style="width: 100%">
    <el-card> 客户列表 </el-card>
    <el-card style="margin-top: 20px">
      <el-table
        :data="
          userData.slice((currentPage - 1) * pageSize, currentPage * pageSize)
        "
        row-key="id"
        border
        style="width: 950px"
      >
        <el-table-column label="序号" width="50">
          <template slot-scope="scope">{{ scope.$index + 1 }}</template>
        </el-table-column>
        <el-table-column
          prop="phone"
          label="用户名(手机)"
          width="165"
        ></el-table-column>
        <el-table-column
          prop="username"
          label="用户昵称"
          width="155"
        ></el-table-column>
        <el-table-column
          prop="user_gender"
          label="性别"
          width="65"
        ></el-table-column>
        <el-table-column label="状态" width="65">
          <template slot-scope="scope">
            <span v-if="scope.row.is_del == 1">正常</span>
            <span v-else>拉黑</span>
          </template>
        </el-table-column>
        <el-table-column prop="icon" label="图标" width="95">
          <template slot-scope="scope">
            <el-avatar
              v-if="scope.row.icon"
              shape="square"
              :size="58"
              :src="static_url + scope.row.icon"
            ></el-avatar>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              type="primary"
              plain
              size="mini"
              @click="handleUserAddress(scope.$index, scope.row)"
              >收货地址
            </el-button>
            <el-button
              type="primary"
              plain
              size="mini"
              @click="handleUserOrders(scope.$index, scope.row)"
              >用户订单
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        background
        :current-page="currentPage"
        :page-sizes="[5, 10, 15, 20]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="pageTotal"
      >
      </el-pagination>
    </el-card>
  </div>
</template>

<script>
import global from "@/components/global";
import { reqUserList } from "../../../api/index";
export default {
  name: "userlist",
  data() {
    return {
      static_url: global.static_url,
      manager_token: global.manager_token,
      userData: [],
      currentPage: 1,
      pageSize: 5,
      pageTotal: 0,
    };
  },
  mounted() {
    this.getUserData();
  },
  methods: {
    async getUserData() {
      if (!this.manager_token) {
        this.manager_token = localStorage.getItem("manager_token");
      }
      const result = await reqUserList(this.manager_token);
      if (result.code === 0) {
        this.userData = result.data;
        this.pageTotal = this.userData.length;
        console.log(this.userData);
      }
    },
    handleSizeChange(size) {
      this.pageSize = size;
    },
    handleCurrentChange(currentPage) {
      this.currentPage = currentPage;
    },
    handleUserAddress(index, row) {
    //   this.$router.push({ name: "userlist", params: { userInfo: row } });
        console.log(row);
    },
    handleUserOrders(index, row) {
        console.log(row);
    }
  },
};
</script>