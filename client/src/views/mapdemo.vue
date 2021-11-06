<template>
  <div class="amap-wrapper">
    <el-amap  class="amap-box" vid="map"  :zoom="zoom" :center="center">
    </el-amap>
  </div>
</template>

<script>
export default {
    name: 'mapdemo',
    data () {
      return {
        zoom:16,
        center:[
          121.406051,
          31.179695],
      }
    }
}
</script>
<style>
.amap-wrapper {
  height: 100%;
  width: 100%;
}
</style>
