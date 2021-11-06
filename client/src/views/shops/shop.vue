<template>
    <div class="shop" v-if="shopInfo">
        <!-- 头 -->
        <nav class="header-nav">
            <div class="nav_bg">
                <img :src="static_url + shopInfo.banner" alt="">
            </div>
            <div class="nav_back">
                <i @click="$router.push('/home')" class="fa fa-chevron-left"></i>
            </div>
            <div class="shop_image">
                <img :src="static_url + shopInfo.icon" alt="">
            </div>
        </nav>
        <!-- 商家信息  -->
        <div class="index-rst">
            <div class="rst-name">
                <span @click="showInfoModel = true">{{shopInfo.name}}</span>
                <i class="fa fa-caret-right"></i>
            </div>
            <!-- 弹窗 -->
            <infoModel @closeInfoModel="showInfoModel = false" :rst="shopInfo" :showInfoModel="showInfoModel" :theDistanceMe="theDistanceMe"></infoModel>
            <!-- 默认信息 -->
            <div class="rst-order">
                <span>评分{{shopInfo.rating}}</span>
                <span>销量{{shopInfo.total_count}}单</span>
                <span v-if="theDistanceMe">距离({{theDistanceMe.toFixed(2)}}KM)</span>
                <span>...</span>
            </div>
            <!-- 优惠信息 -->
            <activity :shopId="shopInfo.id" :giveIds="shopInfo.give_id" :saleIds="shopInfo.sale_id"></activity>
            <!-- 公告 -->
            <p class="rst-promotion">公告: {{shopInfo.description}}</p>
        </div>
        <!-- 导航 -->
        <navbar :shopInfo="shopInfo"></navbar>
        <keep-alive>
          <router-view></router-view>
        </keep-alive>
    </div>
</template>  

<script>
import {reqShopInfo} from '../../api/index'
import infoModel from '../../components/shops/infoModel'
import activity from '../../components/shops/activity'
import navbar from '../../components/shops/navbar'
import global from '../../components/global'

export default {
    name:"shop",
    data() {
        return {
          static_url:global.static_url,
          shopInfo:null,
          showInfoModel:false
        }
    },
    created() {
        this.getData()
    },
    computed: {
      theDistanceMe() {
        return this.$store.getters.distanceMetre
      }
    },
    methods: {
       async getData() {
            const result = await reqShopInfo(this.$route.query.id)
            //console.log(result);
            if (result.code === 0) {
                this.shopInfo = result.data
            }
        }
    },
    components:{
        infoModel,
        activity,
        navbar,
        global
    }
    // beforeRouteEnter (to, from, next) {
    //   //console.log(to);
    //    next(vm => {
    //      vm.shopInfo = to.params.shopInfo
    //    })
    // }
}
</script>

<style scoped>
.shop {
  width: 100%;
  height: 100%;
  overflow: auto;
  box-sizing: border-box;
}
.header-nav {
  position: relative;
}
.nav_bg img {
  width: 100%;
  height: 26.666667vw;
}
.nav_back {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 26.666667vw;
  background: rgba(0, 0, 0, 0.5);
}
.nav_back i {
  color: #fff;
  font-size: 1.3rem;
  margin-left: 1.333333vw;
  margin-top: 1.333333vw;
}
.shop_image {
  position: absolute;
  top: 0;
  left: 50%;
  margin-left: -10vw;
  margin-top: 11vw;
}
.shop_image img {
  width: 20vw;
  height: 20vw;
  border-radius: 0.8vw;
}

.index-rst {
  padding: 8vw 0 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  background: #fff;
  box-shadow: inset 0 -0.666667vw 0.666667vw hsla;
}
.index-rst .rst-name {
  flex: 1;
  width: 72vw;
  /* font-size: 1.3rem; */
  font-size: 1.1rem;
  font-weight: 700;
  white-space: nowrap;
  padding-right: 2.666667vw;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 1.6vw 0;
}
.rst-name span {
  text-align: left;
  overflow: hidden;
  text-overflow: ellipsis;
}

.index-rst .rst-order {
  white-space: nowrap;
  height: 3.2vw;
  margin-top: 1.733333vw;
  color: #666;
  text-align: center;
  font-size: 0.6rem;
}
.rst-order span {
  margin: 0 3px;
}
.index-rst .rst-promotion {
  width: 80vw;
  font-size: 0.6rem;
  color: #999;
  overflow: hidden;
  text-overflow: ellipsis;
  margin: 2.266667vw auto 2.666667vw;
  padding: 0;
  white-space: nowrap;
}
</style>