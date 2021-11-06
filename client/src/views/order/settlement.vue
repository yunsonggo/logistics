<template>
    <div class="settlement">
      <!-- header -->
        <hader :isLeft="true" title="确认订单"></hader>
        <!-- 内容 -->
        <div class="view-body" v-if="orderInfo">
            <div>
                <!-- 收货地址 -->
                <section class="cart-address">
                    <p class="title">
                        订单配送至
                        <span class="address-tag" v-if="selectedAddressInfo && selectedAddressInfo.tage">
                            {{selectedAddressInfo.tage}}
                        </span>
                    </p>
                    <p class="address-detail">
                        <span v-if="selectedAddressInfo" @click="checkAddress">
                            {{selectedAddressInfo.address}}{{selectedAddressInfo.bottom}}
                        </span>
                        <span v-else>
                            <span v-if="haveAddress" @click="checkAddress">选择收货地址</span>
                            <span v-else @click="addAddress">新增收货地址</span>
                        </span>
                        <i class="fa fa-angle-right"></i>
                    </p>
                    <h2 v-if="selectedAddressInfo" class="address-name">
                        <span>{{selectedAddressInfo.user_name}}</span>
                        <span v-if="selectedAddressInfo.user_gender">({{selectedAddressInfo.user_gender}})</span>
                        <span class="phone">{{selectedAddressInfo.mobile}}</span>
                        <span class="phone">{{selectedAddressInfo.name}}(附近)</span>
                    </h2>
                </section>
                <!-- 送达时间 -->
                <delivery :orderInfo="orderInfo" :distanceMe="distanceMe"></delivery>
                <!-- 订单内容 -->
                <cartGroup :orderInfo="orderInfo" :totalPrice="totalPrice" :delPrice=" delPrice"></cartGroup>
                <!-- 备注信息 -->
                <section class="checkout-section">
                  <cartItem @click="showWare=true" title="调货换货" :subHead="remarkInfo.ware || '未选择'"></cartItem>
                  <cartItem @click="$router.push('/remark')" title="订单备注" :subHead="remarkInfo.remark || '备注信息'"></cartItem>
                  <cartItem @click="isInvoiceShow=true" title="发票信息" :subHead="remarkInfo.invoice || '不需要开发票'"></cartItem>
                </section>
                <!-- 显示调货换货信息 -->
                <ware :isWareShow="showWare" @close="showWare=false"></ware>
                <invoice :isInvoiceShow="isInvoiceShow" @close="isInvoiceShow=false"></invoice>
            </div>
        </div>
        <!-- 底部 -->
        <footer class="action-bar">
          <span>总计: ¥{{totalPrice.toFixed(2)}}</span>
          <button @click="handlePay()">去支付</button>
        </footer>
    </div>
</template>

<script>
import hader from "../../components/header"
import {reqUserAddress} from "../../api/index"
import delivery from '@/components/orders/delivery'
import cartGroup from '@/components/orders/cartGroup'
import cartItem from '@/components/orders/cartItem'
import ware from '@/components/orders/ware'
import {Toast} from 'mint-ui'
import invoice from '@/components/orders/invoice'
import distance from '@/api/distance'
import global from '@/components/global'

export default {
    name:"settlement",
    data () {
        return {
            userId:Number,
            haveAddress:false,
            allAddress:[],
            showWare:false,
            isInvoiceShow:false,
            showDeliveryPrice:0
        }
    },
    beforeRouteEnter (to, from, next) {
        next(vm => {
          vm.userId = localStorage.getItem("user") 
            if (!vm.selectedAddressInfo) {
                vm.getData()
            }
        })
    },
    computed:{
        selectedAddressInfo () {
          const self = this
          if (this.$store.getters.selectedAddress) {
            localStorage.removeItem("selectedAddress")
            localStorage.setItem("selectedAddress",JSON.stringify(this.$store.getters.selectedAddress))
            distance(this.$store.getters.selectedAddress.lng,this.$store.getters.selectedAddress.lat,self)
            return this.$store.getters.selectedAddress
          } 
          // const self = this
          // if (this.$store.getters.selectedAddress) {
          //   localStorage.removeItem("selectedAddress")
          //   localStorage.setItem("selectedAddress",JSON.stringify(this.$store.getters.selectedAddress))
          //   distance(this.$store.getters.selectedAddress.lng,this.$store.getters.selectedAddress.lat,self)
          //   return this.$store.getters.selectedAddress
          // } else {
          //   return JSON.parse(localStorage.getItem("selectedAddress"))
          // }
        },
        orderInfo () {
          if (this.$store.getters.orderInfo) {
            localStorage.removeItem("orderInfo")
            localStorage.setItem("orderInfo",JSON.stringify(this.$store.getters.orderInfo))
            return this.$store.getters.orderInfo
          } else {
            return JSON.parse(localStorage.getItem("orderInfo"))
          }
          
        },
        totalPrice () {
          return this.$store.getters.totalPrice
        },
        remarkInfo () {
          return this.$store.getters.remarkInfo
        }, 
        distanceMe () {
          return this.$store.getters.distanceMetre
        },
         delPrice () {
          // 基础运费
          let shopDeliveryPrice = 0
         //  超出运费
          let shopOverPrice = 0
         // 基础公里数
         let freeMetre = global.deliMetre
         // 是否签约商铺
         let suppState = 0
         // 运费
         let deliveryPrice = 0
         // 距离
         let metre = 0
          if (this.$store.getters.orderInfo) {
             shopDeliveryPrice = this.$store.getters.orderInfo.shopInfo.delivery_price
             shopOverPrice = this.$store.getters.orderInfo.shopInfo.over_price
             suppState = this.$store.getters.orderInfo.shopInfo.is_union
          }
          if (this.$store.getters.distanceMetre) {
            metre = this.$store.getters.distanceMetre
          }
          if (metre <= freeMetre) {
            deliveryPrice = shopDeliveryPrice
          } else {
            deliveryPrice = ( metre - freeMetre ) * shopOverPrice + shopDeliveryPrice
          }
          this.$store.dispatch('setDeliveryPrice',deliveryPrice)
          if (suppState === 0) {
            // 非签约商家
            this.showDeliveryPrice = deliveryPrice
          } else {
            this.showDeliveryPrice = 0
          }
          return this.showDeliveryPrice
        },
    },
    methods:{
        async getData () {
          let user_id = localStorage.getItem("user_login")
          const result = await reqUserAddress(user_id)
          if (result.code === 0) {
            this.allAddress = result.data
            if (this.allAddress.length > 0) {
                this.haveAddress = true
            } else {
                this.haveAddress = false
            }
          }
         // console.log(this.allAddress);
        },
        checkAddress() {
            this.$router.push({
                name:"myAddress",
                query:{
                    id:this.userId
                }
            })
        },
        addAddress() {
            this.$router.push({
              name:"addAddress",
              params:{
                  title: "添加地址",
                  userData:this.userData,
                  addressInfo: {
                    id:Number,
                    userId:Number,
                    tage:"",
                    gender:"",
                    address:"",
                    lng:Number,
                    lat:Number,
                    name:"",
                    phone:"",
                    bottom:""
                  }
                },
                query:{
                  id:this.userId
                }
              })
        },
        // 支付
        handlePay() {
          if(!this.selectedAddressInfo) {
            Toast({
              message:"请选择收获地址",
              position:"middle",
              duration:2000
            })
            return
          }
          this.$router.push('/pay')
        },
    },
    components:{
        hader,
        reqUserAddress,
        delivery,
        cartGroup,
        cartItem,
        ware,
        Toast,
        invoice,
        distance,
        global
    }
}
</script>


<style scoped>
.settlement {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  padding-top: 44px;
}

.view-body {
  width: 100%;
  height: 100%;
  overflow: auto;
  box-sizing: border-box;
  font-size: 0.9rem;
  color: #333;
  padding-bottom: 14.133333vw;
  padding-left: 1.6vw;
  padding-right: 1.6vw;
  background-image: linear-gradient(
      0deg,
      #f5f5f5,
      #f5f5f5 65%,
      hsla(0, 0%, 96%, 0.3) 75%,
      hsla(0, 0%, 96%, 0)
    ),
    linear-gradient(270deg, #009eef, #009eef);
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}
.cart-address {
  background-color: transparent;
  padding: 4.266667vw 2.133333vw 2.933333vw 2.133333vw;
  min-height: 22.133333vw;
  color: #fff;
  overflow: hidden;
}
.cart-address .title {
  font-size: 0.9rem;
  line-height: 1.21;
  color: hsla(0, 0%, 100%, 0.8);
}
.cart-address .address-detail {
  font-size: 1.3rem;
  font-weight: 600;
  line-height: 1.88;
  overflow: hidden;
  display: flex;
  align-items: center;
}
.address-detail > span {
  display: inline-block;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  max-width: calc(100% - 4vw);
}
.address-detail > i {
  margin-left: 2.133333vw;
}

/* 显示送货地址 */
.address-name {
  font-size: 0.86rem;
  line-height: 1.21;
  margin-bottom: 1.333333vw;
}
.address-name .phone {
  margin-left: 2.133333vw;
}
.address-tag {
  border: 0.133334vw solid hsla(0, 0%, 100%, 0.8);
  margin-left: 1.6vw;
  display: inline-block;
  padding: 0.533333vw;
  white-space: nowrap;
  border-radius: 0.133333vw;
  font-size: 0.6rem !important;
  line-height: 2.666667vw;
}

.checkout-section {
  margin-bottom: 2.133333vw;
  padding: 0 5.333333vw;
  background: #fff;
  box-shadow: 0 0.133333vw 0.266667vw 0 rgba(0, 0, 0, 0.05);
}

/* 底部支付样式 */
.action-bar {
  height: 11.733333vw;
  position: fixed;
  left: 0;
  bottom: 0;
  width: 100%;
  background: #3c3c3c;
  z-index: 2;
}
.action-bar > span {
  color: #fff;
  font-size: 1.2rem;
  line-height: 11.733333vw;
  padding-left: 2.666667vw;
  vertical-align: middle;
}
.action-bar > button {
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  background: #00e067;
  min-width: 28vw;
  padding: 0 1.333333vw;
  border: none;
  color: #fff;
  font-size: 1.2rem;
}
</style>
