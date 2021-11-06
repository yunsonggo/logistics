<template>
    <div v-if="cateGoods" class="goods">
        <!-- 商家推荐  is top -->
        <div class="recommend">
            <p class="recommend-name">商家推荐</p>
            <div class="recommend-wrap">
                <ul>
                    <li v-for="(item,index) in cateGoods[cateGoods.length-1].goods" :key="index">
                        <img :src="static_url + item.icon" alt="">
                        <div class="recommend-food">
                            <p class="recommend-food-name">{{item.name}}</p>
                            <p class="recommend-food-zm">销量:{{item.sell_count}} 评分:{{item.score}}</p>
                        </div>
                        <div class="recommend-food-price">
                            <p>¥{{item.price}}</p>
                            <cartControll :food="item"></cartControll>
                        </div>
                    </li>
                </ul>
            </div>
        </div>
        <!-- 商品分类 -->
        <div class="menuview">
            <!-- 左侧分类列表 -->
            <div class="menu-wrapper" ref="menuScroll">
                <ul>
                    <li :class="{'current':currentIndex === index}" @click="selectMenu(index)" v-for="(item,index) in cateGoods" :key="index">
                        <img v-if="item.sc_icon" :src="static_url + item.sc_icon" alt="">
                        <span>{{item.sc_name}}</span>
                    </li>
                </ul>
            </div>
            <!-- 右侧商品内容 -->
            <div class="foods-wrapper" ref="foodsScroll">
                <ul>
                    <li  class="food-list-hook" v-for="(item,index) in cateGoods" :key="index">
                        <!-- 内容上 标题 -->
                        <div class="category-title">
                            <strong>{{item.sc_name}}</strong>
                            <span>{{item.desc}}</span>
                        </div>
                        <!-- 内容下 详情 -->
                        <div @click="handelShopPage(val)" class="fooddetails" v-for="(val,key) in item.goods" :key="key">
                            <!-- 图片 -->
                            <img :src="static_url + val.icon" alt="">
                            <!-- 文字 -->
                            <section class="fooddetails-info">
                                <h4>{{val.name}}</h4>
                                <p class="fooddetails-des">{{val.description}}</p>
                                <p class="fooddetails-sales">销量:{{val.sell_count}} 评分:{{val.score}}</p>
                                <div class="fooddetails-price">
                                    <span class="price">¥{{val.price}}</span>
                                    <span v-if="val.old_price" class="price-line-through">-¥{{val.old_price}}-</span>
                                    <cartControll :food="val"></cartControll>
                                </div>
                            </section>
                        </div>
                    </li>
                </ul>
            </div>
        </div>
        <!--  购物车 -->
        <shopcart :cateGoods="cateGoods" :shopInfo="shopInfo"></shopcart>
        <!-- 商品详情 -->
        <goodpage :showGoodPage = "showGoodPage" :pageData = "pageData" @closeGoodPage = "showGoodPage = false"></goodpage>
    </div>
</template>

<script>
import {reqShopInfo,reqShopAllGoods,reqShopCates} from '../../api/index'
import global from '../../components/global'
import cartControll from '@/components/shops/cartControll'
import BScroll from 'better-scroll'
import shopcart from './shopCart'
import goodpage from '@/views/shops/goodpage'

export default {
    name:"goods",
    data () {
        return {
            static_url:global.static_url,
            urlId:Number,
            shopInfo:null,
            cateGoods:null,
            menuScroll:{},
            goodScroll:{},
            showGoodPage:false,
            pageData:{},
            scrollY:0,      // 右侧商品 滚动Y 值
            listHeight:[]   // 左侧分类 高度值
        }
    },
    computed: {
        // 根据右侧Y 位置 确定 对应的 索引分类下标
        currentIndex() {
            for(let i=0; i<this.listHeight.length; i++ ) {
                // 相邻两个中 上边一个
                let heightOver = this.listHeight[i]
                // 相邻两个中 下边一个
                let heightDown = this.listHeight[i+1]
                // 判断是否在两个 中间
                if (this.scrollY >= heightOver && this.scrollY < heightDown ) {
                    return i
                }
            }
            return 0
        }
    },
    created() {
        this.getUrlId(),
        this.getShopData(),
        this.getShopCates()
    },
    methods:{
        getUrlId() {
            if (this.$route.query.id) {
                this.urlId = parseInt(this.$route.query.id)
            } else if (this.shopInfo) {
                this.urlId = parseInt(this.shopInfo.id)
            }
        },
        // 获取商家信息 有待优化
        async getShopData() {
          localStorage.removeItem("orderInfo")
            const result = await reqShopInfo(this.urlId)
            //console.log(result);
            if (result.code === 0) {
                this.shopInfo = result.data
            }
        },
        // 获取商家商品分类 
        async getShopCates() {
            const result = await reqShopCates(this.urlId)
            if (result.code === 0) {
                var topGoods = {"id":9999}
                result.data.push(topGoods)
                result.data.forEach(ele => {
                    ele.goods = []
                })
                this.getGoodsData(result.data)
            }
        },
        // 获取商家商品 有待优化
        async getGoodsData(cateData) {
            const result = await reqShopAllGoods(this.urlId)
           // console.log(result);
            if (result.code === 0) {
                result.data.forEach(ele => {
                    ele.count = 0
                });
                this.handleData(cateData,result.data)
            }
        },
        handleData(cateData,goodData) {
            cateData.forEach(cate => {
                goodData.forEach(good => {
                    if (good.cate_id === cate.id && good.is_top != 1) {
                        cate.goods.push(good)
                    }else if (cate.id === 9999 && good.is_top == 1) {
                        cate.goods.push(good)
                    }
                })
            })
            this.cateGoods = cateData
            //console.log(this.cateGoods);
            this.$nextTick(() => {
                   //DOM 已更新 页面渲染完毕 调用 初始化 better-scroll插件
                   this.initScroll()
                   //计算 各个 分类 的高度 调用此方法
                   this.calculateHeight()
               })
        },
        // 初始化 better-scroll    
        initScroll() {
            this.menuScroll = new BScroll(this.$refs.menuScroll,{
                click:true
            })
            this.goodScroll = new BScroll(this.$refs.foodsScroll,{
                probeType:3,
                click:true
            })
            this.goodScroll.on("scroll",pos => {
                //console.log(pos.y);
                this.scrollY = Math.abs(Math.round(pos.y)) // 四舍五入 绝对值  当前Y 值
            })
        },
        selectMenu(index) {
            //console.log(index);
           let goodlist =  this.$refs.foodsScroll.getElementsByClassName("food-list-hook");
           let ele = goodlist[index]
           //console.log(ele);
           // 滚动到对应元素位置
           this.goodScroll.scrollToElement(ele,250)
        },
        //计算 高度 
        calculateHeight() {
            let foodlist = this.$refs.foodsScroll.getElementsByClassName("food-list-hook")
            //console.log(foodlist);
            // 每个分类区高度添加进数组
            let height = 0
            this.listHeight.push(height)
            for(let i=0; i<foodlist.length-1; i++) {
                let item = foodlist[i];
                //累加
                height += item.clientHeight;
                this.listHeight.push(height)
            }
           // console.log(this.listHeight);
        },
        //商品详情
         handelShopPage(val) {
           //console.log(val);
           this.pageData = val
           this.showGoodPage = true
         }
    },
    components:{
        global,
        cartControll,
        BScroll,
        shopcart,
        goodpage
    }
}
</script>

<style scoped>
.goods {
  position: relative;
  height: calc(100% - 10.666667vw);
}

.recommend {
  padding-top: 4.266667vw;
  background-color: #fff;
}
.recommend-name {
  padding-left: 4.266667vw;
  color: #333;
  font-size: 1rem;
  font-weight: 700;
  margin-bottom: 2.666667vw;
}
.recommend-wrap {
  overflow-x: scroll;
  display: flex;
  width: 100%;
}
.recommend-wrap ul {
  display: flex;
  padding-left: 4.266667vw;
}
.recommend-wrap ul li {
  flex: none;
  width: 32vw;
  margin-right: 2.666667vw;
}
.recommend-wrap li img {
  display: block;
  width: 32vw;
  height: 32vw;
  border-top-left-radius: 0.8vw;
  border-top-right-radius: 0.8vw;
  max-width: 100%;
}
.recommend-food .recommend-food-name {
  color: #333;
  font-size: 0.8rem;
  margin: 1.866667vw 0 1.233333vw;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
}
.recommend-food .recommend-food-zm {
  color: #999;
  font-size: 0.6rem;
  margin-bottom: 1.866667vw;
  min-height: 1em;
}
.recommend-food-price {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-right: 0.266667vw;
}
.recommend-food-price p {
  font-size: 1rem;
  color: #ff5339;
}

.recommend-food-price .price-line-through {
  font-style:oblique;
  font-size: 1rem;
  line-height: 4.266667vw;
  color: #dadada;
  padding-bottom: 0.933333vw;
  display: flex;
  align-items: baseline;
  text-decoration: line-through
}



::-webkit-scrollbar {
  width: 0 !important;
}

.menuview {
  box-sizing: border-box;
  height: 100%;
  padding-bottom: 10.8vw;
  background-color: #fff;
  display: flex;
}
.menu-wrapper {
  overflow-y: hidden;
  /* height: 100%; */
  height: calc(100% - 12.8vw);
  background-color: #eeeeee;
  padding-bottom: 10.666667vw;
  width: 20.533333vw;
}
.menu-wrapper li {
  padding: 4.666667vw 2vw;
  font-size: 0.6rem;
  color: #666;
  line-height: 1.2;
}
.menu-wrapper li img {
  max-width: 100%;
  width: 3.466667vw;
  height: 3.466667vw;
  vertical-align: top;
  margin-right: 0.8vw;
}

.foods-wrapper {
  overflow-y: hidden;
  /* height: 100%; */
  height: calc(100% - 12.8vw);
  width: 79.466667vw;
  padding-bottom: 10.666667vw;
}
.category-title {
  margin-left: 2.666667vw;
  padding: 2vw 8vw 2vw 0;
  display: flex;
  align-items: center;
  overflow: hidden;
}
.category-title strong {
  margin-right: 1.333333vw;
  font-weight: 700;
  font-size: 0.8rem;
  color: #666;
  flex: none;
}
.category-title span {
  flex: 1;
  color: #999;
  font-size: 0.6rem;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.fooddetails {
  min-height: 30.666667vw;
  padding: 2.666667vw 0 2.666667vw 2.666667vw;
  margin-bottom: 0.133333vw;
  display: flex;
}
.fooddetails img {
  width: 25.333333vw;
  height: 25.333333vw;
  flex: none;
  margin-right: 2.666667vw;
  border-radius: 0.533333vw;
}
.fooddetails-info {
  flex: 1;
  padding-bottom: 6.666667vw;
  padding-right: 4vw;
}
.fooddetails-info h4 {
  padding-right: 4vw;
  font-weight: 700;
  overflow: hidden;
  font-size: 1rem;
  white-space: nowrap;
  width: 40vw;
  text-overflow: ellipsis;
  color: #333;
}
.fooddetails-des {
  margin: 1.333333vw 0;
  font-size: 0.6rem;
  color: #999;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  width: 42.666667vw;
}
.fooddetails-sales {
  margin: 1.733333vw 0 !important;
  color: #999;
  font-size: 0.6rem;
  line-height: 1;
  min-height: 1em;
}
.fooddetails-price {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 3.733333vw;
}
.fooddetails-price .price {
  font-size: 1rem;
  line-height: 4.266667vw;
  color: #ff5339;
  padding-bottom: 0.933333vw;
  display: flex;
  align-items: baseline;
}

.fooddetails-price .price-line-through {
  font-style:oblique;
  font-size: 1rem;
  line-height: 4.266667vw;
  color: #b8b8b8;
  padding-bottom: 0.933333vw;
  display: flex;
  align-items: baseline;
  text-decoration: line-through
}

.menu-wrapper .current {
  background-color: #fff !important;
  color: #333 !important;
}
</style>
