<template>
  <div class="index">
    <router-view></router-view>
    <tabbar :data="tabbarData"></tabbar>
  </div>
</template>

<script>
import tabbar from '../components/tabbar'

export default {
  name: 'index',
  data () {
    return {
      tabbarData: [
        {title:'首页',icon:'home',path:'/home'},
        {title:'订单',icon:'file-text-o',path:'/order'},
        {title:'我的',icon:'user',path:'/me'},
      ],
    }
  },
  components: {
    tabbar
  },
}
</script>

<style scoped>
  .index {
    width: 100%;
    height: calc(100% - 45px);
  }
</style>