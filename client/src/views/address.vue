<template>
    <div class="address">
        <hader :isLeft="true" title="选择收获地址"></hader>
        <div class="city_search">
          <div class="search">
            <span class="city" @click="$router.push('/city')">
              <i class="fa fa-angle-down">
                {{city}}
              </i>
            </span>
            <i class="fa fa-search">
              <input type="text" v-model="search_val" placeholder="小区/学校/单位...">
            </i>
          </div>
          <location :address="address" @click="selectAddress"></location>
        </div>
        <div class="area" v-if="search_val !== ''">
          <ul class="area_list" v-for="(item,index) in areaList" :key="index">
              <li @click="selectAddress(item)">
                <h4>{{item.name}}</h4>
                <p>{{item.district}}{{item.address}}</p>
              </li>
          </ul>
        </div>
        <div class="amap-wrapper" v-else>
          <el-amap  class="amap-box" vid="amap"  :zoom="zoom"  :center="center" :plugin="theMapPlugin" >
          </el-amap>
        </div>
      <div>
        {{address}}
      </div>
    </div>
</template>

<script>
import hader from '../components/header'
import location from '../components/location'
import distance from '@/api/distance'

export default {
    name:"theAddress",
    data () {
      return {
        city: "",
        search_val: "",
        areaList: [],
        zoom:16,
        center:[115.058494,35.760257],
        theMapPlugin:[{
          enableHighAccuracy: true,//是否使用高精度定位，默认:true
          timeout: 100,          //超过10秒后停止定位，默认：无穷大
          maximumAge: 0,           //定位结果缓存0毫秒，默认：0
          convert: true,           //自动偏移坐标，偏移后的坐标为高德坐标，默认：true
          showButton: true,        //显示定位按钮，默认：true
          buttonPosition: 'RB',    //定位按钮停靠位置，默认：'LB'，左下角
          showMarker: true,        //定位成功后在定位到的位置显示点标记，默认：true
          showCircle: true,        //定位成功后用圆圈表示定位精度范围，默认：true
          panToLocation: true,     //定位成功后将定位到的位置作为地图中心点，默认：true
          zoomToAccuracy:true,    //定位成功后调整地图视野范围使定位位置及精度范围视野内可见，默认：f
          extensions:'all',
          pName: 'Geolocation',
          // events: {
            //init(o) {
              // o 是高德地图定位插件实例
            // o.getCurrentPosition((status, result) => {
            //     console.log(result)
            //     if (result && result.position) {
            //       self.lng = result.position.lng;
            //       self.lat = result.position.lat;
            //       self.center = [self.lng, self.lat];
            //       self.loaded = true;
            //       self.$nextTick();
            //     }
            //   });
            // }
          // }
        }],
      }
    },
    computed: {
      address () {
        //console.log(this.$store)
        return this.$store.getters.location.formattedAddress
      }  
    },
    watch: {
      search_val () {
        this.searchPlace()  
      }
    },
    methods: {
      searchPlace() {
        //console.log(this.search_val);
        const self = this
        AMap.plugin('AMap.Autocomplete', function(){
        // 实例化Autocomplete
        var autoOptions = {
          //city 限定城市，默认全国
          city: self.city
        }
        var autoComplete= new AMap.Autocomplete(autoOptions);
          autoComplete.search(self.search_val, function(status, result) {
            // 搜索成功时，result即是对应的匹配数据
            //console.log(result);
            self.areaList = result.tips
          })
        })
      },
      selectAddress(item) {
        let self = this
        // 设置点选地址
        if (item) {
          this.$store.dispatch('setAddress',item.district + item.address + item.name)
          this.$store.dispatch("setPosition",item.location)
          distance(item.location.lng,item.location.lat,self)
        } else {
          this.$store.dispatch('setAddress',this.address)
        }
        this.$router.push('/home')
      }
    },
    components: {
        hader,
        location,
        distance
    },

  beforeRouteEnter (to, from, next) {
      //console.log(to);
      next(vm => {
        vm.city = to.params.city
      })
    }
}
</script>

<style scoped>
.address {
  width: 100%;
  height: 100%;
  overflow: auto;
  box-sizing: border-box;
  padding-top: 45px;
}

.city_search {
  background-color: #fff;
  padding: 10px 20px;
  color: #333;
}

.search {
  background-color: #eee;
  height: 40px;
  border-radius: 10px;
  box-sizing: border-box;
  line-height: 40px;
}
.search .city {
  padding: 0 10px;
}
.city i {
  margin-right: 10px;
}
.search input {
  margin-left: 5px;
  background-color: #eee;
  outline: none;
  border: none;
}

.area {
  margin-top: 16px;
  background: #fff;
}
.area li {
  border-bottom: 1px solid #eee;
  padding: 8px 16px;
  color: #aaa;
}
.area li h4 {
  font-weight: bold;
  color: #333;
  margin-bottom: 5px;
}
.amap-wrapper {
  height: calc(100% - 180px) ;
  width: 100%;
}
/* .amap-box {

} */
</style>
