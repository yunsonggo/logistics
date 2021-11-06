<template>
  <div style="width: 100%">
    <el-card>编辑商家信息</el-card>
    <el-card style="margin-top: 20px; width: 1024px">
      <el-form
        ref="suppInfoRef"
        :model="suppInfo"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item label="商家名称" prop="name">
          <el-input
            v-model="suppInfo.name"
            :placeholder="suppInfo.name"
            @input="change($event)"
            style="width:550px"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="商家类别" prop="cate_id">
          <el-select v-model="cateGroySelected" @change="getCateGroySelected">
            <el-option
              v-for="(cate, index) in cateGroy"
              :key="index"
              :label="cate.name"
              :value="cate.id"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="公告" prop="description">
          <el-input
            type="textarea"
            :rows="2"
            style="width:320px"
            maxlength="55"
            show-word-limit
            v-model="suppInfo.description"
            :placeholder="suppInfo.description"
            @input="change($event)"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="营业时间" prop="opening_hours">
          <el-input
            v-model="suppInfo.opening_hours"
            :placeholder="suppInfo.opening_hours"
            @input="change($event)"
            style="width: 260px"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="商家状态">
          <el-switch
            v-model="suppState"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="正常"
            inactive-text="关闭"
            @change="getSuppState"
          >
          </el-switch>
        </el-form-item>
        <el-form-item label="是否签约">
          <el-switch
            v-model="suppUnion"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="签约"
            inactive-text="未签"
            @change="getSuppUnion"
          >
          </el-switch>
        </el-form-item>
        <el-form-item label="展示位置">
          <el-checkbox-group
            v-model="suppShowPosition"
            @change="getSuppShowPosition"
          >
            <el-checkbox label="首页" name="is_main"></el-checkbox>
            <el-checkbox label="置顶" name="is_top"></el-checkbox>
            <el-checkbox label="热卖" name="is_hot"></el-checkbox>
            <el-checkbox label="新店" name="is_new"></el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item label="提供服务">
          <el-checkbox-group
            v-model="suppGive"
            @change="getSuppGive"
          >
            <el-checkbox v-for="item in giveData" :key="item.id" :label="item.name" ></el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item label="促销活动">
          <el-checkbox-group
            v-model="suppSalePromote"
            @change="getsuppSalePromote"
          >
            <el-checkbox v-for="item in salePromoteData" :key="item.id" :label="item.name" ></el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item label="起送数量" prop="min_num">
          <el-input
            type="number"
            v-model="suppInfo.min_num"
            :placeholder="suppInfo.min_num"
            @input="change($event)"
            style="width: 80px"
          >
          </el-input>件
        </el-form-item>
        <el-form-item label="配送价格" prop="delivery_price">
          <el-input
            type="number"
            v-model="suppInfo.delivery_price"
            :placeholder="suppInfo.delivery_price"
            @input="change($event)"
            style="width: 80px"
          >
          </el-input>元 (10KM以内)
        </el-form-item>
        <el-form-item label="超出运费" prop="over_price">
          <el-input
            type="number"
            v-model="suppInfo.over_price"
            :placeholder="suppInfo.over_price"
            @input="change($event)"
            style="width: 80px"
          >
          </el-input>元 (10KM以外：每公里价格)
        </el-form-item>
        <el-form-item label="订单数量" prop="total_count">
          <el-input
            type="number"
            v-model="suppInfo.total_count"
            :placeholder="suppInfo.total_count"
            @input="change($event)"
            style="width: 80px"
          >
          </el-input>单 (非必要请勿修改)
        </el-form-item>
        <el-form-item label="商家评分">
           <el-input-number v-model="suppInfo.serve_score" :precision="2" :step="0.1" :max="10" style="width:150px"></el-input-number>服务评分
           <el-input-number v-model="suppInfo.delivery_score" :precision="2" :step="0.1" :max="10" style="width:150px"></el-input-number>配送评分
           <el-input-number v-model="suppInfo.pack_score" :precision="2" :step="0.1" :max="10" style="width:150px"></el-input-number>包装评分
           <el-input-number v-model="suppInfo.rating" :precision="2" :step="0.1" :max="10" style="width:150px"></el-input-number>综合评分
        </el-form-item>
        <el-form-item label="固定电话" prop="phone">
          <el-input
            v-model="suppInfo.phone"
            :placeholder="suppInfo.phone"
            @input="change($event)"
            style="width: 130px"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="手机号码" prop="mobile">
          <el-input
            v-model="suppInfo.mobile"
            :placeholder="suppInfo.mobile"
            @input="change($event)"
            style="width: 130px"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="经纬度">
            经度
          <el-input
            type="number"
            v-model="suppInfo.lng"
            :placeholder="suppInfo.lng"
            @input="change($event)"
            style="width: 150px"
          >
          </el-input>
          纬度
          <el-input
            type="number"
            v-model="suppInfo.lat"
            :placeholder="suppInfo.lat"
            @input="change($event)"
            style="width: 150px"
          >
          </el-input>
          (统一固定仓库,无需定位)
        </el-form-item>
        <el-form-item label="联系地址" prop="address">
          <el-input
            v-model="suppInfo.address"
            :placeholder="suppInfo.address"
            @input="change($event)"
            style="width: 530px"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="banner" prop="banner">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/manager/supplier/banner/upload"
            name="supp_banner_file"
            :data="upBannerParamData"
            :show-file-list="false"
            :on-success="handleBannerSuccess"
            :before-upload="beforeBannerUpload"
            multiple="multiple"
          >
            <img
              v-if="suppInfo.banner"
              :src="static_url + suppInfo.banner"
              class="avatar"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>点击图片上传
        </el-form-item>
        <el-form-item label="Logo" prop="icon">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/manager/supplier/logo/upload"
            name="supp_logo_file"
            :data="upLogoParamData"
            :show-file-list="false"
            :on-success="handleLogoSuccess"
            :before-upload="beforeLogoUpload"
            multiple="multiple"
          >
            <img
              v-if="suppInfo.icon"
              :src="static_url + suppInfo.icon"
              class="logo"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>点击图片上传
        </el-form-item>
        <el-form-item label="更多介绍">
            <el-input
            @input="change($event)"
            type="textarea"
            :rows="2"
            :placeholder="suppInfo.content"
            v-model="suppInfo.content"
            style="width:600px"
            maxlength="300"
            show-word-limit
            >
            </el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleEditSuppForm"
            >立即提交</el-button
          >
          <el-button @click="editSuppFormReset">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import global from "@/components/global";
import {reqSuppEdit} from "../../../api/index"
export default {
  name: "suppedit",
  data() {
    return {
      suppInfo: {},
      static_url: global.static_url,
      manager_token: global.manager_token,
      rules: {
        name: [
          { required: true, message: "请输入商家名称", trigger: "blur" },
          {
            min: 4,
            max: 20,
            message: "长度在 4 到 20 个字符",
            trigger: "blur",
          },
        ],
      },
      cateGroy: [],
      cateGroySelected: "",
      suppState: true,
      suppUnion: false,
      suppShowPosition: [],
      giveData: [],
      suppGive: [],
      salePromoteData: [],
      suppSalePromote: [],
      upBannerParamData:{
        manager_token: global.manager_token,
        old_banner_url:"",
        supp_id:""
      },
      upLogoParamData:{
        manager_token: global.manager_token,
        old_logo_url:"",
        supp_id:""
      }
    };
  },
  beforeRouteEnter(to, from, next) {
    next((vm) => {
      vm.suppInfo = to.params.suppInfo;
      vm.suppInfo.state == 1 ? (vm.suppState = true) : (vm.suppState = false);
      vm.suppInfo.is_union == 1
        ? (vm.suppUnion = true)
        : (vm.suppUnion = false);
      if (vm.suppInfo.is_main == 1) {
        vm.suppShowPosition.push("首页");
      }
      if (vm.suppInfo.is_top == 1) {
        vm.suppShowPosition.push("置顶");
      }
      if (vm.suppInfo.is_new == 1) {
        vm.suppShowPosition.push("新店");
      }
      if (vm.suppInfo.is_hot == 1) {
        vm.suppShowPosition.push("热卖");
      }
    });
  },
  mounted() {
    this.getCateGroy();
    this.getGiveData();
    this.getSalePromoteData();
  },
  methods: {
    change(e) {
      this.$forceUpdate(e);
    },
    getCateGroy() {
      this.cateGroy = this.$store.getters.categroy;
      this.cateGroy.forEach((cate) => {
        if (cate.id == this.suppInfo.cate_id) {
          this.cateGroySelected = cate.name;
        }
      });
    },
    getCateGroySelected() {
        this.suppInfo.cate_id = this.cateGroySelected;
    },
    getSuppState() {
      this.suppState ? (this.suppInfo.state = 1) : (this.suppInfo.state = 2);
    },
    getSuppUnion() {
      this.suppUnion
        ? (this.suppInfo.is_union = 1)
        : (this.suppInfo.is_union = 2);
      //console.log(this.suppInfo.is_union);
    },
    getSuppShowPosition() {
      console.log(this.suppShowPosition);
      this.suppInfo.is_main = 2;
      this.suppInfo.is_top = 2;
      this.suppInfo.is_hot = 2;
      this.suppInfo.is_new = 2;
      this.suppShowPosition.forEach((item) => {
        if (item == "首页") {
          this.suppInfo.is_main = 1;
        }
        if (item == "置顶") {
          this.suppInfo.is_top = 1;
        }
        if (item == "热卖") {
          this.suppInfo.is_hot = 1;
        }
        if (item == "新店") {
          this.suppInfo.is_new = 1;
        }
      });
    },
    getGiveData() {
      if (this.$store.getters.giveData) {
        this.giveData = this.$store.getters.giveData;
        if (this.suppInfo.give_id) {
          this.giveData.forEach((give) => {
            if (this.suppInfo.give_id.indexOf(give.id) != -1) {
              this.suppGive.push(give.name);
            }
          });
        }
      }
    },
    getSuppGive() {
        var nameArr = []
        this.giveData.forEach((give) => {
            this.suppGive.forEach((giveName) => {
                if(give.name == giveName) {
                    nameArr.push(give.id)
                }
            })
        })
        var nameStr = nameArr.toString()
        this.suppInfo.give_id = nameStr
    },
    getSalePromoteData() {
        if(this.$store.getters.salePromote) {
            this.salePromoteData = this.$store.getters.salePromote
            if (this.suppInfo.sale_id) {
                this.salePromoteData.forEach((item) => {
                    if (this.suppInfo.sale_id.indexOf(item.id) != -1) {
                        this.suppSalePromote.push(item.name)
                    }
                })
            }
        }
    },
    getsuppSalePromote() {
        var nameArr = []
        this.salePromoteData.forEach((sale) => {
            this.suppSalePromote.forEach((saleName) => {
                if (sale.name == saleName) {
                    nameArr.push(sale.id)
                }
            })
        })
        var nameStr = nameArr.toString()
        this.suppInfo.sale_id = nameStr
        //console.log(this.suppInfo.sale_id);
    },
    // banner 上传成功
    handleBannerSuccess(res,file) {
        if (res != "") {
            this.$message({
                message: '上传成功',
                type: 'success'
            });
            this.suppInfo.banner = res.path;
        } else {
            this.$message.error('上传失败');
        }
    },
    // banner 上传验证
    beforeBannerUpload(file) {
        const isJPG = file.type === 'image/jpeg';
        const isLt2M = file.size / 1024 / 1024 < 2;
        if (!isJPG) {
          this.$message.error('上传图片只能是 JPG 格式!');
        }
        if (!isLt2M) {
          this.$message.error('上传图片大小不能超过 2MB!');
        }
        if (isJPG && isLt2M) {
            this.upBannerParamData.old_banner_url = this.suppInfo.banner
            this.upBannerParamData.supp_id = this.suppInfo.id
        } else {
            this.upBannerParamData.old_banner_url = ""
        }
        //console.log(this.upBannerParamData);
        return isJPG && isLt2M;
    },
    // Logo上传成功
    handleLogoSuccess(res,file) {
        if (res != "") {
            this.$message({
                message: '上传成功',
                type: 'success'
            });
            this.suppInfo.icon = res.path;
        } else {
            this.$message.error('上传失败');
        }
    },
    // Logo 上传验证
    beforeLogoUpload(file) {
        const isJPG = file.type === 'image/jpeg';
        const isLt2M = file.size / 1024 / 1024 < 2;
        if (!isJPG) {
          this.$message.error('上传图片只能是 JPG 格式!');
        }
        if (!isLt2M) {
          this.$message.error('上传图片大小不能超过 2MB!');
        }
        if (isJPG && isLt2M) {
            this.upLogoParamData.old_logo_url = this.suppInfo.icon
            this.upLogoParamData.supp_id = this.suppInfo.id
        } else {
            this.upLogoParamData.old_logo_url = ""
        }
        console.log(this.upLogoParamData);
        return isJPG && isLt2M;
    },
    // 提交
    handleEditSuppForm() {
        this.$delete(this.suppInfo,'create_time')
        if (!this.manager_token) {
            this.manager_token = localStorage.getItem("manager_token")
        }
        var minNum = parseInt(this.suppInfo.min_num)
        if (isNaN(minNum)) {
            this.$message.error('起送数量不是数字哦');
            return
        } else {
            this.suppInfo.min_num = minNum
        }
        var deliveryPrice = parseFloat(this.suppInfo.delivery_price) 
        if (isNaN(deliveryPrice)) {
            this.$message.error('配送价格不是数字哦');
            return
        } else {
            this.suppInfo.delivery_price = deliveryPrice
        }
        var overPrice = parseFloat(this.suppInfo.over_price)
        if (isNaN(overPrice)) {
            this.$message.error('超出运费不是数字哦');
            return
        } else {
            this.suppInfo.over_price = overPrice
        }
        var totalCount = parseInt(this.suppInfo.total_count)
        if (isNaN(totalCount)) {
            this.$message.error('订单数量不是数字哦');
            return
        } else {
            this.suppInfo.total_count = totalCount
        }
        var serveScore = parseFloat(this.suppInfo.serve_score)
        if (isNaN(serveScore)) {
            this.$message.error('服务评分不是数字哦');
            return
        } else {
            this.suppInfo.serve_score = serveScore
        }
        var deliveryScore = parseFloat(this.suppInfo.delivery_score)
        if (isNaN(deliveryScore)) {
            this.$message.error('配送评分不是数字哦');
            return
        } else {
            this.suppInfo.delivery_score = deliveryScore
        }
        var packScore = parseFloat(this.suppInfo.pack_score)
        if (isNaN(packScore)) {
            this.$message.error('包装评分不是数字哦');
            return
        } else {
            this.suppInfo.pack_score = packScore
        }
        var Rating = parseFloat(this.suppInfo.rating)
        if (isNaN(Rating)) {
            this.$message.error('综合评分不是数字哦');
            return
        } else {
            this.suppInfo.rating = Rating
        }
        var LngNum = parseFloat(this.suppInfo.lng)
        if (isNaN(LngNum)) {
            this.$message.error('经度不是数字哦');
            return
        } else {
            this.suppInfo.lng = LngNum
        }
        var LatNum = parseFloat(this.suppInfo.lat)
        if (isNaN(LatNum)) {
            this.$message.error('纬度不是数字哦');
            return
        } else {
            this.suppInfo.lat = LatNum
        }
        var supp = this.suppInfo
        this.handleEditSupp(supp)
    },
    // 取消
    editSuppFormReset() {
        this.$router.go(-1)
    },
    // 实现数据更新
    async handleEditSupp(supp) {
        //console.log(supp);
        const result = await reqSuppEdit(this.manager_token,supp)
        if (result.code === 0) {
            this.$message({
                    message: '商家信息修改成功',
                    type: 'success'
                })
                this.$router.push('/yunsong/home/supplist')
                return
        } else {
            this.$message.error('编辑商家信息失败');
            return
        }
    }
  },
};
</script>

<style scoped>
.el-upload {
    display: inline-block;
    text-align: center;
    cursor: pointer;
    outline: none;
}
.avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 98px;
    height: 98px;
    line-height: 98px;
    text-align: center;
  }
  .avatar {
    border: 1px dashed #d9d9d9;
    width: 298px;
    height: 118px;
    display: block;
  }
  .logo {
    border: 1px dashed #d9d9d9;
    width: 98px;
    height: 98px;
    display: block;
  }
</style>