export default function  distance (getLng,getLat,self) {
    let shopLng = 115.058494
    let shopLat = 35.760257
    let endLng = getLng
    let endLat = getLat
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
                return result
                   
            } else {
                console.log('获取驾车数据失败：' + result);
                return ""
            }
        });
        
    })
    //console.log(self.theDistance);
 }