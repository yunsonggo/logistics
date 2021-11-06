<template>
  <div style="width: 100%">
    <el-card> 优惠活动列表 </el-card>
    <el-card style="margin-top: 20px">
      <el-table
        :data="
          promoteData.slice(
            (currentPage - 1) * pageSize,
            currentPage * pageSize
          )
        "
        row-key="id"
        border
        style="width: 650px"
      >
        <el-table-column label="序号" width="50">
          <template slot-scope="scope">{{ scope.$index + 1 }}</template>
        </el-table-column>
        <el-table-column prop="name" label="名称" width="115"></el-table-column>
        <el-table-column
          prop="seat"
          label="排位"
          width="85"
          sortable
        ></el-table-column>
        <el-table-column label="状态" width="65">
          <template slot-scope="scope">
            <span v-if="scope.row.state == 1">正常</span>
            <span v-else>关闭</span>
          </template>
        </el-table-column>
        <el-table-column label="描述" prop="desc">
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              type="primary"
              plain
              size="mini"
              @click="handleEdit(scope.$index, scope.row)"
              >编辑
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
import { reqSalePromote } from "../../../api/index";
export default {
  name: "salepromotelist",
  data() {
    return {
      static_url: global.static_url,
      manager_token: global.manager_token,
      promoteData: [],
      currentPage: 1,
      pageSize: 5,
      pageTotal: 0,
    };
  },
  mounted() {
    this.getPromoteData();
    console.log(this.promoteData);
  },
  methods: {
    // 获取服务数据
    getPromoteData() {
      if (this.$store.getters.salePromote) {
        this.promoteData = this.$store.getters.salePromote;
        this.pageTotal = this.promoteData.length;
      } else {
        this.getPromote();
      }
    },
    async getPromote() {
      if (!this.manager_token) {
        this.manager_token = localStorage.getItem("manager_token");
      }
      const result = await reqSalePromote(this.manager_token);
      if (result.code === 0) {
        this.$store.dispatch("setSalePromote", result.data);
        this.promoteData = result.data;
        this.pageTotal = this.promoteData.length;
      }
    },
    handleEdit(index, row) {
      this.$router.push({
        name: "salepromoteedit",
        params: { promoteInfo: row },
      });
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