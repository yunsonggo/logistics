<template>
  <section class="index-container" v-if="restaurant.name">
    <div class="index-shopInfo" @click="$router.push({name:'goods',query:{id:restaurant.id}})">
      <!-- 左侧图片 -->
      <div class="logo_container">
        <img :src="static_url + restaurant.icon" alt>
      </div>
      <!-- 右侧内容 -->
      <div class="index_main">
        <!-- 第一行 品牌 -->
        <div class="index_shopname">
          <i v-if="restaurant.is_union">签约</i>
          <span>{{restaurant.name}}</span>
        </div>

        <!-- 第二行 星级 -->
        <div class="index-rateWrap">
          <div>
            <rating :rating = "restaurant.rating"></rating>
            <span class="rate">{{restaurant.rating}}</span>
            <span>销量{{restaurant.total_count}}单</span>
          </div>
          <div v-if="restaurant.sale_id" class="delivery">
            <span class="icon-hollow">正在促销</span>
          </div>
        </div>

        <!-- 第三行 配送 -->
        <div class="index-moneylimit">
          <div>
            <span>{{restaurant.min_num}}件起送</span>
            |
            <span v-if="restaurant.is_union == 1" class="notdelivery">- 配送费: ¥{{restaurant.delivery_price}} -</span>
            <span v-else>配送费: ¥{{restaurant.delivery_price}}</span>
          </div>
          <div class="index-distanceWrap">
            <span v-if="theDistanceMe">距离: {{theDistanceMe.toFixed(2)}}km</span>
            <span v-if="restaurant.is_union == 1 " class="icon-union">推荐</span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import rating from '../components/rating'
import global from '../components/global'
export default {
  name: "indexShop",
  data () {
    return {
      static_url:global.static_url,
    }
  },
  props: {
    restaurant: Object,
    theDistanceMe:""
  },
  
  components: {
    rating,
    global
  }
};
</script>

<style scoped>
.index-container {
  background: #fff;
  color: #666;
  padding: 4vw 0;
  border-bottom: 0.133333vw solid #eee;
}
.index-shopInfo {
  display: flex;
  justify-content: flex-start;
  padding: 0 2.666667vw;
  align-items: stretch;
}
.logo_container {
  width: 17.333333vw;
  height: 17.333333vw;
}
.logo_container img {
  display: block;
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  border: 0.133333vw solid rgba(0, 0, 0, 0.08);
  border-radius: 0.533333vw;
}
.index_main {
  display: flex;
  justify-content: space-between;
  overflow: hidden;
  flex-direction: column;
  padding-left: 2.666667vw;
  font-size: 0.66rem;
  flex-grow: 1;
}
.index_shopname {
  align-items: center;
  color: #333;
  font-weight: 700;
  font-size: 0.9rem;
}
.index_shopname i {
  background: #ffe800;
  margin-right: 1.333333vw;
  padding: 0.266667vw 0.666667vw;
  text-align: center;
  white-space: nowrap;
  font-size: 0.6rem;
}
.index_shopname span {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.index-rateWrap {
  display: flex;
  align-items: center;
  overflow: hidden;
  justify-content: space-between;
}

.index-rateWrap .rate {
  margin-right: 1.066667vw;
}
.index-moneylimit {
  width: 100%;
  display: flex;
  justify-content: space-between;
}
.index-moneylimit .notdelivery {
  font-style:oblique;
  text-decoration: line-through
}
.index-moneylimit .index-distanceWrap {
  color: #999;
}
.index-moneylimit .index-distanceWrap .icon-union{
  color: #fff;
  background-color: #f07373;
  padding: 1px;
  box-sizing: border-box;
}
.delivery {
  display: flex;
  align-items: center;
  font-size: 0.6rem;
  margin-left: 1.066667vw;
}
.delivery .icon-hollow {
  color: #fff;
  background-color: #2395ff;
  padding: 2px;
  box-sizing: border-box;
}
</style>
