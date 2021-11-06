<template>
  <div style="width: 100%">
    <el-card> 编辑类别 </el-card>
    <el-card style="margin-top: 20px; width: 600px">
      <el-form
        :model="catesInfo"
        :rules="rules"
        ref="catesInfoRef"
        label-width="100px"
      >
        <el-form-item label="类别名称" prop="name">
          <el-input
            v-model="catesInfo.name"
            prefix-icon="el-icon-tickets"
            :placeholder="catesInfo.name"
            @input="change($event)"
            style="width: 165px"
          ></el-input>
        </el-form-item>
        <el-form-item label="类别状态">
          <el-switch
            v-model="cateState"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="启用"
            inactive-text="关闭"
            @change="getCateState"
          >
          </el-switch>
        </el-form-item>
        <el-form-item label="类别排位" prop="position">
          <el-input
            type="Number"
            v-model="catesInfo.position"
            prefix-icon="el-icon-tickets"
            :placeholder="catesInfo.position"
            @input="change($event)"
            style="width: 85px"
          ></el-input>
        </el-form-item>
        <el-form-item label="图标" prop="icon">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/manager/category/icon/upload"
            name="cate_icon_file"
            :data="uploadCateIconParamData"
            :show-file-list="false"
            :on-success="handleCateIconSuccess"
            :before-upload="beforeCateIconUpload"
            multiple="multiple"
          >
            <img
              v-if="catesInfo.icon"
              :src="static_url + catesInfo.icon"
              class="logo"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i> </el-upload
          >点击图片上传 (必须是 PNG 格式)
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleCateEditForm"
            >立即提交</el-button
          >
          <el-button @click="cateEditFormReset">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import global from "@/components/global";
import { reqCategoryEdit } from "../../../api/index";
export default {
  name: "catesedit",
  data() {
    return {
      static_url: global.static_url,
      manager_token: global.manager_token,
      catesInfo: {},
      cateState: true,
      rules: {
        name: [
          { required: true, message: "类别名称", trigger: "blur" },
          {
            min: 2,
            max: 6,
            message: "长度在 2 到6 个字符",
            trigger: "blur",
          },
        ],
      },
      uploadCateIconParamData: {
        manager_token: global.manager_token,
        old_cateicon_url: "",
        cate_id: "",
      },
    };
  },
  beforeRouteEnter(to, from, next) {
    next((vm) => {
      vm.catesInfo = to.params.catesInfo;
      vm.catesInfo.state == 1 ? (vm.cateState = true) : (vm.cateState = false);
    });
  },
  methods: {
    // 多层嵌套无法输入解决方法
    change(e) {
      this.$forceUpdate(e);
    },
    getCateState() {
      this.cateState ? (this.catesInfo.state = 1) : (this.catesInfo.state = 2);
    },
    handleCateIconSuccess(res, file) {
      if (res != "") {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.catesInfo.icon = res.path;
      } else {
        this.$message.error("上传失败");
      }
    },
    beforeCateIconUpload(file) {
      const isJPG = file.type === "image/png";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG) {
        this.$message.error("上传图片只能是 PNG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传图片大小不能超过 2MB!");
      }
      if (isJPG && isLt2M) {
        this.uploadCateIconParamData.old_cateicon_url = this.catesInfo.icon;
        this.uploadCateIconParamData.cate_id = this.catesInfo.id;
      } else {
        this.uploadCateIconParamData.old_cateicon_url = "";
      }
      console.log(this.uploadCateIconParamData);
      return isJPG && isLt2M;
    },
    handleCateEditForm() {
      this.$delete(this.catesInfo, "shop_ids");
      this.$delete(this.catesInfo, "is_top");
      this.$delete(this.catesInfo, "is_main");
      this.$delete(this.catesInfo, "is_hot");
      this.$delete(this.catesInfo, "description");
      var Position = parseInt(this.catesInfo.position);
      if (isNaN(Position)) {
        this.$message.error("排位不是数字哦");
        return;
      } else {
        this.catesInfo.position = Position;
      }
      var data = this.catesInfo;
      var manager_token = "";
      if (this.manager_token) {
        manager_token = this.manager_token;
      } else {
        manager_token = localStorage.getItem("manager_token ");
      }
      this.editCategory(manager_token, data);
    },
    cateEditFormReset() {
      this.$router.go(-1);
    },
    async editCategory(manager_token, data) {
      const result = await reqCategoryEdit(manager_token, data);
      if (result.code === 0) {
        this.$message({
          message: "编辑类别信息成功",
          type: "success",
        });
        this.$router.push("/yunsong/home/cateslist");
        return;
      } else {
        this.$message.error("编辑类别信息失败");
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