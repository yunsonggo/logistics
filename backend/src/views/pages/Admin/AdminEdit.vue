<template>
  <div style="width: 100%">
    <el-card>编辑人员信息</el-card>
    <el-card style="margin-top: 20px; width: 600px">
      <el-form
        :model="adminInfo"
        :rules="rules"
        ref="adminInfoRef"
        label-width="100px"
      >
        <el-form-item label="注册名称" prop="username">
          <el-input
            v-model="adminInfo.username"
            prefix-icon="el-icon-s-custom"
            placeholder="请输入用户名"
            @input="change($event)"
          ></el-input>
        </el-form-item>
        <el-form-item label="昵称" prop="usernick">
          <el-input v-model="adminInfo.usernick"></el-input>
        </el-form-item>
        <el-form-item label="性别" prop="user_gender">
          <el-input v-model="adminInfo.user_gender"></el-input>
        </el-form-item>
        <el-form-item label="电话" prop="phone">
          <el-input v-model="adminInfo.phone"></el-input>
        </el-form-item>
        <el-form-item label="上传头像" prop="user_icon">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/manager/admin/avatar/upload"
            name="admin_avatar_file"
            :data="upParamData"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload"
            multiple="multiple"
          >
            <img
              v-if="adminInfo.user_icon"
              :src="static_url + adminInfo.user_icon"
              class="avatar"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="权限" prop="power">
          <el-radio-group v-model="adminInfo.power">
            <el-radio label="7">管理员</el-radio>
            <el-radio label="77">系统管理员</el-radio>
            <el-radio label="777">超级管理员</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSignUpForm"
            >立即提交</el-button
          >
          <el-button @click="signUpFormReset">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card style="margin-top: 20px; width: 600px">
        <span>修改密码功能暂未实现</span>
    </el-card>
  </div>
</template>

<script>
import global from "@/components/global";
import {reqEditAdminInfo} from "../../../api/index"

export default {
  name: "adminedit",
  data() {
    return {
      static_url: global.static_url,
      upParamData:{
        manager_token: global.manager_token,
        old_avatar_url:"",
        admin_id:""
      },
      adminInfo: {
        id:0,
        username: "",
        usernick: "",
        user_icon: "",
        user_gender: "",
        phone: "",
        power: "",
      },
      rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          {
            min: 6,
            max: 20,
            message: "长度在 6 到 20 个字符",
            trigger: "blur",
          },
        ],
      },
    };
  },
  beforeRouteEnter(to, from, next) {
    if (to.params.adminInfo) {
      next((vm) => {
        vm.adminInfo = to.params.adminInfo;
      });
    }
  },
  methods:{
    // 多层嵌套无法输入解决方法
    change(e) {
      this.$forceUpdate(e);
    },
    // 取消编辑方法
    signUpFormReset() {
      this.$router.go(-1)
    },
    // 注册方法
    handleSignUpForm() {
      // console.log('123456')
      this.$refs.adminInfoRef.validate((valid) => {
        // console.log(valid)
        if (!valid) return;
        //console.log(valid);
        this.editAdminById();
      });
    },
    // 头像上传成功
    handleAvatarSuccess(res, file) {
        this.adminInfo.user_icon = res.path;
        if (res != "") {
            this.$message({
                message: '上传成功',
                type: 'success'
            });
        } else {
            this.$message.error('上传失败');
        }
    },
    // 上传验证
    beforeAvatarUpload(file) {
        const isJPG = file.type === 'image/jpeg';
        const isLt2M = file.size / 1024 / 1024 < 2;

        if (!isJPG) {
          this.$message.error('上传头像图片只能是 JPG 格式!');
        }
        if (!isLt2M) {
          this.$message.error('上传头像图片大小不能超过 2MB!');
        }
        if (isJPG && isLt2M) {
            this.upParamData.old_avatar_url = this.adminInfo.user_icon
            this.upParamData.admin_id = this.adminInfo.id
        } else {
            this.upParamData.old_avatar_url = ""
        }
        //console.log(this.upParamData.old_avatar_url);
        //console.log(this.upParamData.admin_id);
        return isJPG && isLt2M;
      },

    // 实现 人员基本信息 编辑
    async editAdminById() {
        this.$delete(this.adminInfo,'create_time')
        const result = await reqEditAdminInfo(this.upParamData.manager_token,this.adminInfo)
        if (result.code === 0) {
            this.$message({
                message: '编辑人员基本信息成功',
                type: 'success'
            });
            this.$router.push('/yunsong/home/adminlist')
            return
        } else {
            this.$message.error('编辑人员基本信息失败');
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
    width: 98px;
    height: 98px;
    display: block;
  }
</style>