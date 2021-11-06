<template>
    <div class="search">
        <hader :isLeft="true" title="搜索">返回</hader>
        <div class="search_header">
            <form action="" class="search_wrap">
                <i class="fa fa-search"></i>
                <input type="text" v-model="key_word" placeholder="输入商家/商品信息" @input="clearResult">
                <button @click.prevent="searchHandle">搜索</button>
            </form>
        </div>
        <div class="shop" v-if="result && !showShop">
          <div v-if="empty" class="empty_wrap">
            <img src="https://fuss10.elemecdn.com/d/60/70008646170d1f654e926a2aaa3afpng.png" alt="">
            <div class="empty_txt">
              <h4>没有搜索结果</h4>
              <span>换个关键词试试吧</span>
            </div>
          </div>
          <div v-else>
            <searchIndex @listClick="shopItemClick" :data="result" ></searchIndex>
          </div>
        </div>
        <!-- 商家信息 -->
        <div class="container" v-if="showShop">
          <!--    菜单导航-->
          <filterview :menuData="menuRes" @update="update"></filterview>
          <div class="shoplist" 
               v-infinite-scroll="loadMore"
               :infinite-scroll-disabled="loading"
          >
            <indexShop v-for="(item,index) in supplierData" :key="index" :restaurant = "item"></indexShop>
          </div>
        </div>
    </div>
</template>

<script>
import hader from '../components/header'
import searchIndex from '../components/searchIndex'
import {reqSearch,reqWithKeyword,reqProShop,reqMenu} from '../api'
import filterview from '@/components/filterview'
import indexShop from '@/components/indexShop'

export default {
    name: 'search',
    data () {
        return {
            key_word : "",
            result:null,
            empty:false,
            showShop: false,
            menuRes: null,
            page: 0,
            size: 7,
            keys:'is_main',
            condition:"",
            supplierData:[],
            restaurant: Object,
            loading:false
        }
    },
    watch : {
        key_word () {
            this.keyWordChange()
            this.AwithKeyWord()
        }
    },
    created() {
      //this.loadData()
      this.getMenu()
    },
    methods: {
       // 菜单
      async getMenu() {
        const result = await reqMenu()
        if (result.code === 0) {
          this.menuRes = result.data
          //console.log(this.menuRes)
        }
      },
      // 点击搜索
      searchHandle() {
        if(!this.key_word) return;
        // 有内容
        if(this.result && this.result.length > 0) {
          this.shopItemClick()
        } else {
          this.empty = true
        }
      },
      //请求后台 获取搜索商家
        async keyWordChange() {
            const result = await reqSearch(this.key_word)
            if (result.data) {
              this.result = result.data
            } 
        },
      // 获取相关搜索词
        async AwithKeyWord() {
          const result = await reqWithKeyword('b2c','utf-8',1,this.key_word)
          result.result.forEach(ele => {
            this.result.push(ele)
          });
          //console.log(this.result);
        },
        clearResult () {
          this.empty = false
          this.result = []
          this.showShop = false
        },
        // 点击搜索列表 跳转实现
        shopItemClick() {
          this.showShop = true
          this.supplierData = this.result
        },
        // 更新数据
        update(key,condition) {
          if (key) {
            this.keys = key.key
            this.condition = condition.condition
            console.log(this.keys);
            console.log(this.condition)
            this.loadData()
          }else {
            this.supplierData = this.result
          }
        },
        async loadData() {
          this.page = 1
          const result = await reqProShop(this.keys,this.page,this.size,this.condition)
          if (result.code === 0 ) {
            this.supplierData = result.data
            //console.log(this.supplierData);
          }
        },
        async loadMore() {
            this.page ++ ;
            const result = await reqProShop(this.keys,this.page,this.size,this.condition)
            if (result.data !== null) {
              // 加载完成 重新 渲染
              result.data.forEach(item => {
                this.supplierData.push(item)
              })
          }
        },
    },
    components: {
        hader,
        searchIndex,
        filterview,
        indexShop
    }
}
</script>



<style scoped>
.search {
  width: 100%;
  height: 100%;
  overflow: auto;
  box-sizing: border-box;
}
.search_header {
  margin-top: 45px;
  background: #fff;
  padding: 0 4.266667vw;
}
.search_header .search_wrap {
  padding: 2.933333vw 2.933333vw 2.933333vw 0;
  display: flex;
  align-items: center;
  position: relative;
}
.search_wrap .fa-search {
  width: 2.933333vw;
  height: 2.933333vw;
  color: #888;
  position: absolute;
  top: 4.6666666vw;
  left: 2.933333vw;
}
.search_wrap input {
  flex-grow: 1;
  border-radius: 0.266667vw;
  background-color: #f8f8f8;
  padding: 1.733333vw 4vw 1.733333vw 8.8vw;
  color: #666;
  outline: none;
  border: none;
}
.search_wrap button {
  outline: none;
  border: none;
  color: #333;
  font-size: 0.426667rem;
  background: #fff;
  font-weight: 700;
  margin-left: 2.933333vw;
  font-size: 14px;
}

.shop {
  width: 100%;
  height: calc(100% - 95px);
  overflow: auto;
}

.empty_wrap {
  width: 100%;
  height: 100%;
  overflow: hidden;
  background: #fff;
  display: flex;
  justify-content: center;
}
.empty_wrap img {
  width: 35vw;
  height: 35vw;
}
.empty_txt h4 {
  color: #666;
  font-size: 1rem;
  margin: 12vw 0 2vw 0;
}
.empty_txt span {
  color: #999;
  font-size: 0.8rem;
}
</style>
