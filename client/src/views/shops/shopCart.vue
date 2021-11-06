<template>
    <div class="shopcart">
        <transition name="fade">
            <!-- 购物车 蒙版 -->
            <div class="cartview-cartmask" @click.self = "showCartView = false" v-if="showCartView && !isEmpty">
                <!-- 购物车 内容 -->
                <div class="cartview-cartbody">
                    <!-- 购物车 头部 -->
                    <div class="cartview-cartheader">
                        <span>已选商品</span>
                        <button>
                            <i class="fa fa-trash-p"></i>
                            <span @click="clearGoods">清空</span>
                        </button>
                    </div>
                    <!--  -->
                    <div class="entityList-cartbodyScroller">
                        <ul class="entityList-cartlist">
                            <li class="entityList-entityrow" v-for="(good,index) in selectGoods" :key="index">
                                <h4><span>{{good.name}}</span></h4>
                                <span class="entityList-entitytotal">{{(good.price * good.count).toFixed(2)}}</span>
                                <cartControll :food="good"></cartControll>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </transition>
        <!-- 购物车 底部 -->
        <div class="bottomNav-cartfooter" :class="{'bottomNav-carticon-empty':isEmpty}">
            <span class="bottomNav-carticon" @click="showCartView = !showCartView">
                <i class="fa fa-cart-plus"></i>
                <span v-if="totalCount">{{totalCount}}</span>
            </span>
            <div class="bottomNav-cartInfo" @click="showCartView = !showCartView">
                <p class="bottomNav-carttotal">
                    <span v-if="isEmpty">未选购商品</span>
                    <span v-else>¥{{totalGoodsPrice.toFixed(2)}}</span>
                </p>
                <p v-if="shopInfo.is_union != 1 " class="bottomNav-cartdelivery">
                    基本配送费:¥{{shopInfo.delivery_price}}元
                </p>
                <p v-else class="bottomNav-notcartdelivery">
                    -配送费:¥{{shopInfo.delivery_price}}元(商家支付)-
                </p>
            </div>
            <button class="submit-btn">
                <span v-if="isEmpty">{{shopInfo.min_num}}件起送</span>
                <span v-else @click="settlement">去结算</span>
            </button>
        </div>
    </div>
</template>

<script>
import cartControll from '@/components/shops/cartControll'
export default {
    name: "shopcate",
    data() {
        return {
            totalCount:0,
            totalGoodsPrice:0,
            selectGoods:[],
            showCartView:false,
            charData : []
        }
    },
    props:{
        cateGoods:Array,
        shopInfo:Object
    },
    created () {
      this.initSelectData()
    },
    computed:{
      isEmpty() {
          let empty = true
          this.totalCount = 0
          this.totalGoodsPrice = 0
          this.selectGoods = []
          let selectedGoods = []
          if (localStorage.getItem("selectedGoods")) {
            selectedGoods = JSON.parse(localStorage.getItem("selectedGoods"))
            this.charData = []
          } else if (this.charData) {
            selectedGoods = this.charData
          }
          if (selectedGoods) {
            this.cateGoods.forEach(item => {
              item.goods.forEach(good => {
                selectedGoods.forEach(ele => {
                  if (ele.id === good.id) {
                    good.count = ele.count
                  }
                })
              })
            })
            localStorage.removeItem("selectedGoods")
            this.charData = []
          }
          this.cateGoods.forEach(item => {
            item.goods.forEach(good => {
                if (good.count) {
                    empty = false
                    this.totalCount += good.count
                    this.totalGoodsPrice += good.price * good.count
                    this.selectGoods.push(good)
                }
              })
          })
          // 签约用户
          if (this.shopInfo.is_union == 1) {
            if (this.totalCount < this.shopInfo.min_num) {
              empty = true
            } else {
              empty = false
            }
          } else if (localStorage.getItem("orderInfo")) {
             let orderData = JSON.parse(localStorage.getItem("orderInfo"))
             if (orderData.shopInfo.is_union == 1) {
               if (this.totalCount < orderData.shopInfo.min_num) {
                  empty = true
                } else {
                  empty = false
                }
             }
          } 
          return empty
      } 
    },
    methods:{
        initSelectData () {
          let orderData = JSON.parse(localStorage.getItem("orderInfo"))
          if (orderData) {
            this.charData = orderData.selectedGoods
          }
        },
        clearGoods() {
            this.cateGoods.forEach(item => {
                item.goods.forEach(good => {
                    good.count = 0
                })
            })
            localStorage.removeItem("orderInfo")
            localStorage.removeItem("selectedGoods")
            localStorage.removeItem("remarkInfo")
            this.selectGoods = []
            this.$store.dispatch("setOrderInfo",{
              shopInfo: this.shopInfo,
              selectedGoods:this.selectGoods
            })
        },
        settlement() {
          this.$store.dispatch("setOrderInfo",{
            shopInfo: this.shopInfo,
            selectedGoods:this.selectGoods
          })
          localStorage.setItem("orderInfo",JSON.stringify(
            {
              shopInfo: this.shopInfo,
              selectedGoods:this.selectGoods
            }
          ))
          localStorage.setItem("selectedGoods",JSON.stringify(this.selectGoods))
          this.$router.push({
            name:"settlement"
          })
        }
    },
    components:{
        cartControll
    }
}
</script>

<style scoped>
.bottomNav-cartfooter {
  position: fixed;
  right: 0;
  bottom: 0;
  left: 0;
  display: flex;
  align-items: center;
  padding-left: 21.066667vw;
  background-color: rgba(61, 61, 63, 0.9);
  height: 12.8vw;
  z-index: 1000;
}

.bottomNav-carticon {
  position: absolute;
  left: 3.2vw;
  bottom: 2vw;
  width: 13.333333vw;
  height: 13.333333vw;
  box-sizing: border-box;
  border-radius: 100%;
  background-color: #3190e8;
  border: 1.333333vw solid #444;
  box-shadow: 0 -0.8vw 0.533333vw 0 rgba(0, 0, 0, 0.1);
}
.bottomNav-carticon > i {
  position: absolute;
  top: 7px;
  right: 0;
  bottom: 0;
  left: 7px;
  color: #fff;
  font-size: 1.6rem;
}
.bottomNav-cartInfo {
  flex: 1;
}
.bottomNav-carttotal {
  font-size: 0.8rem;
  line-height: normal;
  color: #fff;
}
.bottomNav-cartdelivery {
  color: rgb(211, 211, 211);
  font-size: 0.6rem;
}
.bottomNav-notcartdelivery {
  font-style:oblique;
  color: rgb(211, 211, 211);
  font-size: 0.6rem;
  text-decoration: line-through
}
.submit-btn {
  height: 100%;
  width: 28vw;
  background-color: #38ca73;
  color: #fff;
  text-align: center;
  text-decoration: none;
  font-size: 0.9rem;
  font-weight: 600;
  line-height: 12.8vw;
  border: none;
  outline: none;
}

.bottomNav-carticon-empty > span {
  background-image: radial-gradient(circle, #363636 6.266667vw, #444 0);
}
.bottomNav-carticon-empty > span > i {
  color: #606065 !important;
}
.bottomNav-carticon-empty .bottomNav-carttotal > span {
  color: #999;
}
.bottomNav-carticon-empty .submit-btn {
  background-color: #535356 !important;
}

.bottomNav-carticon > span {
  position: absolute;
  right: -1.2vw;
  top: -1.333333vw;
  line-height: 1;
  background-image: linear-gradient(-90deg, #ff7416, #ff3c15 98%);
  color: #fff;
  border-radius: 3.2vw;
  padding: 0.833333vw 1.333333vw;
  font-size: 0.6rem;
}

.cartview-cartmask {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-color: rgba(0, 0, 0, 0.4);
  z-index: 200;
}
.cartview-cartbody {
  position: fixed;
  left: 0;
  width: 100%;
  background-color: #fff;
  bottom: 12.8vw;
  z-index: 201;
  opacity: 1;
  font-size: 1rem;
}
.cartview-cartheader {
  display: flex;
  align-items: center;
  padding: 0 4vw;
  border-bottom: 0.133333vw solid #ddd;
  background-color: #eceff1;
  color: #666;
  height: 10.666667vw;
}
.cartview-cartheader > span {
  flex: 1;
  display: flex;
  align-items: center;
}
.cartview-cartheader > button {
  border: none;
  outline: none;
  flex: none;
  display: flex;
  align-items: center;
  padding-left: 4vw;
  color: #666;
  text-decoration: none;
  font-size: 0.8rem;
  line-height: 4vw;
  background: none;
}
.cartview-cartheader > button > span {
  margin-left: 0.8vw;
}
.entityList-cartbodyScroller {
  overflow: auto;
  max-height: 80vw;
}
.entityList-entityrow {
  border-bottom: 0.133333vw solid #eee;
  display: flex;
  align-items: center;
  padding: 2vw 3.333333vw 2vw 0;
  min-height: 12.666667vw;
  margin-left: 3.333333vw;
}
.entityList-entityrow > h4 {
  flex: 5.5;
  line-height: normal;
}
.entityList-entityrow > h4 > span {
  display: inline-block;
  font-style: normal;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  vertical-align: middle;
  max-width: 46.666667vw;
}
.entityList-entitytotal {
  color: rgb(255, 83, 57);
  flex: 2.5;
  text-align: left;
  white-space: nowrap;
  font-weight: 700;
}
.entityList-entitytotal::before {
  content: "\A5";
  font-size: 0.6rem;
  color: currentColor;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s ease-out;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
