<template>
  <div :class="{'open':isSort || isUnion || isFilter || isGood}" @click.self = "hideView">
<!--    导航  -->
    <div v-if="menuData" class="filter_wrap">
      <aside class="filter">
        <div class="filter-nav"
             v-for="(item,index) in menuData.MenuList"
             :key="index"
             :class="{'filter-bold':currentFilter === index}"
             @click="filterSort(index)"
        >
          <span>{{item.name}}</span>
          <i v-if="item.icon" :class="'fa fa-' + item.icon"></i>
        </div>
      </aside>
    </div>
<!--    商家排序-->
    <section class="filter-extend" v-if="isSort">
      <ul>
        <li v-for="(item,index) in menuData.Sort" :key="index" @click="selectSort('sort',item,index)">
          <span :class="{'selectName':currentSort === index}">{{item.name}}</span>
          <i class="fa fa-check" v-show="currentSort === index"></i>
        </li>
      </ul>
    </section>
   <!-- 商品排序-->
    <section class="filter-extend" v-if="isGood">
      <ul>
       <li v-for="(item,index) in menuData.MenuGoods" :key="index" @click="selectSort('good',item,index)">
          <span :class="{'selectName':currentSort == index}">{{item.name}}</span>
          <i class="fa fa-check" v-show="currentSort == index"></i>
       </li>
     </ul>
    </section>
      <!-- 联盟商家-->
    <section class="filter-extend" v-if="isUnion">
      <ul>
        <div v-for="(item,index) in menuData.MenuList" :key="index">
          <li v-if="index === 2 " @click="selectSort('union',item,index)">
            <span :class="{'selectName':currentSort == index}">{{item.name}}</span>
            <i class="fa fa-check" v-show="currentSort == index"></i>
          </li>
        </div>
      </ul>
    </section>
    <!-- 商品筛选 -->
    <section class="filter-extend" v-if="isFilter">
      <div class="filter-sort" >
        <div class="morefilter">
          <p class="title">商家服务(多选)</p>
          <ul>
            <li v-for="(item,index) in menuData.MenuGive" :key="index" 
                @click="selectScreen(item,'menuGive',menuData.MenuGive)"
                :class="{'selected':item.select}"
                >
                <img v-if="item.icon" :src="static_url + item.icon" alt="">
                <span>{{item.name}}</span>
            </li>
          </ul>
        </div>
        <div class="morefilter">
          <p class="title">优惠活动(多选)</p>
          <ul>
            <li v-for="(item,index) in menuData.MenuSale" :key="index"
              :class="{'selected':item.select}"
              @click="selectScreen(item,'menuSale',menuData.MenuSale)"
            >
                <img v-if="item.icon" :src="static_url + item.icon" alt="">
                <span>{{item.name}}</span>
            </li>
          </ul>
        </div>
        <div class="morefilter">
          <p class="title">商品价格(单选)</p>
          <ul>
            <li v-for="(item,index) in menuData.MenuPrice" :key="index"
              :class="{'selected':item.select}"
              @click="selectScreen(item,'menuPrice',menuData.MenuPrice)"
            >
                <img v-if="item.icon" :src="static_url + item.icon" alt="">
                <span v-if="item.minValue == '' ">¥{{item.maxValue}}以下</span>
                <span v-else-if="item.maxValue == '' ">¥{{item.minValue}}以上</span>
                <span v-else>¥{{item.minValue}}-{{item.maxValue}}</span>
            </li>
          </ul>
        </div>
      </div>
      <div class="morefilter-btn">
        <mt-button type="default" size="large" @click="clearFilter">清空</mt-button>   
        <mt-button type="primary" size="large" @click="filterDone">确定</mt-button>   
        <!-- <span class="morefilter-clear">清空</span>
        <span class="morefilter-ok">确定</span> -->
      </div>
    </section>
  </div>

</template>

<script>
import global from '@/components/global' 

export default {
  name: "filterview",
  props: {
    menuData: {}
  },
  data () {
    return {
      static_url:global.static_url,
      menuSaveList: {},
      currentFilter: 0, // 点击菜单下标
      isSort:false,
      isGood:false,
      isUnion:false,
      isFilter:false,
      currentSort:0
    }
  },
  components:{
    global
  },
  computed:{
    
  },
  methods: {
    filterSort(index) {
       //console.log(index);
      // 点击加粗
      this.currentFilter = index
      // 点击 添加蒙版
      switch (index) {
        case 0:
          this.isSort = true,
          this.isGood = false,
          this.isUnion = false,
          this.isFilter = false,
          this.$emit("searchFixed",true)
          break;
        case 1:
          this.isSort = false,
          this.isGood = true,
          this.isUnion = false,
          this.isFilter = false,
          this.$emit("searchFixed",true)
          break;
        case 2: 
          this.isSort = false,
          this.isGood = false,
          this.isUnion = true,
          this.isFilter = false,
          this.$emit("searchFixed",true)
          break;
        case 3: 
          this.isSort = false,
          this.isGood = false,
          this.isUnion = false,
          this.isFilter = true,
          this.$emit("searchFixed",true)
          break;
        default:
          this.hideView()
          break;
      }
    },
    // 取消蒙版
    hideView () {
      this.isSort = false;
      this.isUnion = false;
      this.isFilter = false;
      this.isGood = false;
      this.$emit("searchFixed",false)
    },
    // 菜单点选处理
    selectSort(key,item,index) {
      switch (key) {
        case 'sort': 
            this.currentSort = index
            this.menuData.MenuList[0].name = this.menuData.Sort[index].name
            this.hideView()
            // 更新数据
            this.$emit("update",{key:'sort'},{condition:item.value})
          break;
        case 'good': 
            this.currentSort = index
            this.menuData.MenuList[1].name = this.menuData.MenuGoods[index].name
            this.hideView()
            // 更新数据
            this.$emit("update",{key:'good'},{condition:item.value})
          break;
        case 'union': 
            this.currentSort = index
            this.menuData.MenuList[2].name = this.menuData.MenuList[index].name
            this.hideView()
            this.$emit("update",{key:'union'},{condition:item.value})
          break;
        default: 
        this.hideView()
        break;
      }
      
    },
    // 筛选点选处理
    selectScreen(item,key,data) {
      if (key === 'menuPrice') {
        //单选
        data.forEach(ele => {
          ele.select = false;
        })
      }
      item.select = !item.select
    },
    // 清空筛选
    clearFilter() {
      this.menuData.MenuGive.forEach(ele => {
        ele.select = false
      })
      this.menuData.MenuSale.forEach(ele => {
        ele.select = false
      })
      this.menuData.MenuPrice.forEach(ele => {
        ele.select = false
      })
    },
    // 确定 筛选
    filterDone() {
      let requestData = ""
      this.menuData.MenuGive.forEach(ele => {
        if (ele.select === true) {
          requestData += ("g"+ele.id+",")
        }
      
      })
      this.menuData.MenuSale.forEach(ele => {
        if (ele.select === true) {
          requestData += ("s"+ele.id+",")
        }
      })
      this.menuData.MenuPrice.forEach(ele => {
        if (ele.select === true) {
          requestData += ("min"+ele.minValue+","+"max"+ele.maxValue+",") 
        }
      })
      // 向后端发送请求
      this.$emit("update",{key:'screen'},{condition:requestData})
      this.hideView()
    }
  }
}
</script>

<style scoped>
.filter_wrap {
  background: #fff;
  position: sticky;
  top: 54px;
  z-index: 10;
}
.filter {
  position: relative;
  border-bottom: 1px solid #ddd;
  line-height: 10.4vw;
  z-index: 101;
  height: 10.666667vw;
  display: flex;
  justify-content: space-around;
}
.filter-nav {
  flex: 1;
  text-align: center;
  color: #666;
  font-size: 0.8333rem;
}
.filter-nav i {
  width: 1.6vw;
  height: 0.8vw;
  margin-left: 1.333333vw;
  margin-bottom: 0.533333vw;
  fill: #333;
  will-change: transform;
}

.filter-bold {
  font-weight: 600;
  color: #333;
}

.open {
  /* position: fixed; */
  position:fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  transition: all 0.3s ease-in-out;
  z-index: 3;
}

.filter-extend {
  background-color: #fff;
  color: #333;
  padding-top: 2.133333vw;
  /* position: absolute; */
  position: sticky;
  width: 100%;
  z-index: 4;
  left: 0;
  top: 25.533333vw;
  /* top: 24.533333vw; */
}
.filter-extend li {
  position: relative;
  padding-left: 5.333333vw;
  line-height: 10.666667vw;
  overflow: hidden;
}
.fa-check {
  float: right;
  color: #009eef;
  margin-right: 3.733333vw;
  line-height: 10.666667vw;
}

.selectName {
  color: #009eef;
}
/* 筛选 */
.filter-sort {
  background: #fff;
  padding: 0 2.666667vw;
  line-height: normal;
}
.morefilter {
  margin: 2.666667vw 0;
  overflow: hidden;
}
.morefilter .title {
  margin-bottom: 2vw;
  color: #666;
  font-size: 0.5rem;
}
.morefilter ul {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  font-size: 0.8rem;
}
.morefilter li {
  box-sizing: border-box;
  width: 30%;
  height: 9.333333vw;
  line-height: 9.333333vw;
  margin: 0.8vw 1%;
  background: #fafafa;
}
.morefilter li img {
  width: 3.466667vw;
  height: 3.466667vw;
  vertical-align: middle;
  margin-right: 0.8vw;
}
.morefilter li span {
  margin-left: 2%;
  vertical-align: middle;
}

.morefilter-btn {
  display: flex;
  justify-content: space-around;
  align-items: center;
  background-color: #fafafa;
  box-shadow: 0 -0.266667vw 0.533333vw 0 #ededed;
  line-height: 11.466667vw;
  box-sizing: border-box;
}
.morefilter-btn span {
  font-size: 0.826667rem;
  text-align: center;
  text-decoration: none;
  flex: 1;
}
.morefilter-clear {
  color: #ddd;
  background: #fff;
}
.morefilter-ok {
  color: #fff;
  background: #00d762;
  border: 0.133333vw solid #00d762;
}

.selected {
  color: #3190e8 !important;
  background-color: #edf5ff !important;
}

.edit {
  color: #333 !important;
}
</style>
