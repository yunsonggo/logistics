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
        <el-form-item label="名称" prop="suppName">
          <el-input
            style="width: 250px"
            @focus="clearInput"
            v-model="formInline.suppName"
            placeholder="商家名称"
          ></el-input>
        </el-form-item>
        <el-form-item label="分类" prop="catesName">
          <el-select
            v-model="formInline.catesName"
            placeholder="默认"
            style="width: 110px"
          >
            <el-option label="默认" value="0"></el-option>
            <el-option
              v-for="(cate, index) in categroyData"
              :key="index"
              :label="cate.name"
              :value="cate.id"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select
            v-model="formInline.status"
            placeholder="默认"
            style="width: 85px"
          >
            <el-option label="默认" value="0"></el-option>
            <el-option
              v-for="(state, index) in statesData"
              :key="index"
              :label="state.name"
              :value="state.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="签约" prop="union">
          <el-select
            v-model="formInline.union"
            placeholder="默认"
            style="width: 85px"
          >
            <el-option label="默认" value="0"></el-option>
            <el-option
              v-for="(union, index) in unionData"
              :key="index"
              :label="union.name"
              :value="union.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="展示" prop="position">
          <el-select
            v-model="formInline.position"
            placeholder="默认"
            style="width: 85px"
          >
            <el-option label="默认" value="0"></el-option>
            <el-option
              v-for="(posi, index) in positionData"
              :key="index"
              :label="posi.name"
              :value="posi.key"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="服务" prop="give">
          <el-select
            v-model="formInline.give"
            placeholder="默认"
            style="width: 250px"
          >
            <el-option label="默认" value="0"></el-option>
            <el-option
              v-for="item in giveData"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="活动" prop="salePromote">
          <el-select v-model="formInline.salePromote" placeholder="默认">
            <el-option label="默认" value="0"></el-option>
            <el-option
              v-for="item in salePromoteData"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
          <el-button type="" @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card style="margin-top: 20px">
      <el-table
        :data="
          suppData.slice((currentPage - 1) * pageSize, currentPage * pageSize)
        "
        row-key="id"
        border
        style="width: 1248px"
      >
        <el-table-column label="序号" width="50">
          <template slot-scope="scope">{{ scope.$index + 1 }}</template>
        </el-table-column>
        <el-table-column
          prop="categroy"
          label="类别"
          width="85"
        ></el-table-column>
        <el-table-column
          prop="name"
          label="商家名称"
          width="210"
        ></el-table-column>
        <el-table-column prop="user_icon" label="LOGO" width="90">
          <template slot-scope="scope">
            <el-avatar
              v-if="scope.row.icon"
              shape="square"
              :size="58"
              :src="static_url + scope.row.icon"
            ></el-avatar>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="55">
          <template slot-scope="scope">
            <span v-if="scope.row.state == 1">正常</span>
            <span v-else>关闭</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="total_count"
          label="订单数"
          width="95"
          sortable
        ></el-table-column>
        <el-table-column label="签约" width="55">
          <template slot-scope="scope">
            <span v-if="scope.row.is_union == 2">未签</span>
            <span v-else style="font-weight: bold">已签</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="mobile"
          label="手机号码"
          width="115"
        ></el-table-column>
        <el-table-column label="商家商品" width="125">
          <template slot-scope="scope">
            <el-button
              type="primary"
              plain
              size="mini"
              @click="handleGoodsList(scope.$index, scope.row)"
              >商品列表
            </el-button>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button size="mini" @click="handleInfo(scope.$index, scope.row)"
              >详情
            </el-button>
            <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
              >编辑
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.row)"
              >关闭
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
    <el-drawer
      title="商家详情"
      ref="suppInfo"
      :visible.sync="drawer"
      :with-header="false"
      size="35%"
    >
      <el-card class="m-page-box">
        <div class="headClass">
          <el-backtop target=".headClass">
            <div
              style="
                 {
                  height: 100%;
                  width: 100%;
                  background-color: #f2f5f6;
                  box-shadow: 0 0 6px rgba(0, 0, 0, 0.12);
                  text-align: center;
                  line-height: 40px;
                  color: #1989fa;
                }
              "
            >
              up
            </div>
          </el-backtop>
          <div><h3>商家详情</h3></div>
          <div>
            <img :src="static_url + suppInfoData.banner" alt="" />
          </div>
          <div>
            <h4>{{ suppInfoData.name }}</h4>
            <p>商家类别:{{ suppInfoData.categroy }}</p>
            <p>商家公告:{{ suppInfoData.description }}</p>
            <p>营业时间:{{ suppInfoData.opening_hours }}</p>
            <p>订单总量:{{ suppInfoData.total_count }}</p>
          </div>
          <div>
            <h4>商家状态:</h4>
            <p v-if="suppInfoData.state == '1'">营业状态:正常</p>
            <p v-if="suppInfoData.state == '2'">营业状态:关闭</p>
            <p v-if="suppInfoData.is_union == '1'">是否签约:签约</p>
            <p v-if="suppInfoData.is_union == '2'">是否签约:未签</p>
            <p v-if="suppInfoData.create_time">
              注册时间:{{ suppInfoData.create_time }}
            </p>
          </div>
          <div>
            <h4>商家展示:</h4>
            <p v-if="suppInfoData.is_main == '1'">首页展示</p>
            <p v-if="suppInfoData.is_top == '1'">置顶展示</p>
            <p v-if="suppInfoData.is_new == '1'">新店展示</p>
            <p v-if="suppInfoData.is_hot == '1'">热卖展示</p>
          </div>
          <div>
            <h4>配送相关:</h4>
            <p>起送数量:{{ suppInfoData.min_num }}件</p>
            <p>配送价格:{{ suppInfoData.delivery_price }}</p>
            <p>超出运费:{{ suppInfoData.over_price }}</p>
          </div>
          <div>
            <h4>参与服务:</h4>
            <p>{{ handelGive }}</p>
          </div>
          <div>
            <h4>参与活动:</h4>
            <p>{{ handelSalePromote }}</p>
          </div>
          <div>
            <h4>商家评分:</h4>
            <p>服务评分:{{ suppInfoData.serve_score }}</p>
            <p>包装评分:{{ suppInfoData.pack_score }}</p>
            <p>配送评分:{{ suppInfoData.delivery_score }}</p>
            <p>
              综合评分:
              <el-rate
                v-model="suppInfoData.rating"
                disabled
                show-score
                text-color="#ff9900"
                score-template="{value}"
              >
              </el-rate>
            </p>
          </div>
          <div>
            <h4>更多介绍:</h4>
            <p>
              LOGO：<img
                :src="static_url + suppInfoData.icon"
                style="width: 50px; height: 50px"
                alt=""
              />
            </p>
            <p>{{ suppInfoData.content }}</p>
          </div>
          <div>
            <h4>联系方式:</h4>
            <p>固话:{{ suppInfoData.phone }}</p>
            <p>手机:{{ suppInfoData.mobile }}</p>
            <p>地址:{{ suppInfoData.address }}</p>
            <p>经度:{{ suppInfoData.lng }} 纬度:{{ suppInfoData.lat }}</p>
          </div>
        </div>
      </el-card>
    </el-drawer>
  </div>
</template>

<script>
import global from "@/components/global";
import { reqSuppList } from "../../../api/index";

export default {
  name: "supplist",
  data() {
    return {
      drawer: false,
      formInline: {
        suppName: "商家名称",
        catesName: "默认",
        status: "默认",
        union: "默认",
        position: "默认",
        give: "默认",
        salePromote: "默认",
      },
      rules: {
        // suppName: [
        //   { required: true, message: "请输入", trigger: "blur" },
        //   { min: 3, max: 5, message: "长度在 3 到 5 个字符", trigger: "blur" },
        // ],
        // catesName: [{ required: true, message: "请选择", trigger: "change" }],
      },
      static_url: global.static_url,
      manager_token: global.manager_token,
      suppData: [],
      saveSuppData: [],
      categroyData: [],
      tempArray: [],
      suppInfoData: {},
      statesData: [
        { name: "正常", Key: "state", value: "1" },
        { name: "关闭", Key: "state", value: "2" },
      ],
      unionData: [
        { name: "签约", key: "is_union", value: "1" },
        { name: "未签", key: "is_union", value: "2" },
      ],
      positionData: [
        { name: "首页", key: "is_main", value: "1" },
        { name: "新店", key: "is_new", value: "1" },
        { name: "置顶", key: "is_top", value: "1" },
        { name: "热卖", key: "is_hot", value: "1" },
      ],
      giveData: [],
      salePromoteData: [],
      currentPage: 1,
      pageSize: 5,
      pageTotal: 0,
    };
  },
  computed: {
    handelGive() {
      var giveName = "";
      if (this.suppInfoData.give_id) {
        //   var arr = this.suppInfoData.give_id.split(',')
        //   console.log('a',arr);
        this.giveData.forEach((give) => {
          if (this.suppInfoData.give_id.indexOf(give.id) != -1) {
            giveName += give.name + ",";
          }
        });
      }
      return giveName;
    },
    handelSalePromote() {
      var saleName = "";
      if (this.suppInfoData.sale_id) {
        this.salePromoteData.forEach((sale) => {
          if (this.suppInfoData.sale_id.indexOf(sale.id) != -1) {
            saleName += sale.name + ",";
          }
        });
      }
      return saleName;
    },
  },
  mounted() {
    this.getSuppData();
    this.getCateGroy();
    this.getGiveData();
    this.getSalePromote();
  },
  methods: {
    clearInput() {
      this.resetSearch();
    },
    // 查询
    onSubmit() {
      this.tempArray = [];
      // 名称模糊查询
      if (
        this.formInline.suppName != "" &&
        this.formInline.suppName != "默认" &&
        this.formInline.suppName != "商家名称"
      ) {
        this.suppData.forEach((supp) => {
          let str = ["", ...this.formInline.suppName, ""].join(".*");
          let reg = new RegExp(str);
          if (reg.test(supp.name)) {
            this.tempArray.push(supp);
          }
        });
      }
      // 分类查询
      if (this.tempArray.length > 0) {
        if (
          this.formInline.catesName != "默认" &&
          this.formInline.catesName != "0"
        ) {
          let byCateArray = [];
          this.tempArray.forEach((supp) => {
            if (supp.cate_id == this.formInline.catesName) {
              byCateArray.push(supp);
            }
          });
          this.tempArray = byCateArray;
        }
      } else {
        if (
          this.formInline.catesName != "默认" &&
          this.formInline.catesName != "0"
        ) {
          this.suppData.forEach((supp) => {
            if (supp.cate_id == this.formInline.catesName) {
              this.tempArray.push(supp);
            }
          });
        }
      }
      // 状态查询
      if (this.tempArray.length > 0) {
        if (this.formInline.status != "默认" && this.formInline.status != "0") {
          let byStatusArray = [];
          this.tempArray.forEach((supp) => {
            if (supp.state == this.formInline.status) {
              byStatusArray.push(supp);
            }
          });
          this.tempArray = byStatusArray;
        }
      } else {
        if (this.formInline.status != "默认" && this.formInline.status != "0") {
          this.suppData.forEach((supp) => {
            if (supp.state == this.formInline.status) {
              this.tempArray.push(supp);
            }
          });
        }
      }
      //签约
      if (this.tempArray.length > 0) {
        if (this.formInline.union != "默认" && this.formInline.union != "0") {
          let byUnionArray = [];
          this.tempArray.forEach((supp) => {
            if (supp.is_union == this.formInline.union) {
              byUnionArray.push(supp);
            }
          });
          this.tempArray = byUnionArray;
        }
      } else {
        if (this.formInline.union != "默认" && this.formInline.union != "0") {
          this.suppData.forEach((supp) => {
            if (supp.is_union == this.formInline.union) {
              this.tempArray.push(supp);
            }
          });
        }
      }
      //推荐
      if (this.tempArray.length > 0) {
        if (
          this.formInline.position != "默认" &&
          this.formInline.position != "0"
        ) {
          let byPoArray = [];
          this.tempArray.forEach((supp) => {
            if (this.formInline.position == "is_main") {
              if (supp.is_main == 1) {
                byPoArray.push(supp);
              }
            }
            if (this.formInline.position == "is_new") {
              if (supp.is_new == 1) {
                byPoArray.push(supp);
              }
            }
            if (this.formInline.position == "is_top") {
              if (supp.is_top == 1) {
                byPoArray.push(supp);
              }
            }
            if (this.formInline.position == "is_hot") {
              if (supp.is_hot == 1) {
                byPoArray.push(supp);
              }
            }
          });
          this.tempArray = byPoArray;
        }
      } else {
        if (
          this.formInline.position != "默认" &&
          this.formInline.position != "0"
        ) {
          this.suppData.forEach((supp) => {
            if (this.formInline.position == "is_main") {
              if (supp.is_main == 1) {
                this.tempArray.push(supp);
              }
            }
            if (this.formInline.position == "is_new") {
              if (supp.is_new == 1) {
                this.tempArray.push(supp);
              }
            }
            if (this.formInline.position == "is_top") {
              if (supp.is_top == 1) {
                this.tempArray.push(supp);
              }
            }
            if (this.formInline.position == "is_hot") {
              if (supp.is_hot == 1) {
                this.tempArray.push(supp);
              }
            }
          });
        }
      }
      // 服务查询
      if (this.tempArray.length > 0) {
        if (this.formInline.give != "默认" && this.formInline.give != "0") {
          let byGiveArray = [];
          this.tempArray.forEach((supp) => {
            let res = supp.give_id.indexOf(this.formInline.give);
            // -1 不存在
            if (res != -1) {
              byGiveArray.push(supp);
            }
          });
          this.tempArray = byGiveArray;
        }
      } else {
        if (this.formInline.give != "默认" && this.formInline.give != "0") {
          this.suppData.forEach((supp) => {
            let res = supp.give_id.indexOf(this.formInline.give);
            if (res != -1) {
              this.tempArray.push(supp);
            }
          });
        }
      }

      // 活动查询
      //console.log(this.formInline.salePromote);
      if (this.tempArray.length > 0) {
        if (
          this.formInline.salePromote != "默认" &&
          this.formInline.salePromote != "0"
        ) {
          let bySaleArray = [];
          this.tempArray.forEach((supp) => {
            let res = supp.sale_id.indexOf(this.formInline.salePromote);
            if (res != -1) {
              bySaleArray.push(supp);
            }
          });
          this.tempArray = bySaleArray;
        }
      } else {
        if (
          this.formInline.salePromote != "默认" &&
          this.formInline.salePromote != "0"
        ) {
          this.suppData.forEach((supp) => {
            let res = supp.sale_id.indexOf(this.formInline.salePromote);
            if (res != -1) {
              this.tempArray.push(supp);
            }
          });
        }
      }

      //全是默认
      if (
        this.formInline.suppName == "" &&
        this.formInline.suppName == "默认" &&
        this.formInline.suppName == "商家名称" &&
        this.formInline.catesName == "默认" &&
        this.formInline.catesName == "0" &&
        this.formInline.status == "默认" &&
        this.formInline.status == "0" &&
        this.formInline.union == "默认" &&
        this.formInline.union == "0" &&
        this.formInline.position == "默认" &&
        this.formInline.position == "0" &&
        this.formInline.give == "默认" &&
        this.formInline.give == "0" &&
        this.formInline.salePromote == "默认" &&
        this.formInline.salePromote == "0"
      ) {
        this.resetSearch();
      }
      this.suppData = this.tempArray;
      this.pageTotal = this.suppData.length;
    },
    // 重置查询
    resetSearch() {
      this.formInline.suppName = "";
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
    // 获取商家数据
    async getSuppData() {
      if (!this.manager_token) {
        this.manager_token = localStorage.getItem("manager_token");
      }
      const result = await reqSuppList(this.manager_token);
      if (result.code === 0) {
        result.data.forEach((ele) => {
          this.categroyData.forEach((categroy) => {
            if (ele.cate_id == categroy.id) {
              ele.categroy = categroy.name;
            }
          });
        });
        this.suppData = result.data;
        this.saveSuppData = result.data;
        this.pageTotal = this.suppData.length;
        // console.log(this.suppData);
      } else {
        this.$message.error(result.msg);
      }
    },
    // 获取所有分类数据
    getCateGroy() {
      this.categroyData = this.$store.getters.categroy;
    },
    // 获取所有服务数据
    getGiveData() {
      this.giveData = this.$store.getters.giveData;
      //console.log(this.giveData);
    },
    // 获取所有活动数据
    getSalePromote() {
      this.salePromoteData = this.$store.getters.salePromote;
      //console.log(this.salePromote);
    },
    // 详情
    handleInfo(index, row) {
      this.drawer = true;
      this.suppInfoData = row;
      //console.log(this.suppInfoData);
    },
    // 编辑
    handleEdit(index, row) {
      this.$router.push({ name: "suppedit", params: { suppInfo: row } });
      //console.log(row);
    },
    // 关闭
    handleDelete(row) {
      console.log(row);
      this.$alert("确定要关闭商家,请在编辑页面操作!", "关闭商家: " + row.name, {
        confirmButtonText: "确定",
        callback: () => {
          this.$message({
            type: "info",
            message: `暂未开放此处关闭商家`,
          });
        },
      });
    },
    handleSizeChange(size) {
      this.pageSize = size;
    },
    handleCurrentChange(currentPage) {
      this.currentPage = currentPage;
    },
    handleGoodsList(index,row) {
      console.log(row);
    }
  },
};
</script>

<style  lang="scss">
.el-drawer.rtl {
  overflow: scroll;
}
.m-page-box {
  img {
    width: 380px;
  }

  h4 {
    border-bottom: 1px solid #e6e6e6;
    padding: 6px;
  }

  p {
    padding: 6px;
  }
}
.headClass {
  height: 100vh;
  overflow: auto;
}
</style>