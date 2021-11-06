<template>
    <div class="comment" v-if="evaluation">
        <!-- 商家评分 -->
        <section class="rating-wrap">
            <div class="rating-info">
                <h4>{{shopInfo.rating}}</h4>
                <div class="shop-score">
                    <span>商家评分</span>
                    <rating :rating = "shopInfo.rating"></rating>
                </div>
            </div>
            <div class="other-score">
                <div class="tp-score">
                    <div>
                        <span>服务</span>
                        <p>{{shopInfo.serve_score}}</p>
                    </div>
                    <div>
                         <span>包装</span>
                        <p>{{shopInfo.pack_score}}</p>
                    </div>
                </div>
                <div class="rider-score">
                    <span>配送</span>
                    <p>{{shopInfo.delivery_score}}</p>
                </div>
            </div>
        </section>
        <!-- 评论区 -->
        <div class="shop-info">
            <!-- 标签 -->
            <ul class="tags">
                <li @click="selectComment(item.key)" v-for="(item,index) in tagesData" :key="index" :class="{'unsatisfied':item.key == 'bag'}">
                    {{item.title}}
                    <span v-if="item.num > 0">{{item.num}}</span>
                </li>
            </ul>
            <!-- 内容 -->
            <div v-for="(item,index) in tagesData" :key="index">
              <transition
                enter-active-class="animate__animated animate__fadeInRight animate__fast" 
                leave-active-class="animate__animated animate__fadeOutLeft animate__fast"
              >
                  <ul v-if="item.key == selectKey" class="comments-wrap">
                      <li v-for="(eva,k) in item.evaData" :key="k">
                          <div class="comment-user">
                              <img v-if="eva.user_icon" :src="static_url + eva.user_icon" alt="">
                              <img v-else src="../../assets/avatar.jpg" alt="">
                          </div>
                          <div class="comments-info">
                              <div class="comment-name">
                                  <h4>{{eva.user_name}}</h4>
                                  <small>{{eva.create_time}}</small>
                              </div>
                              <div class="comment-rating">
                                  <rating :rating="eva.score"></rating>
                                  <span :style="{color: ratingcontent(eva.score.toFixed(0)).color}">
                                      {{ratingcontent(eva.score.toFixed(0)).txt}}
                                  </span>
                              </div>
                              <div class="comment-text">
                                  {{eva.content_text}}
                              </div>
                              <div class="comment-reply">
                                  {{eva.reply}}
                              </div>
                              <ul class="comment-imgs" v-if="eva.icon">
                                  <!-- <li v-for="(img,i) in eva.icon" :key = "i">
                                      <img :src="static_url + img" alt="">
                                  </li> -->
                                  <li>
                                      <img :src="static_url + eva.icon" alt="">
                                  </li>
                              </ul>
                          </div>
                      </li>
                  </ul>
              </transition>
            </div>       
        </div>
    </div>
</template>

<script>
import {reqGoodEva,reqShopInfo} from '../../api/index'
import rating from '../../components/rating'
import global from '../../components/global'
import {formatDate} from '@/components/formatDate.js'

export default {
    name:"comments",
    data () {
        return {
            static_url:global.static_url,
            urlId:Number,
            evaluation:null,
            idName:"shop_id",
            shopId:Number,
            page: 1,
            count: 20,
            shopInfo:Object,
            tagesData:[],
            selectKey:"all"
        }
    },
    props:{
        
    },
    created() {
        this.getUrlId(),
        this.getShopInfo(),
        this.getData()
    },
    methods:{
        getUrlId() {
            if (this.$route.query.id) {
                this.urlId = parseInt(this.$route.query.id)
            } 
        },
        // 获取商家信息
        async getShopInfo() {
            const result = await reqShopInfo(this.$route.query.id)
            //console.log(result);
            if (result.code === 0) {
                this.shopInfo = result.data
            }
        },
        // 获取评论 数据
        async getData() {
            const result = await reqGoodEva(this.idName,this.urlId,this.page,this.count)
            if (result.code == 0) {
                this.evaluation = result.data
                this.handleEva()
            }
            //console.log(result);
        },
        // 分类评论数据
        handleEva() {
            var tages = [
                    {key:"all", title: "全部", num: 0,evaData:[]},
                    {key:"new",  title: "最新", num: 0,evaData:[]},
                    {key:"fine",  title: "好评", num: 0,evaData:[]},
                    {key:"general",  title: "中评", num: 0,evaData:[]},
                    {key:"bag",  title: "差评", num: 0,evaData:[]},
                    {key:"pic",  title: "有图", num: 0,evaData:[]},
                ];
            this.evaluation.forEach(ele => {
                tages.forEach(tage => {
                    switch (tage.key) {
                        case "all": 
                            tage.evaData.push(ele)
                            tage.num ++
                        case "new": 
                            if (tage.evaData.length > 5) {
                                break
                            }
                            tage.evaData.push(ele)
                            tage.num ++
                        case "fine": 
                            if (ele.is_fine == 1) {
                                tage.evaData.push(ele)
                                tage.num ++
                                break
                            }
                        case "general": 
                            if (ele.is_general == 1) {
                                tage.evaData.push(ele)
                                tage.num ++
                                break
                            }
                        case "bag": 
                            if (ele.is_bad == 1) {
                                tage.evaData.push(ele)
                                tage.num ++
                                break
                            }
                        case "pic":
                            if (ele.icon) {
                                tage.evaData.push(ele)
                                tage.num ++
                                break
                            }
                        default: 
                            break
                    }
                })
            })
            this.tagesData = tages
            this.selectKey = "all"
            //console.log(this.tagesData);
        },
        // 根据 标签 筛选评论
        selectComment(key) {
            this.selectKey = key
        },
        ratingcontent(rating) {
            const content = [
                { txt: "吐槽", color: 'rgb(137,159,188)' },
                { txt: "较差", color: 'rgb(137, 159, 188)' },
                { txt: "一般", color: 'rgb(251, 154, 11)' },
                { txt: "满意", color: 'rgb(251, 154, 11)' },
                { txt: "超赞", color: 'rgb(255, 96, 0)' }
            ];
            return content[rating - 1];
        }
    },
    components:{
        reqGoodEva,
        reqShopInfo,
        rating,
        global,
        formatDate
    },
    filters: {
        // 时间格式自定义 只需把字符串里面的改成自己所需的格式
        formatDate(time) {
            let date = new Date(time);
            return formatDate(date, 'yyyy年MM月dd日');
        },
        // formatDate2(time) {
        //     let date = new Date(time);
        //     return formatDate(date, 'hh:mm:ss');
        // },
        // formatDate3(time) {
        //     let date = new Date(time);
        //     return formatDate(date, 'yyyy年MM月dd日 hh:mm:ss');
        // }
    }
}
</script>

<style scoped>
.comment {
  height: calc(100% - 10.666667vw);
  line-height: 1.2;
}
.rating-wrap {
  display: flex;
  margin-bottom: 2.133333vw;
  padding: 5.333333vw 0 8vw 6.4vw;
  background-color: #fff;
}
.rating-info {
  display: flex;
  justify-content: space-between;
  width: 33.6vw;
  align-items: center;
  color: #666;
}
.rating-info h4 {
  align-items: center;
  font-size: 2.2rem;
  color: #ff6000;
}
.rating-info .shop-score {
  align-items: flex-start;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.other-score {
  flex: 1;
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: #666;
}
.tp-score {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  flex: 1;
  padding: 0 7.2vw;
  align-items: center;
}
.tp-score > div {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.tp-score > div > span,
.rider-score > span {
  font-size: 0.8rem;
  margin-bottom: 1.333333vw;
}
.tp-score > div > p,
.rider-score > p {
  font-size: 1rem;
}
.rider-score {
  width: 22.933333vw;
  border-left: 1px solid #ddd;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.shop-info {
  background-color: #fff;
  padding: 2.666667vw 3.2vw 0;
  font-size: 0.8rem;
}
.tags {
  padding-bottom: 2.666667vw;
  border-bottom: 1px solid #eee;
}
.tags li {
  color: #6d7885;
  background-color: #ebf5ff;
  display: inline-block;
  padding: 0 2.4vw;
  height: 7.466667vw;
  line-height: 7.466667vw;
  margin: 0.933333vw;
  font-size: 0.7rem;
  border-radius: 0.533333vw;
}
.tags li:hover{
  color: #dddddd;
  background-color: #ff6000;
  display: inline-block;
  padding: 0 2.4vw;
  height: 7.466667vw;
  line-height: 7.466667vw;
  margin: 0.933333vw;
  font-size: 0.7rem;
  border-radius: 0.533333vw;
  opacity: 0.6;
}

.unsatisfied {
  color: #aaa !important;
  background-color: #f5f5f5 !important;
}

.unsatisfied:hover {
  color: #dddddd !important;
  background-color: #ff6000 !important;
  opacity: 0.6;
}

.comments-wrap > li {
  padding: 4vw 0 3.2vw;
  border-bottom: 0.133333vw solid #eee;
  display: flex;
}
.comment-user {
  width: 10.666667vw;
}
.comment-user img {
  width: 30px;
  height: 30px;
  border-radius: 50%;
}
.comments-info {
  flex: 1;
  font-size: 0.9rem;
}
.comment-name {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.comment-name h4 {
  margin-top: 0;
  color: #333;
  margin-right: 1.6vw;
}
.comment-name small {
  font-size: 0.6rem;
  color: #999;
}
.comment-rating {
  display: flex;
  margin: 1.6vw 0 0.533333vw;
  align-items: center;
}
.comment-rating > span {
  font-size: 0.6rem;
  margin-left: 1.066667vw;
}
.comment-text {
  color: #333;
  word-break: break-word;
  margin: 2.133333vw 0;
}
.comment-reply {
  margin: 2.666667vw 0;
  padding: 2.666667vw;
  border-radius: 1.066667vw;
  background: #f3f3f3;
  position: relative;
  font-size: 0.8rem;
}
.comment-reply::after {
  content: " ";
  position: absolute;
  bottom: 100%;
  left: 4vw;
  width: 0;
  height: 0;
  border-style: solid;
  border-width: 0 2.133333vw 2.133333vw;
  border-color: transparent transparent #f3f3f3;
}
.comment-imgs {
  margin-top: 1.066667vw;
  margin-bottom: 3.2vw;
}
.comment-imgs > li {
  display: inline-block;
  margin: 0 0.533333vw;
}
.comment-imgs > li img {
  width: 40vw;
  height: 40vw;
}
</style>
