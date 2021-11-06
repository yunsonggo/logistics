<template>
    <div class="orderInfo">
        <hader :title="title" :isLeft="true"></hader>
        <div class="view-body" v-if="orderDetail">
            <div class="status-head" v-if="orderDetail.state == 2">
                <div class="status-text">订单已完成</div>
                <div class="status-title">感谢您对'下单么'的信任,期待再次光临</div>
            </div>
            <div class="status-head" v-if="orderDetail.state == 1">
                <div class="status-text">订单已付款,暂未送达</div>
                <div class="status-title">小哥正在积极备货配送中...</div>
            </div>
            <div class="status-head" v-if="orderDetail.state == 0">
                <div class="status-text">订单未付款</div>
                <div class="status-title">订单已提交,暂未付款</div>
            </div>
            <!-- 商品信息 -->
            <div class="restaurant-card">
                <ordercartgroup :orderInfo="orderDetail" :totalPrice="orderDetail.total_price" :delPrice="orderDetail.delivery_price"></ordercartgroup>
            </div>
            <!-- 配送信息 -->
            <div class="detail-card">
                <div class="title">配送信息</div>
                <ul class="card-list">
                    <li class="list-item" v-if="orderDetail.state == '2' ">
                        <span >送达时间</span>
                        <div>{{orderDetail.update_time}}</div>
                    </li>
                    <li class="list-item" v-else-if="orderDetail.state != '2' && orderDetail.delivery_datetime">
                        <span>送达时间</span>
                        <div>{{orderDetail.delivery_datetime}}</div>
                    </li>
                    <li class="list-item" v-else>
                        <span>订单时间:</span>
                        <div>{{orderDetail.update_time}}</div>
                    </li>
                    <li class="list-item">
                        <span>送货地址:</span>
                        <div class="content">
                            <span>{{orderDetail.liaison}}{{orderDetail.gender}}</span>
                            <span>{{orderDetail.liaison_phone}}</span>
                            <span>{{orderDetail.address}}</span>
                        </div>
                    </li>
                    <li class="list-item">
                        <span>配送路程:</span>
                        <div>{{orderDetail.metre.toFixed(2)}}KM</div>
                    </li>
                </ul>
            </div>
            <!-- 订单信息 -->
            <div class="detail-card">
                <div class="title">订单信息</div>
                <ul class="card-list">
                    <li class="list-item">
                        <span>订单编号:</span>
                        <div>{{orderDetail.order_code}}</div>
                    </li>
                    <li class="list-item">
                        <span>支付方式:</span>
                        <div v-if="orderDetail.select_pay == '0'">微信支付</div>
                        <div v-if="orderDetail.select_pay == '1'">支付宝支付</div>
                    </li>
                    <li class="list-item">
                        <span>下单时间:</span>
                        <div>{{orderDetail.create_time}}</div>
                    </li>
                </ul>
            </div>
            <!-- 备注信息 -->
            <div class="detail-card">
                <div class="title">备注信息</div>
                <ul class="card-list">
                    <li class="list-item">
                        <span>调货换货:</span>
                        <div>{{orderDetail.exchange}}</div>
                    </li>
                    <li class="list-item">
                        <span>订单备注:</span>
                        <div>{{orderDetail.remark}}</div>
                    </li>
                    <li class="list-item">
                        <span>发票信息:</span>
                        <div>{{orderDetail.invoice}}</div>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script>
import hader from "../../components/header"
import ordercartgroup from "../../components/orders/ordercartgroup"
export default {
    data () {
        return {
            title:"订单详情",
            orderDetail: null
        }
    },
    beforeRouteEnter(to,from,next) {
        next(vm => {
            vm.orderDetail = to.params;
            console.log(vm.orderDetail);
        })
    },
    components:{
        hader,
        ordercartgroup
    }
}
</script>

<style scoped>
.orderInfo {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  padding-top: 45px;
}

.view-body {
  width: 100%;
  height: 100%;
  overflow: auto;
  box-sizing: border-box;
}
.status-head {
  margin: 2.666667vw;
  box-shadow: 0 0.133333vw 0.266667vw 0 rgba(0, 0, 0, 0.05);
  background-color: #fff;
  padding: 0 3.2vw 5.333333vw;
}
.status-head .status-text {
  margin: 0 auto 4.266667vw;
  color: #333;
  font-size: 1.5rem;
  text-align: left;
  padding: 4.266667vw 0;
  border-bottom: 0.133333vw solid #ddd;
}
.status-head .status-title {
  font-size: 1rem;
  color: #333;
  margin-bottom: 1.866667vw;
}

.restaurant-card {
  margin: 2.666667vw;
  box-shadow: 0 0.133333vw 0.266667vw 0 rgba(0, 0, 0, 0.05);
  background-color: #fff;
  padding: 0 3.2vw 5.333333vw;
}

/* 配送和订单 */
.detail-card {
  margin: 2.666667vw;
  box-shadow: 0 0.133333vw 0.266667vw 0 rgba(0, 0, 0, 0.05);
  background-color: #fff;
  padding: 0 3.2vw 5.333333vw;
}

.status-head .status-text {
  margin: 0 auto 4.266667vw;
  color: #333;
  font-size: 1.5rem;
  text-align: left;
  padding: 4.266667vw 0;
  border-bottom: 0.133333vw solid #ddd;
}
.status-head .status-title {
  font-size: 1rem;
  color: #333;
  margin-bottom: 1.866667vw;
}
.title {
  font-size: 1rem;
  color: #333;
  border-bottom: 1px solid #eee;
  line-height: 10.4vw;
}
.list-item {
  color: #6e6e6e;
  border-bottom: 1px solid #f2f2f2;
  display: flex;
  align-items: baseline;
  line-height: 4.8vw;
  font-size: 0.88rem;
  padding: 3.2vw 8vw 2.666667vw 0;
}
.list-item .content {
  line-height: 1.5em;
  padding-bottom: 2.666667vw;
  display: flex;
  flex-direction: column;
  flex: 1;
}
.list-item span:first-child {
  flex: none;
}
.remark {
  text-overflow: ellipsis;
  overflow: hidden;
  flex-grow: 1;
  white-space: nowrap;
}
</style>

