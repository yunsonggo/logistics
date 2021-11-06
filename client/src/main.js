import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import axios from 'axios'
import MintUI from 'mint-ui'
import 'mint-ui/lib/style.css'
// 引入vue-amap
import VueAMap from 'vue-amap';
import { Indicator } from 'mint-ui';
import animated from 'animate.css'

Vue.use(VueAMap);
// 初始化vue-amap
VueAMap.initAMapApiLoader({
    // 高德的key
    key: '79ad201c45fb6a945959a5ef35be770b',
    // 插件集合
    plugin: ['AMap.Autocomplete', 'AMap.PlaceSearch', 'AMap.Scale', 'AMap.OverView', 'AMap.ToolBar', 'AMap.MapType', 'AMap.PolyEditor', 'AMap.CircleEditor','AMap.Driving'],
    // 高德 sdk 版本，默认为 1.4.4
    v: '1.4.4'
});

Vue.use(animated)
Vue.use(MintUI)
Vue.config.productionTip = false
//Vue.prototype.$axios = axios
// 请求拦截
axios.interceptors.request.use(config => {
    // if (config.method == 'POST') {
    //     config.data = qs.stringify(config.data)
    // }
    //加载动画
    Indicator.open();
    return config;
},error => {
    return Promise.reject(error);
})
// 响应拦截
axios.interceptors.response.use(response => {
    Indicator.close();
    return response;
},error => {
    Indicator.close();
    return Promise.reject(error);
})

new Vue({
    router,
    store,
    render: h => h(App)
  }).$mount('#app');
