<template>
    <div class="city">
        <div class="search_wrap">
            <div class="search">
                <i class="fa fa-search"></i>
                <input v-model="city_val" type="text" placeholder="输入城市名">
            </div>
<!--            <button @click="$router.go(-1)">取消</button>-->
            <button @click="$router.push({name:'theAddress',params:{city:city}})">取消</button>
        </div>
        <div style="height: 100%" v-if="searchList.length === 0">
            <div class="location">
                <location :address="city"></location>
            </div>
            <alphabet :citiesInfo="citiesInfo" @selectCity= "selectCity" ></alphabet>
        </div>
        <div class="search_list" v-else>
            <ul>
              <li v-for="(item,index) in searchList" :key="index" @click="selectCity(item)">
                {{item.name}}
              </li>
            </ul>
        </div>
    </div>
</template>

<script>
import location from '@/components/location'
import alphabet from '@/components/alphabet'
import { reqCitiesData } from '@/api'

export default {
    name:'city',
    data () {
        return {
            city_val: "",
            citiesInfo: [],
            keys: [],
            citiesObj: [],
            searchList: []
        }
    },
  computed: {
    city() {
      return (
          this.$store.getters.location.addressComponent.city ||
          this.$store.getters.location.addressComponent.province
      );
    }
  },
    components: {
        location,
        alphabet
    },
    // 钩子 加载城市信息
    created () {
        this.getCityInfo ()
    },
    watch: {
      city_val () {
        //console.log(this.city_val)
        //查询city
        this.searchCity()
      }
    },
    methods: {
      async getCityInfo () {
       const result = await reqCitiesData()
       if (result.code === 0) {
           this.citiesInfo = result.data
          result.data.forEach(key => {
            //console.log(key.index)
            this.keys.push(key.index)
            if (key.index !== 'hot') {
              key.list.forEach(k => {
                this.citiesObj.push(k)
              })
            }
          })
         this.keys.shift();
         //console.log( this.citiesObj)
       }
      },
      selectCity(city) {
        //console.log(city)
        this.$router.push({name:'theAddress',params:{city:city.name}})
      },
      searchCity() {
        if (!this.city_val) {
          this.searchList = []
        } else {
          // 检索城市 存入 searchList
          this.searchList = this.citiesObj.filter(item => {
            return item.name.indexOf(this.city_val) !== -1
          })
        }
        //console.log(this.searchList)
      }
    }
}
</script>

<style scoped>
.city {
  width: 100%;
  height: 100%;
  overflow: auto;
  box-sizing: border-box;
  padding-top: 45px;
}
.search_wrap {
  position: fixed;
  top: 0;
  height: 45px;
  width: 100%;
  background: #fff;
  box-sizing: border-box;
  padding: 3px 16px;
  display: flex;
  justify-content: space-between;
}
.search {
  background-color: #eee;
  border-radius: 10px;
  line-height: 40px;
  width: 85%;
  box-sizing: border-box;
  padding: 0 16px;
}
.search input {
  background: #eee;
  outline: none;
  border: none;
  margin-left: 5px;
}
.search_wrap button {
  outline: none;
  color: #009eef;
}

.location {
  background: #fff;
  padding: 8px 16px;
  height: 65px;
  box-sizing: border-box;
}

.search_list {
  background: #fff;
  padding: 5px 16px;
}
.search_list li {
  padding: 10px;
  border-bottom: 1px solid #eee;
}
</style>
