<template>
  <div style="width: 100%">
    <el-card> 入驻商家列表 </el-card>
    <el-card style="margin-top: 20px">
      <el-form
        :inline="true"
        :model="formInline"
        ref="formInline"
        :rules="rules"
      >
        <el-form-item label="搜索内容" prop="selectInfo">
          <el-input
            style="width: 250px"
            @focus="clearInput"
            v-model="formInline.selectInfo"
            placeholder="联系人/电话/地址"
          ></el-input>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card style="margin-top: 20px">
      <el-table
        :data="
          ordersData.slice((currentPage - 1) * pageSize, currentPage * pageSize)
        "
        row-key="id"
        border
        style="width: 1980px"
      >
        <el-table-column label="序号" width="50">
          <template slot-scope="scope">{{ scope.$index + 1 }}</template>
        </el-table-column>
        <el-table-column label="状态" width="115">
          <template slot-scope="scope">
            <span v-if="scope.row.state == 0">未付款</span>
            <span v-else-if="scope.row.state == 1">已付款,未送达</span>
            <span v-else>已完成</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="create_time"
          label="创建时间"
          width="110"
          sortable
        >
        </el-table-column>
        <el-table-column
          prop="order_code"
          label="订单编号"
          width="110"
          sortable
        >
        </el-table-column>
        <el-table-column prop="liaison" label="联系人" width="85">
        </el-table-column>
        <el-table-column prop="gender" label="性别" width="55">
        </el-table-column>
        <el-table-column prop="liaison_phone" label="联系电话" width="115">
        </el-table-column>
        <el-table-column prop="address" label="收货地址" width="225">
        </el-table-column>
        <el-table-column prop="metre" label="距离(KM)" width="115" sortable>
        </el-table-column>
        <el-table-column
          prop="delivery_datetime"
          label="预计送达"
          width="115"
          sortable
        >
        </el-table-column>
        <el-table-column
          prop="goods_price"
          label="商品价格"
          width="115"
          sortable
        >
        </el-table-column>
        <el-table-column
          prop="delivery_price"
          label="配送价格"
          width="115"
          sortable
        >
        </el-table-column>
        <el-table-column label="签约商家" width="85">
          <template slot-scope="scope">
            <span v-if="scope.row.is_union == 1">签约</span>
            <span v-else>未签</span>
          </template>
        </el-table-column>
        <el-table-column prop="total_price" label="总价" width="115" sortable>
        </el-table-column>
        <el-table-column label="下单用户" width="115">
          <template slot-scope="scope">
            <span v-if="scope.row.user_name">{{ scope.row.user_name }}</span>
            <span v-else>{{ scope.row.user_phone }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="goods_info" label="商品信息" width="115">
          <template slot-scope="scope">
            <el-button type="text" @click="getSaleGoodsByOrderId(scope.row)"
              >点击查看</el-button
            >
            <el-dialog title="订单商品信息" :visible.sync="dialogTableVisible" style="width: 1900px">
              <el-table :data="saleGoodsData"  show-summary >
                <el-table-column label="序号" width="50">
                  <template slot-scope="scope">{{ scope.$index + 1 }}</template>
                </el-table-column>
                <el-table-column
                  property="sale_code"
                  label="识别码"
                  width="100"
                ></el-table-column>
                <el-table-column
                  property="shop_name"
                  label="商家名称"
                  width="160"
                ></el-table-column>
                <el-table-column
                  property="id"
                  label="商品ID"
                  width="65"
                ></el-table-column>
                <el-table-column
                  property="good_name"
                  label="商品名称"
                  width="115"
                ></el-table-column>
                <el-table-column
                  property="good_price"
                  label="单价"
                  width="85"
                ></el-table-column>
                <el-table-column
                  property="good_count"
                  label="数量"
                  width="85"
                ></el-table-column>
                <el-table-column
                  property="exchange"
                  label="调换货"
                  width="85"
                ></el-table-column>
                <el-table-column
                  property="total_price"
                  label="小计"
                  width="105"
                ></el-table-column>
              </el-table>
            </el-dialog>
          </template>
        </el-table-column>
        <el-table-column prop="invoice" label="发票" width="65">
        </el-table-column>
        <el-table-column prop="remark" label="备注" width="115">
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
import { reqOrderList, reqSaleGoodsByOrderId } from "../../../api/index";
export default {
  name: "orderlist",
  data() {
    return {
      static_url: global.static_url,
      manager_token: global.manager_token,
      ordersData: [],
      currentPage: 1,
      pageSize: 5,
      pageTotal: 0,
      saleGoodsData: [],
      dialogTableVisible: false,
      formInline: {
        selectInfo: "联系人/电话/地址",
        catesName: "默认",
        status: "默认",
        union: "默认",
        position: "默认",
        give: "默认",
        salePromote: "默认",
      },
    };
  },
  mounted() {
    this.getOrders();
  },
  methods: {
    clearInput() {
      this.resetSearch();
    },
    // 重置查询
    resetSearch() {
      this.formInline.selectInfo = "";
      this.formInline.catesName = "默认";
      this.formInline.status = "默认";
      this.formInline.union = "默认";
      this.formInline.position = "默认";
      (this.formInline.give = "默认"),
        (this.formInline.salePromote = "默认"),
        (this.tempArray = []);
      this.suppData = this.saveSuppData;
      this.currentPage = 1;
      this.pageTotal = this.suppData.length;
    },
    async getOrders() {
      if (!this.manager_token) {
        this.manager_token = localStorage.getItem("manager_token");
      }
      const result = await reqOrderList(this.manager_token);
      if (result.code === 0) {
        this.ordersData = result.data;
        this.pageTotal = this.ordersData.length;
      }
      console.log(result);
    },
    async getSaleGoodsByOrderId(row) {
      this.saleGoodsData = [];
      this.$delete(row, "create_time");
      this.$delete(row, "update_time");
      this.$delete(row, "shop_info");
      const result = await reqSaleGoodsByOrderId(this.manager_token, row);
      if (result.code === 0) {
        result.data.forEach((element) => {
          if (element.good_price) {
            var price = element.good_price.toFixed(2);
            element.good_price = price;
            element.total_price = (
              element.good_price * element.good_count
            ).toFixed(2);
          }
        });
        console.log(result.data);
        this.dialogTableVisible = true;
        this.saleGoodsData = result.data;
      }
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