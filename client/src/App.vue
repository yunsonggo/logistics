<template>
  <div id="app">
    
      <router-view></router-view>
   
  </div>
</template>

<script>

export default {
  name:'app',
  created () {
    this.getLocation();
  },
  methods: {
    getLocation () {
      const self = this
      AMap.plugin('AMap.Geolocation', function() {
        var geolocation = new AMap.Geolocation({
          // 是否使用高精度定位，默认：true
          enableHighAccuracy: true,
          // 设置定位超时时间，默认：无穷大
          timeout: 10000,
        })
        geolocation.getCurrentPosition();
        AMap.event.addListener(geolocation, 'complete', onComplete);
        AMap.event.addListener(geolocation, 'error', onError);
        function onComplete (data) {
          // data是具体的定位信息
          //console.log(data);
          self.$store.dispatch('setLocation',data);
          self.$store.dispatch('setAddress',data.formattedAddress);
          self.distance()
        }
        function onError (data) {
          // 定位出错
          self.getLnglatLocation();
          //console.log(data);
        }
      })
    },
    getLnglatLocation() {
      const self = this;
      AMap.plugin('AMap.CitySearch', function () {
        var citySearch = new AMap.CitySearch();
        citySearch.getLocalCity(function (status, result) {
          if (status === 'complete' && result.info === 'OK') {
            // 查询成功，result即为当前所在城市信息
            // console.log(result);
            // -----------------------
              AMap.plugin('AMap.Geocoder', function() {
                var geocoder = new AMap.Geocoder({
                  // city 指定进行编码查询的城市，支持传入城市名、adcode 和 citycode
                  city: result.adcode
                })

                var lnglat = result.rectangle.split(";")[0].split(",");

                geocoder.getAddress(lnglat, function(status, data) {
                  if (status === 'complete' && data.info === 'OK') {
                      // result为对应的地理位置详细信息
                      self.$store.dispatch('setLocation',{
                        addressComponent: {
                          city: result.city,
                          province: result.province
                        },
                        formattedAddress: data.regeocode.formattedAddress
                      });
                      //
                      self.$store.dispatch("setAddress",data.regeocode.formattedAddress)
                      //console.log(data);
                      
                  }
                })
              })
            // ------------------------
          }
        })
      })
      
    },
    
      distance () {
        const self = this
        let shopLng = 115.058494
        let shopLat = 35.760257
        let endLng = 0
        let endLat = 0
        // 搜索并点选了地址
        if (self.$store.getters.position.lng) {
          endLng = self.$store.getters.position.lng
          endLat = self.$store.getters.position.lat
        } else if (self.$store.getters.location.position) {
          // 默认定位地址
          endLng = self.$store.getters.location.position.lng
          endLat = self.$store.getters.location.position.lat
        } 
        
        AMap.plugin('AMap.Driving', () => {
            // 实例化Autocomplete
            var autoOptions = {
                policy: AMap.DrivingPolicy.LEAST_TIME, // 其它policy参数请参考 https://lbs.amap.com/api/javascript-api/reference/route-search#m_DrivingPolicy
                ferry: 1, // 是否可以使用轮渡
                extensions: 'all',
                province: '豫', // 车牌省份的汉字缩写 
            }
            //var autoComplete= new AMap.Autocomplete(autoOptions);
            var driving = new AMap.Driving(autoOptions)
                // 根据起终点经纬度规划驾车导航路线
            driving.search(new AMap.LngLat(shopLng,shopLat), new AMap.LngLat(endLng,endLat), (status, result) => {
                // result 即是对应的驾车导航信息，相关数据结构文档请参考  https://lbs.amap.com/api/javascript-api/reference/route-search#m_DrivingResult
                if (status === 'complete') {
                    //console.log('绘制驾车路线完成')
                    //console.log(result)
                    //console.log(result.routes[0].distance);
                    //self.distanceRouterData = result
                    //self.theDistance = result.routes[0].distance/1000
                    self.$store.dispatch('setDistancesRouters',result)
                    self.$store.dispatch('setDistanceMetre',result.routes[0].distance/1000)
                } else {
                    console.log('获取驾车数据失败：' + result);
                }
            });
            
        })
        //console.log(self.theDistance);
     }
  },
  
}
</script>

<style>
#app {
  width: 100%;
  height: 100%;
  font-size: 14px;
  background: #f1f1f1;
}
</style>
