<template>
    <!-- <transition class="animate__animated animate__bounceInUp"> -->
      <transition 
        enter-active-class="animate__animated animate__fadeInUp animate__faster" 
        leave-active-class="animate__animated animate__fadeOutDown animate__faster"
      >
        <div class="food" v-if="showGoodPage">
          <div class="foodpanel-close" @click="$emit('closeGoodPage')">
            <img :src="static_url + '/asset/images/closexx.png'" alt="">
          </div>
          <div class="foodpanel-body">
              <div class="foodpanel-foodimg">
                <img :src="static_url + pageData.icon" alt="">
              </div>
              <div class="foodpanel-foodinfo">
                <h4>{{pageData.name}}</h4>
                <div class="foodpanel-foodsales">
                  <span>销量: {{pageData.sell_count}}</span>
                  <span>评分: {{pageData.score}}</span>
                </div>
                <div class="foodpanel-priceLine">
                  <span>¥{{pageData.price}}</span>
                  <cartControll :food = "pageData" class="cart-btn"></cartControll>
                </div>
                <p>{{pageData.content}}</p>
              </div>
          </div>
        </div>
    </transition>
</template>

<script>
import global from '../../components/global'
import cartControll from '@/components/shops/cartControll'

export default {
    name: "goodpage",
    props: {
        pageData: Object,
        showGoodPage: Boolean
    },
    data () {
      return {
        static_url:global.static_url,
      }
    },
    components:{
      global,
      cartControll
    }
}
</script>

<style scoped>
.food {
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  z-index: 100;
}
.foodpanel-close {
  width: 7.466667vw;
  height: 7.466667vw;
  border-radius: 50%;
  background-color: rgba(0, 0, 0, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
  position: absolute;
  z-index: 1;
  top: 2.133333vw;
  right: 4vw;
}
.foodpanel-close img {
  width: 5.333333vw;
  height: 5.333333vw;
}
.foodpanel-body {
  position: relative;
  background-color: #fff;
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  overflow-y: auto;
  padding-bottom: 20vw;
  box-sizing: border-box;
  overflow-x: hidden;
}
.foodpanel-foodimg {
  width: 100%;
  height: 100vw;
  display: block;
}
.foodpanel-foodimg img {
  width: 100%;
  height: 100%;
}
.foodpanel-foodinfo {
  padding: 4vw 4vw 0;
  width: 100%;
  min-height: 29.333333vw;
}
.foodpanel-foodinfo h4 {
  display: flex;
  align-items: center;
  margin-bottom: 2.4vw;
  font-size: 1.2rem;
  font-weight: 700;
  width: 74.666667vw;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.foodpanel-foodsales {
  font-size: 0.8rem;
  color: #666;
  line-height: 1;
  margin-bottom: 2.4vw;
}
.foodpanel-priceLine {
  display: flex;
  align-items: center;
  margin-bottom: 2.4vw;
  position: relative;
}
.foodpanel-priceLine > span {
  font-size: 1rem;
  line-height: 4.266667vw;
  color: #ff5339;
  padding-bottom: 0.933333vw;
  display: flex;
  align-items: baseline;
}
.foodpanel-priceLine .cart-btn {
  position: absolute;
  right: 8vw;
}
.foodpanel-foodinfo > p {
  font-size: 0.8rem;
  color: #666;
  line-height: 3.733333vw;
  padding-right: 8vw;
}

/* .move-enter-active,
.move-leave-active {
  transition: all 0.55s ease-in-out;
  transform: translate(0, -500px);
}

.move-enter,
.move-leave-to {
  transform: translate(0, 500px);
} */
</style>
