<template>
    <div class="login">
        <div class="logo">
            <img src="../assets/logo.jpg" alt="下单么">
        </div>
        <!-- 手机号 -->
        <InputGroup type="number" v-model="phone" placeholder="手机号" :btnTitle="btnTitle" :disabled="disabled" :error="errors.phone" @btnClick="getVerifyCode"></InputGroup>
        <!-- 验证码 -->
        <InputGroup type="number" v-model="verifyCode" placeholder="验证码" :error="errors.code"></InputGroup>
        <!-- 服务协议 -->
        <div class="login_des">
            <p>
                新用户登录即自动注册,表示已同意
                <span>《用户服务协议》</span>
            </p>
            <p v-if="codeStr">{{codeStr}}</p>
        </div>
        <!-- 登录按钮 -->
        <div class="login_btn">
            <button :disabled="isClick" @click="handleLogin">登录</button>
        </div>
    </div>
</template>

<script>
import InputGroup from '@/components/InputGroup'
import {reqSendCode, reqLogin} from '../api'

export default {
    name: 'login',
    data () {
        return {
            phone: '',
            verifyCode: '',
            errors: {},
            btnTitle: '获取验证码',
            disabled: false,
            codeStr:''
        }
    },
    computed: {
        isClick() {
            if (!this.phone || !this.verifyCode) return true
            else return false
        }
    },
    methods: {
        // 获取验证码
        async getVerifyCode () {
            if (this.validatePhone()) {
                // 发送验证码按钮 倒计时验证
                this.validateBtn()
                // 发送网络请求
                // 聚合数据短信替换id  secret  阿里云 短信 验证id secret
                const result = await reqSendCode(this.phone,'id','secret')
                this.codeStr = result.data
                if (result.code === 1) {
                    this.errors = {
                        code: '获取验证码失败'
                    }
                }
            }
        },
        // 验证手机号
        validatePhone () {
            if (!this.phone) {
                this.errors = {
                    phone: '请输入手机号'
                }
                return false
            } else if (!/^1[345678]\d{9}$/.test(this.phone)) {
                this.errors = {
                    phone: '手机号输入有误'
                }
                return false
            } else {
                this.errors = {}
                return true
            }
        },
        //发送验证码按钮倒计时
        validateBtn() {
            let time = 60
            let timer = setInterval(() => {
                if (time == 0) {
                    clearInterval(timer)
                    this.btnTitle = '获取验证码'
                    this.disabled = false
                } else {
                    // 倒计时
                    this.btnTitle = time + '秒后重试'
                    this.disabled = true
                    time --
                }
            },1000)
        },
        // 登录
        async handleLogin () {
            // 取消错误提醒
            this.errors = {}
            // 发送 登录请求
            const result = await reqLogin(this.phone,this.verifyCode)
            //console.log(result);
            // 登录成功
            if (result.code === 0) {
                localStorage.setItem('user_login',result.data)
                this.$router.push('/')
            }
            //登录失败
            console.log(result);
        }
    },
    components: {
        InputGroup
    }
}
</script>

<style >
    .login {
        width: 100%;
        height: 100%;
        padding: 30px;
        box-sizing: border-box;
        background: #fff;
    }

    .logo {
        text-align: center;
    }
    .logo img {
        width: 150px;
    }
    .text_group,
    .login_des,
    .login_btn {
        margin-top: 20px;
    }
    .login_des {
        color: #aaa;
        line-height: 22px;
    }
    .login_des span {
        color: #4d90fe;
    }
    .login_btn button {
        width: 100%;
        height: 40px;
        background-color: #48ca38;
        border-radius: 4px;
        color: white;
        font-size: 14px;
        border: none;
        outline: none;
    }
    .login_btn button[disabled] {
        background-color: #8bda81;
    }
</style>