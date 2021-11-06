<template>
    <div>
        <hader :title="title" :isLeft="true"></hader>
        <div class="order">
            <div v-if="allOrdersData">
                <div class="order-card-body" v-for="(order,index) in allOrdersData" :key="index">
                <div class="order-card-wrap" v-if="order.shop_info" @click="$router.push({name:'orderinfo',params:order})">
                    <img v-if="order.shop_info.icon" :src="static_url + order.shop_info.icon" alt="">
                    <!-- 订单内容 -->
                    <div class="order-card-content">
                        <!-- 头 -->
                        <div class="order-card-head">
                            <div class="title">
                                <a>
                                    <span>{{order.shop_name}}</span>
                                    <i class="fa fa-angle-right"></i>
                                </a>
                                <p v-if="order.state == 0">订单未付款</p>
                                <p v-if="order.state == 1">已付款,未送达</p>
                                <p v-if="order.state == 2">订单已完成</p>
                            </div>
                            <p class="date-time">{{order.update_time}}</p>
                        </div>
                        <!--  -->
                        <div class="order-card-detail" v-for="(good,i) in order.sale_goods_list" :key="i">
                            <p class="detail">{{good.good_name}}</p>
                            <p class="price">¥{{good.good_price.toFixed(2)}}</p>
                        </div>
                        <div class="order-card-detail">
                            <p class="detail">总价:</p>
                            <p class="price">¥{{order.total_price.toFixed(2)}}</p>
                        </div>
                    </div>
                </div>
                <div class="order-card-bottom">
                    <button class="cardbutton" v-if="order.state === '0' ">去支付</button>
                    <button class="cardbutton" @click="$router.push({name:'shop',query:{id:order.shop_id}})">再来一单</button>
                </div>
            </div>
            </div>
            <div v-else>
                <div class="order-card-body">
                    <div class="order-card-wrap">
                        <div class="order-card-content">
                            <div class="order-card-head">
                                <div class="title">
                                    <p>暂未订单</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    
</template>

<script>
import hader from "../components/header"
import {reqOrdersData} from "../api/index"
import global from '@/components/global'
export default {
    name:"order",
    data() {
        return {
            title:"我的订单",
            static_url:global.static_url,
            // 所有订单数据
            allOrdersData:[]
        }
    },
    props:{
        user_id:Number,
        userid:Number,
        // 当前订单数据
        orderList:Object
    },
    beforeRouteEnter(to,from,next) {
        next(vm => {
            vm.getData()
        })
    },
    methods:{
        async getData() {
            let uId = 0
            let userId = 0
            if (this.user_id) {
                uId = this.user_id
            } else {
                uId = localStorage.getItem("user_login")
            }
            if (this.userid) {
                userId = this.userid
            } else {
                userId = localStorage.getItem("user")
            }
            const result = await reqOrdersData(uId,userId)
            if (result.code === 0) {
                this.allOrdersData = result.data
            }
            //console.log(result);
        }
    },
    components:{
        reqOrdersData,
        global,
        hader
    }
}
</script>


<style scoped>
.order {
  width: 100%;
  height: 100%;
  overflow: auto;
  box-sizing: border-box;
  margin-bottom: 2.666667vw;
  padding-top: 45px;
}
.order-card-body {
  margin-top: 2.666667vw;
  background-color: #fff;
  padding: 3.733333vw 0 0 4vw;
}
.order-card-wrap {
  display: flex;
}
.order-card-wrap > img {
  height: 8.533333vw;
  border-radius: 0.533333vw;
  border: 1px solid #eee;
  width: 8.533333vw;
  margin-right: 2.666667vw;
}
.order-card-content {
  flex: 1;
}
.order-card-head {
  border-bottom: 1px solid #eee;
  padding-right: 3.466667vw;
  padding-bottom: 2.666667vw;
}
.order-card-head .title {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.order-card-head .title > a {
  font-size: 1rem;
  line-height: 1.5em;
  color: #333;
  text-decoration: none;
}
.order-card-head .title > a > span {
  display: inline-block;
  max-width: 10em;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.order-card-head .title > a > i {
  margin-left: 1.333333vw;
  color: #ccc;
  vertical-align: super;
}
.order-card-head .title > p {
  font-size: 0.8rem;
  text-align: right;
  color: #333;
  max-width: 14em;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.date-time {
  font-size: 0.6rem;
  color: #999;
}
.order-card-detail {
  display: flex;
  justify-content: space-between;
  padding: 4vw 3.466667vw 4vw 0;
  font-size: 0.8rem;
}
.order-card-detail .detail {
  color: #666;
  display: flex;
  align-items: center;
}
.order-card-detail .price {
  flex-basis: 16vw;
  text-align: right;
  color: #333;
  font-weight: 700;
}
.order-card-bottom {
  display: flex;
  border-top: 1px solid #eee;
  padding: 2.666667vw 4.266667vw;
  justify-content: flex-end;
}
.cardbutton {
  padding: 1.333333vw 2.666667vw;
  border: 1px solid #2395ff;
  border-radius: 0.533333vw;
  background-color: transparent;
  outline: none;
  font-size: 0.8rem;
  color: #2395ff;
  margin-left: 2vw;
}
</style>
