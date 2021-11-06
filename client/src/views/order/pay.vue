<template>
    <div class="pay">
        <hader :isLeft="true" :title="title"></hader>
        <div class="main" v-if="orderInfo">
            <div class="tip">
                <p class="countdown-title">支付剩余时间</p>
                <h3 class="countdown-text">{{countDown}}</h3>
            </div>
            <section class="home">
                <ul class="list info-list">
                    <li>
                        <span class="order-name">{{orderInfo.shopInfo.name}}</span>
                        <span class="text-highlight">¥{{totalPrice.toFixed(2)}}</span>
                    </li>
                </ul>
                <div class="title">在线支付</div>
                <ul class="list payment-list">
                    <li class="payment-list-item" @click="selectPay = '0'">
                        <div class="icon">
                            <img src="../../assets/wechart.jpg" alt="">
                            <span>微信</span>
                        </div>
                        <i v-if="selectPay == '0'" class="fa fa-check-circle selected"></i>
                    </li>
                    <li class="payment-list-item" @click="selectPay = '1'">
                        <div class="icon">
                            <img src="../../assets/zhifubao.png" alt="">
                            <span>支付宝</span>
                        </div>
                        <i v-if="selectPay == '1'" class="fa fa-check-circle selected"></i>
                    </li>
                </ul>
            </section>
            <button @click="addOrder" :disabled="timeOut" class="btn-submit">确认支付</button>
        </div>
    </div>
</template>

<script>
import hader from '../../components/header'
import {reqOrderSend} from '../../api/index'
import {Toast} from 'mint-ui'

export default {
    name: "pay",
    data() {
        return {
            title:"在线支付",
            countDown:"00:15:00",
            timer:null,
            timeOut:false,
            selectPay:'0',
        }
    },
    beforeRouteEnter (to,from,next) {
        next(vm => {
            vm.selectPay = "0"
            vm.countTimeDown()
        })
    },
    computed:{
        orderInfo () {
          if (this.$store.getters.orderInfo) {
            return this.$store.getters.orderInfo
          } else {
            if (localStorage.getItem("orderInfo")) {
              return JSON.parse(localStorage.getItem("orderInfo"))
            }
          }
          
        },
        totalPrice () {
          if (this.$store.getters.totalPrice) {
            return this.$store.getters.totalPrice
          } 
        },
        addressInfo() {
          if (this.$store.getters.selectedAddress) {
            return this.$store.getters.selectedAddress
          } else {
            if(localStorage.getItem("selectedAddress")) {
              return JSON.parse(localStorage.getItem("selectedAddress"))
            }
          }
        },
        remarkInfo () {
          if (this.$store.getters.remarkInfo) {
            return this.$store.getters.remarkInfo
          } else {
            if(localStorage.getItem("remarkInfo")) {
              return JSON.parse(localStorage.getItem("remarkInfo"))
            }
          }
        },
        goodsPrice () {
          if (this.$store.getters.goodsPrice) {
            return this.$store.getters.goodsPrice
          }
        },
        deliveryPrice () {
          if (this.$store.getters.deliveryPrice) {
            return this.$store.getters.deliveryPrice
          }
        },
        distanceMetre () {
          if (this.$store.getters.distanceMetre) {
            return this.$store.getters.distanceMetre
          }
        },
        deliveryDateTime () {
          if (this.$store.getters.deliveryDateTime) {
            return this.$store.getters.deliveryDateTime
          }
        }
    },
    methods:{
        countTimeDown() {
            let minute = 14
            let second = 59
            this.timer = setInterval(() => {
                second --
                if (second == "00" && minute == "00") {
                    this.countDown = "订单支付已超时"
                    clearInterval(this.timer)
                    this.timeOut = true
                    return
                }
                if (second == "00") {
                    second = 59
                    minute --
                    if (minute < 10) {
                        minute = "0" + minute
                    }
                }
                if (second < 10 ) {
                    second = "0" + second
                }
                this.countDown = "00:" + minute + ":" + second
            },1000)
        },
        // 向后台发送订单数据 并 完成支付
        async addOrder() {
          let user_id = localStorage.getItem("user_login")
          let userid = localStorage.getItem("user") 
          let orderList = {
            userid: userid,
            orderInfo:this.orderInfo,
            totalPrice:this.totalPrice,
            addressInfo:this.addressInfo,
            remarkInfo:this.remarkInfo,
            selectPay:this.selectPay,
            goodsPrice:this.goodsPrice,
            deliveryPrice:this.deliveryPrice,
            distanceMetre:this.distanceMetre,
            orderDelTime:this.deliveryDateTime
          }
          console.log(orderList);
          const result = await reqOrderSend(user_id,userid,orderList)
          //console.log(result);
          if (result.code === 0) {
            this.saveOrder(user_id,userid,orderList)
          } else {          
            Toast({
              message:"支付出现错误,请重试",
              position:"middle",
              duration:2000
            })
            this.$router.push('/pay')
          }
        },
        // 前端添加订单
        saveOrder (user_id,userid,orderList) {
          if (orderList) {
            this.$router.push({
              name:"order",
              params:{
                user_id:user_id,
                userid:userid,
                orderList: orderList
              }
            })
          }
        }

        // pay() {
        //     const data = {
        //         body:"下单么",
        //         out_trade_no: new Date().getTime().toString(),
        //         total_fee: 1 // 单位 分
        //     }
        //     alert("进入到pay方法中");
        //     // 请求 http://www.thenewstep.cn/wxzf/example/jsapi.php
        //     fetch("http://www.thenewstep.cn/wxzf/example/jsapi.php", {
        //         method: "POST",
        //         headers: {
        //             "Content-type": "application/json"
        //         },
        //         body: JSON.stringify(data)
        //     })
        //     .then(res => res.json())
        //     .then(data => {
        //         this.onBridgeReady(data);
        //     })
        //     .catch(err => {
        //         alert("请求失败");
        //     });
        // },
        // // 微信浏览器内部调用
        // onBridgeReady(data) {
        //     WeixinJSBridge.invoke("getBrandWCPayRequest", data, res => {
        //         if (res.err_msg == "get_brand_wcpay_request:ok") {
        //         // 使用以上方式判断前端返回,微信团队郑重提示：
        //         //res.err_msg将在用户支付成功后返回ok，但并不保证它绝对可靠。
        //         // alert("支付成功");
        //         // 生成订单
        //         //this.addOrder();
        //         }
        //     });
        // },
    },
    components:{
        hader
    }
}
</script>


<style scoped>
.pay {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  padding-top: 45px;
}

.main {
  width: 100%;
  height: 100%;
  overflow: auto;
  box-sizing: border-box;
}
.tip {
  display: block;
  text-align: center;
  background-color: #fff;
  color: #333;
  padding: 2rem;
  line-height: 2rem;
}
countdown-title {
  font-size: 0.88rem;
  color: #999;
}
.countdown-text {
  height: 2rem;
  color: #333;
  font-size: 1.8rem;
}
down-text {
  height: 2rem;
  color: #333;
  font-size: 1.8rem;
}
.list {
  border-bottom: 0.5px solid #eee;
  background: #fff;
}
.info-list {
  padding: 0 1.5rem;
}
.info-list li {
  border-top: 0.5px solid #eee;
  display: flex;
  padding: 1.5rem 0;
  font-size: 1rem;
  line-height: 1rem;
  justify-content: space-between;
}
.info-list li .order-name {
  color: #333;
  display: inline-block;
  vertical-align: bottom;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  margin-right: 16px;
  max-width: 60%;
}
.text-highlight {
  color: #ff6000;
}
.title {
  background-color: #f5f5f5;
  font-size: 0.88rem;
  padding: 1.2rem 1.2rem 1rem;
  color: #999;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.payment-list-item {
  border-bottom: 0.5px solid #eee;
  padding: 0.9rem 1rem;
  font-size: 1.1rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.icon {
  display: flex;
  align-items: center;
}
.payment-list-item img {
  width: 2.4rem;
  height: 2.4rem;
  margin-right: 20px;
}
.payment-list-item > i {
  font-size: 1.5rem;
  color: #eee;
}
.selected {
  color: #4cd964 !important;
}
.btn-submit-wrap {
  margin: 1rem auto;
  width: 90%;
}
.btn-submit {
  font-size: 1.1rem;
  line-height: 1.5rem;
  background-color: #4cd964;
  width: 100%;
  outline: none;
  border: none;
  color: #fff;
  border-radius: 5px;
  padding: 0.5rem;
  box-sizing: border-box;
}

/* 不可点击btn */
.btn-submit[disabled] {
  background-color: #bbb !important;
}
</style>

