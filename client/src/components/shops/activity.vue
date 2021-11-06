<template>
    <div>
        <div class="rst-activity">
            <div class="activity-txt">
                <span style="background: #41b883">服务</span>
                <span> 商家提供免费服务项目: </span>
            </div>
            <div class="activity-count" @click="showGive = !showGive">
                <span>{{giveData.length}}项</span>
                <i class="fa fa-caret-down"></i>
            </div>
            <transition name="giveFade">
                <div class="act_model" v-show="showGive">
                    <div class="activity-sheet">
                        <i class="fa fa-remove" @click="showGive = false"></i>
                        <h2>商家服务</h2>
                        <ul>
                            <li v-for="(item,index) in giveData" :key="index">
                                <span style="background: #41b883">{{item.name}}</span>
                                <span> {{item.desc}} </span>
                            </li>
                        </ul>
                    </div>
                </div>
            </transition>
        </div>
        <div class="rst-activity">
            <div class="activity-txt">
                <span style="background: #f07373">优惠</span>
                <span> 商家商品参与优惠活动: </span>
            </div>
            <div class="activity-count" @click="showSale = !showSale">
                <span>{{saleData.length}}项</span>
                <i class="fa fa-caret-down"></i>
            </div>
            <transition name="saleFade">
                <div class="act_model" v-show="showSale">
                    <div class="activity-sheet">
                        <i class="fa fa-remove" @click="showSale = false"></i>
                        <h2>优惠活动</h2>
                        <ul>
                            <li v-for="(item,index) in saleData" :key="index">
                                <span style="background: #f07373">{{item.name}}</span>
                                <span> {{item.desc}} </span>
                            </li>
                        </ul>
                    </div>
                </div>
            </transition>
         </div>
    </div>
</template>

<script>
import {reqShopGive,reqShopSale} from '../../api/index'

export default {
    name:"activity",
    data () {
        return {
            giveData:[],
            saleData:[],
            showGive:false,
            showSale:false,
        }
    },
    props:{
        shopId :Number,
        giveIds:String,
        saleIds:String
    },
    created () {
        this.getGive(),
        this.getSale()
    },
    methods: {
        async getGive() {
            const result = await reqShopGive(this.shopId,this.giveIds)
            //console.log(result);
            this.giveData = result.data
        },
        async getSale() {
            const result = await reqShopSale(this.shopId,this.saleIds)
            //console.log(result);
            this.saleData = result.data
        }
    }
}
</script>

<style scoped>
.rst-activity {
  display: flex;
  color: #333;
  margin: 3.2vw auto 0;
  align-items: center;
  font-size: 0.6rem;
  width: 80vw;
}
.rst-activity .activity-txt {
  flex: 1;
  overflow: hidden;
}
.activity-txt span:first-child {
  color: #fff;
  height: 4.466667vw;
  line-height: 3.466667vw;
  font-size: 0.6rem;
  padding: 0.533333vw 1.2vw;
  margin-right: 1.6vw;
  display: inline-block;
  box-sizing: border-box;
  border-radius: 0.266667vw;
}
.activity-count {
  width: 16.266667vw;
  flex-shrink: 0;
  padding-right: 2.933333vw;
  text-align: right;
  color: #999;
}
.activity-count span {
  margin-right: 1vw;
}
.act-model {
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  background: rgba(0, 0, 0, 0.5);
  z-index: 99;
}
.activity-sheet {
  position: absolute;
  background-color: #f5f5f5;
  box-shadow: 0 -1px 5px 0 rgba(0, 0, 0, 0.4);
  bottom: 0;
  left: 0;
  right: 0;
  z-index: 100;
  padding: 5.333333vw 6.933333vw;
  box-sizing: border-box;
}
.activity-sheet h2 {
  text-align: center;
  font-size: 1.1rem;
  font-weight: 600;
  margin-bottom: 4.133333vw;
}
.activity-sheet ul {
  height: 50.8vw;
  overflow-y: scroll;
}
.activity-sheet ul li {
  margin-bottom: 3.066667vw;
  display: flex;
  font-size: 0.82rem;
  align-items: center;
}
.activity-sheet ul li span:first-child {
  height: 4.266667vw;
  padding: 0.533333vw 1.6vw;
  margin-right: 1.6vw;
  display: inline-block;
  box-sizing: border-box;
  border-radius: 0.266667vw;
  color: #fff;
  white-space: nowrap;
}
.activity-sheet ul li span:last-child {
  line-height: 1.38;
  flex: 1;
}
.act-model .fa-remove {
  position: absolute;
  font-size: 1.4rem;
  right: 2.666667vw;
  top: 2.666667vw;
  z-index: 1002;
  color: #888;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s ease-out;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}

.giveFade-enter-active,
.giveFade-leave-active {
  transition: opacity 0.25s ease-out;
}

.giveFade-enter,
.giveFade-leave-to {
  opacity: 0;
}

.saleFade-enter-active,
.saleFade-leave-active {
  transition: opacity 0.25s ease-out;
}

.saleFade-enter,
.saleFade-leave-to {
  opacity: 0;
}
</style>

