<template>
  <div style="width: 100%">
    <el-card> 添加服务 </el-card>
    <el-card style="margin-top: 20px; width: 600px">
      <el-form
        :model="giveInfo"
        :rules="rules"
        ref="giveInfoRef"
        label-width="100px"
      >
        <el-form-item label="服务名称" prop="name">
          <el-input
            v-model="giveInfo.name"
            prefix-icon="el-icon-tickets"
            placeholder="请输入服务名称"
            @input="change($event)"
            style="width: 165px"
          ></el-input>
        </el-form-item>
        <el-form-item label="服务状态">
          <el-switch
            v-model="giveState"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="启用"
            inactive-text="关闭"
            @change="getGiveState"
          >
          </el-switch>
        </el-form-item>
        <el-form-item label="服务排位" prop="seat">
          <el-input
            type="Number"
            v-model="giveInfo.seat"
            prefix-icon="el-icon-tickets"
            :placeholder="giveInfo.seat"
            @input="change($event)"
            style="width: 85px"
          ></el-input>
        </el-form-item>
        <el-form-item label="图标" prop="icon">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/manager/give/icon/upload"
            name="give_icon_file"
            :data="uploadGiveIconParamData"
            :show-file-list="false"
            :on-success="handleGiveIconSuccess"
            :before-upload="beforeGiveIconUpload"
            multiple="multiple"
          >
            <img
              v-if="giveInfo.icon"
              :src="static_url + giveInfo.icon"
              class="logo"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i> </el-upload
          >点击图片上传 (必须是 PNG 格式)
        </el-form-item>
        <el-form-item label="描述" prop="desc">
          <el-input
            style="width: 320px"
            maxlength="30"
            show-word-limit
            v-model="giveInfo.desc"
            placeholder="请输入服务描述"
            @input="change($event)"
          >
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleGiveAddForm"
            >立即提交</el-button
          >
          <el-button @click="giveAddFormReset">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import global from "@/components/global";
import { reqGiveAdd } from "../../../api/index";
export default {
  name: "givesadd",
  data() {
    return {
      static_url: global.static_url,
      manager_token: global.manager_token,
      giveInfo: {
        name: "",
        state: 1,
        icon: "",
        seat: 0,
        desc:""
      },
      rules: {
        name: [
          { required: true, message: "服务名称", trigger: "blur" },
          {
            min: 2,
            max: 6,
            message: "长度在 2 到6 个字符",
            trigger: "blur",
          },
        ],
      },
      giveState: true,
      uploadGiveIconParamData: {
        manager_token: global.manager_token,
        old_giveicon_url: "",
        give_id: "",
      },
    };
  },
  methods: {
    // 多层嵌套无法输入解决方法
    change(e) {
      this.$forceUpdate(e);
    },
    getGiveState() {
      this.giveState ? (this.giveInfo.state = 1) : (this.giveInfo.state = 2);
    },
    handleGiveIconSuccess(res, file) {
      if (res != "") {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.giveInfo.icon = res.path;
      } else {
        this.$message.error("上传失败");
      }
    },
    beforeGiveIconUpload(file) {
      const isJPG = file.type === "image/png";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG) {
        this.$message.error("上传图片只能是 PNG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传图片大小不能超过 2MB!");
      }
      if (isJPG && isLt2M) {
        this.uploadGiveIconParamData.old_giveicon_url = this.giveInfo.icon;
        this.uploadGiveIconParamData.give_id = this.giveInfo.id;
      } else {
        this.uploadGiveIconParamData.old_giveicon_url = "";
      }
      return isJPG && isLt2M;
    },
    handleGiveAddForm() {
      var Seat = parseInt(this.giveInfo.seat);
      if (isNaN(Seat)) {
        this.$message.error("排位不是数字哦");
        return;
      } else {
        this.giveInfo.seat = Seat;
      }
      var data = this.giveInfo;
      var manager_token = "";
      if (this.manager_token) {
        manager_token = this.manager_token;
      } else {
        manager_token = localStorage.getItem("manager_token ");
      }
      this.addGive(manager_token, data);
    },
    giveAddFormReset() {
      this.$router.go(-1);
    },
    async addGive(token, data) {
      const result = await reqGiveAdd(token, data);
      if (result.code === 0) {
        this.$message({
          message: "添加服务信息成功",
          type: "success",
        });
        this.$router.push("/yunsong/home/giveslist");
        return;
      } else {
        this.$message.error("添加服务信息失败");
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