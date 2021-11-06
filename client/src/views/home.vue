<template>
  <div class="home">
    <!-- <mt-loadmore 
      :top-method="loadData" 
      :bottom-method="loadMore" 
      :bottom-all-loaded="allLoaded" 
      :auto-fill = false
      :bottomPullText = "bottomPullText"
      ref="loadmore"> -->
      <div class="header">
          <div class="address_map" @click="$router.push({name:'theAddress',params:{city:city}})">
              <i class="fa fa-map-marker"></i>
              <span>{{address}}</span>
              <i class="fa fa-sort-desc"></i>
          </div>
      </div>
      <div class="search_wrap" :class="{'fixedview':showFilter}" @click="$router.push('/search')">
          <div class="shop_search">
              <i class="fa fa-search"></i>
              <input type="text" v-model="search_val" placeholder="搜索商品/商家...">
          </div>
      </div>

      <div id="container">
<!--        轮播 -->
        <mt-swipe :auto="2000" class="swiper" :show-indicators="false">
          <mt-swipe-item v-for="(good,index) in swipeGoods" :key="index">
            <img :src="static_url + good.icon" alt="">
          </mt-swipe-item>
        </mt-swipe>
<!--        分类 -->
        <div class="entries">
            <div class="foodentry" v-for="(item,i) in catesRes" :key="i" @click="selectByCate(item)">
              <div class="img_wrap">
                <img :src="static_url + item.icon" alt>
              </div>
              <span>{{item.name}}</span>
            </div>
        </div>
      </div>
    <!--    推荐商家-->
    <div class="shoplist-title">推荐商家</div>    
    <!--    菜单导航-->
    <filterview :menuData="menuRes" @searchFixed="showFilterView" @update="update"></filterview>  
    <!-- 商家信息 -->
    <mt-loadmore 
      :top-method="loadData" 
      :bottom-method="loadMore" 
      :bottom-all-loaded="allLoaded" 
      :auto-fill = false
      :bottomPullText = "bottomPullText" 
      ref="loadmore">
      <div v-if="!supplierData" class="shop">
          <div class="empty_wrap">
            <img src="https://fuss10.elemecdn.com/d/60/70008646170d1f654e926a2aaa3afpng.png" alt="">
            <div class="empty_txt">
              <h4>没有搜索结果</h4>
              <span>换个关键词试试吧</span>
            </div>
          </div>
      </div>
      <div class="shoplist" v-else>
        <indexShop v-for="(item,index) in supplierData" :key="index" :restaurant="item" :theDistanceMe="theDistanceMe"></indexShop>
      </div>
    </mt-loadmore>
    
  </div>
</template>

<script>
import {reqProfileShoping, reqCategroy, reqMenu, reqProShop,reqShopByCate} from '@/api'
import filterview from '../components/filterview'
import indexShop from '../components/indexShop'
import global from '../components/global'

export default {
    name: 'home',
    data () {
      return {
        static_url:global.static_url,
        search_val: "",
        swipeGoods: [],
        catesRes: [],
        menuRes: [],
        showFilter:false,
        page: 1,
        size: 10,
        keys:'is_main',
        condition:"",
        supplierData:[],
        allLoaded: false,
        bottomPullText:"上拉加载更多",
      }
    },
  computed: {
      address () {
        return this.$store.getters.address
      },
      city () {
        return (
            this.$store.getters.location.addressComponent.city ||
            this.$store.getters.location.addressComponent.province
        );
      },
      theDistanceMe() {
        return this.$store.getters.distanceMetre
      }
    },
    mounted(){
      this.getProfileShoping();
      this.getProCates();
      this.getMenu();
      this.getProShop() 
    },
  methods: {
      //首页轮播商品
    async getProfileShoping() {
        const result = await reqProfileShoping()
        if (result.code === 0) {
          this.swipeGoods = result.data
        }
        //console.log(this.swipeGoods)
      },
    // 首页类别
    async getProCates() {
      const result = await reqCategroy()
      if (result.code === 0) {
        this.catesRes = result.data
       // console.log(this.catesRes)
      }
    },
    // 首页菜单
    async getMenu() {
      const result = await reqMenu()
      if (result.code === 0) {
        this.menuRes = result.data
        //console.log(this.menuRes)
      }
    },
    // 首页商家信息
    getProShop() {
     this.loadData()
    },
    // 显示排序 筛选
    showFilterView(isShow) {
      //console.log(isShow);
      this.showFilter = isShow
    },
   
    // 更新数据
    update(key,condition) {
      this.keys = key.key
      this.condition = condition.condition
      //console.log(this.keys);
      //console.log(this.condition)
      this.loadData()
    },
    async loadData() {
      this.page = 1
      this.allLoaded = false
      this.bottomPullText = "上拉加载更多"
      const result = await reqProShop(this.keys,this.page,this.size,this.condition)
      if (result.code === 0 ) {
        this.$refs.loadmore.onTopLoaded()
        this.supplierData = result.data
        //console.log(this.supplierData);
      }
    },
    async loadMore() {
      if (!this.allLoaded) {
        this.page ++ ;
        const result = await reqProShop(this.keys,this.page,this.size,this.condition)
        if (result.data !== null) {
          // 加载完成 重新 渲染
           this.$refs.loadmore.onBottomLoaded()
           result.data.forEach(item => {
             this.supplierData.push(item)
           })
           if (result.data < this.size) {
             this.allLoaded = true;
             this.bottomPullText = "没有更多了哦..."
           }
        } else {
            // 数据空
            this.allLoaded = true;
            this.bottomPullText = "没有更多了哦..."
        }
      }
    },
    // 根据类别搜索
    async selectByCate(item) {
      let cateId = item.id 
      const result = await reqShopByCate(cateId)
      if (result.code == 0) {
        this.supplierData = result.data
      }
    }
  },
  components: {
    filterview,
    indexShop,
    global
  }
}
</script>

<style scoped>
  .home {
    width: 100%;
    height: 100%;
    overflow: auto;
    box-sizing: border-box;
  }
  .header,
  .search_wrap {
    background-color: #009eef;
    padding: 10px 16px;
  }
  .header .address_map {
    color: #fff;
    font-weight: bold;
  }
  .address_map i {
    margin: 0 3px;
    font-size: 18px;
  }
  .address_map span {
    display: inline-block;
    width: 80%;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
  }
  .search_wrap .shop_search {
    /*margin-top: 10px;*/
    background-color: #fff;
    padding: 10px 0;
    border-radius: 4px;
    text-align: center;
    color: #aaa;
  }

  .search_wrap {
    position: sticky;
    top: 0px;
    z-index: 999;
    box-sizing: border-box;
  }
  .swiper {
    height: 150px;
  }
  .swiper img {
    width: 100%;
    height: 150px;
  }

  .entries {
    background: #fff;
    height: 47.2vw;
    text-align: center;
    overflow: hidden;
  }
  .foodentry {
    width: 20%;
    float: left;
    position: relative;
    margin-top: 2.933333vw;
  }
  .foodentry .img_wrap {
    position: relative;
    display: inline-block;
    width: 14vw;
    height: 14vw;
  }
  .img_wrap img {
    width: 100%;
    height: 100%;
  }
  .foodentry span {
    display: block;
    color: #666;
    font-size: 0.66rem;
  }
  /* 推荐商家 */
  .shoplist-title {
    display: flex;
    align-items: flex;
    justify-content: center;
    height: 9.6vw;
    line-height: 9.6vw;
    font-size: 16px;
    color: #333;
    background: #fff;
  }
  .shoplist-title:after,
  .shoplist-title:before {
    display: block;
    content: "一";
    width: 5.333333vw;
    height: 0.266667vw;
    color: #999;
  }
  .shoplist-title:before {
    margin-right: 3.466667vw;
  }
  .shoplist-title:after {
    margin-left: 3.466667vw;
  }

  .fixedview {
    width: 100%;
    position: fixed;
    top: 0px;
    z-index: 999;
  }

  .mint-loadmore {
    height: calc(100% - 95px);
    overflow: auto;
  }
</style>
