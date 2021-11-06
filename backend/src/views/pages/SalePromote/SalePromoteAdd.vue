<template>
  <div style="width: 100%">
    <el-card> 添加活动 </el-card>
    <el-card style="margin-top: 20px; width: 600px">
      <el-form
        :model="promoteInfo"
        :rules="rules"
        ref="promoteInfoRef"
        label-width="100px"
      >
        <el-form-item label="活动名称" prop="name">
          <el-input
            v-model="promoteInfo.name"
            prefix-icon="el-icon-tickets"
            placeholder="请输入活动名称"
            @input="change($event)"
            style="width: 165px"
          ></el-input>
        </el-form-item>
        <el-form-item label="活动状态">
          <el-switch
            v-model="promoteState"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="启用"
            inactive-text="关闭"
            @change="getPromoteState"
          >
          </el-switch>
        </el-form-item>
        <el-form-item label="活动排位" prop="seat">
          <el-input
            type="Number"
            v-model="promoteInfo.seat"
            prefix-icon="el-icon-tickets"
            placeholder="请输入活动排位"
            @input="change($event)"
            style="width: 85px"
          ></el-input>
        </el-form-item>
        <el-form-item label="描述" prop="desc">
          <el-input
            style="width: 320px"
            maxlength="30"
            show-word-limit
            v-model="promoteInfo.desc"
            placeholder="请输入活动描述"
            @input="change($event)"
          >
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handlePromoteAddForm"
            >立即提交</el-button
          >
          <el-button @click="addFormReset">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import global from "@/components/global";
import {reqSalePromoteAdd} from "../../../api/index"
export default {
  name: "salepromoteadd",
  data() {
    return {
      static_url: global.static_url,
      manager_token: global.manager_token,
      promoteInfo: {
        name: "",
        seat: 0,
        state: 1,
        desc:""
      },
      rules: {
        name: [
          { required: true, message: "请输入活动名称", trigger: "blur" },
          {
            min: 2,
            max: 6,
            message: "长度在 2 到6 个字符",
            trigger: "blur",
          },
        ],
      },
      promoteState: true,
    };
  },
  methods: {
    // 多层嵌套无法输入解决方法
    change(e) {
      this.$forceUpdate(e);
    },
    getPromoteState() {
      this.promoteState
        ? (this.promoteInfo.state = 1)
        : (this.promoteInfo.state = 2);
    },
    handlePromoteAddForm() {
      var Seat = parseInt(this.promoteInfo.seat);
      if (isNaN(Seat)) {
        this.$message.error("排位不是数字哦");
        return;
      } else {
        this.promoteInfo.seat = Seat;
      }
      var data = this.promoteInfo;
      var manager_token = "";
      if (this.manager_token) {
        manager_token = this.manager_token;
      } else {
        manager_token = localStorage.getItem("manager_token");
      }
      this.addPromote(manager_token, data);
    },
    addFormReset() {
      this.$router.go(-1);
    },
    async addPromote(token, data) {
      const result = await reqSalePromoteAdd(token, data);
      if (result.code === 0) {
        this.$message({
          message: "添加活动信息成功",
          type: "success",
        });
        this.$router.push("/yunsong/home/salepromotelist");
        return;
      } else {
        this.$message.error("添加活动信息失败");
        return;
      }
    },
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
  border-color: #409eff;
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