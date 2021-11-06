<template>
    <div class="container">
    <div class="logo-box">
      <img src="../assets/images/logo.png" alt="" />
    </div>
    <div class="login-box">
      <div class="login-text">
        注 册 新 用 户
        <br />
        <span><el-button type="text" @click="$router.push({name:'login'})"
            >返回登录</el-button
          ></span>
      </div>
      <el-form :model="signUpForm" ref="signUpFormRef" :rules="signUpFormRules">
        <el-form-item prop="username">
          <el-input
            v-model="signUpForm.username"
            prefix-icon="el-icon-s-custom"
            placeholder="请输入用户名"
            @input="change($event)"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="signUpForm.password"
            prefix-icon="el-icon-unlock"
            placeholder="请输入密码"
            @input="change($event)"
            type="password"
          ></el-input>
        </el-form-item>
        <el-form-item prop="checkPassword">
          <el-input
            v-model="signUpForm.checkPassword"
            prefix-icon="el-icon-unlock"
            placeholder="请确认密码"
            @input="change($event)"
            type="password"
          ></el-input>
        </el-form-item>
        <el-form-item prop="captcha_token">
          <div ref="captchaBox" class="captcha">
            <input type="hidden" v-model="signUpForm.captcha_token">
          </div>
        </el-form-item>
        <el-form-item class="btns">
          <el-button type="primary" @click="handleSignUpForm">注册</el-button>
          <el-button type="primary" @click="signUpFormReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {reqManagerVerifyCaptcha,reqManagerSignUp} from '../api/index'
export default {
    name: "signup",
    data() {
    var validatePass2 = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请再次输入密码'));
        } else if (value !== this.signUpForm.password) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
    };

    return {
      appId: "62602185bf8b6f56e5b1e70fe7da4e60",
      signUpForm: {
        username: "",
        password: "",
        checkPassword: "",
        captcha_token: "",
      },
      // 登录表单的校验对象
      signUpFormRules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          {
            min: 6,
            max: 20,
            message: "长度在 6 到 20 个字符",
            trigger: "blur",
          },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          {
            min: 6,
            max: 20,
            message: "长度在 6 到 20 个字符",
            trigger: "blur",
          },
        ],
        checkPassword: [
          { required: true, message: "请确认密码", trigger: "blur" },
          {
            validator: validatePass2,
            trigger: "blur",
          },
        ],
        captcha_token: [
          { required: true, message: "请滑动验证后登录", trigger: "blur" },
        ]
      },
    };
  },
  beforeRouteEnter(to, from, next) {
    next((vm) => {
      vm.captchaHandle();
    });
  },
  methods: {
    // 多层嵌套无法输入解决方法
    change(e) {
      this.$forceUpdate(e);
    },
    // 表单重置方法
    signUpFormReset() {
      this.$refs.signUpFormRef.resetFields();
    },
    // 注册方法
    handleSignUpForm() {
      // console.log('123456')
      this.$refs.signUpFormRef.validate((valid) => {
        // console.log(valid)
        if (!valid) return;
        //console.log(valid);
        this.signUpByPass()
      });
    },
    // 加载 顶象验证码
    captchaHandle() {
      var myCaptcha = _dx.Captcha(this.$refs.captchaBox, {
        appId: this.appId,
        //type: 'basic', // <-- 指定为"基础类型"，此参数可省略
        //style: 'embed', // 可省略
        width: 350, // 可省略
      });
      myCaptcha.on("ready", () => {
        console.log("captcha is ready!");
      });
      myCaptcha.on("verifySuccess", (security_code) => {
        // 说明：security_code = token + ':' + constID
        //console.log("security_code is: " + security_code);
        setTimeout(() => {
          myCaptcha.reload();
          this.signUpForm.captcha_token = "";
        }, 10000);
        this.sendCaptchaToken(security_code);
      });
    },
    // 发送TOKEN 后台验证
    async sendCaptchaToken(security_code) {
      const result = await reqManagerVerifyCaptcha(security_code);
      if (result.code === 0) {
        this.signUpForm.captcha_token = "0"
      }
    },
    // 发送注册请求
    async signUpByPass() {
      const result = await reqManagerSignUp(this.signUpForm.username,this.signUpForm.password)
      if (result.code === 1) {
          // 注册失败
           this.$message.error(result.msg)
           return
      } else if (result.code === 0) {
          // 注册成功 跳转登录
            this.$message({
                message: '注册成功!',
                type: 'success'
            });
            setTimeout(() => {
                this.$router.push({name:'login',params:{username:this.signUpForm.username}})
                return
            },1000)   
      } else {
          return
      }
    }
  },
}
</script>

<style scoped>
.container {
  width: 100%;
  height: 100%;
  background-image: url("../assets/images/login_bg.png");
  background-size: 100% 100%;
}
.logo-box {
  position: absolute;
  top: 30px;
  left: 30px;
}
.login-box {
  position: absolute;
  top: 50%;
  left: 50%;
  width: 400px;
  height: 690px;
  transform: translate(-50%, -50%);
  border: 1px solid #ccc;
  border-radius: 5px;
  background-color: azure;
}
.login-text {
  width: 100%;
  font-size: 24px;
  text-align: center;
  color: #199efc;
  margin-bottom: 50px;
  box-sizing: border-box;
  padding: 20px;
}
.el-form-item {
  width: 90%;
  margin: 20px auto;
}
.el-form-item .captcha {
  width: 90%;
  margin: 20px auto;
}
.login-text span {
  line-height: 30px;
  font-size: 18px;
  color: #199efc;
}
.btns {
  width: 100%;
  text-align: center;
  margin-top: 50px;
}
</style>