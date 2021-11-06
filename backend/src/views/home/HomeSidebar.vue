<template>
    <aside class="sidebar" :class="{'sidebar-hide': !openNav}">
        <el-menu :default-active="selectMenu" class="sidebar-menu" :collapse="!openNav"
                :collapse-transition="false" :router="true" v-if="menuList">
          <template v-for="menu in menuList">
            <el-menu-item v-if="!menu.children || menu.children.length === 0" :key="menu.menu_title" :index="menu.menu_path">
              <i :class="menu.icon"></i>
              <span slot="title">{{menu.menu_title}}</span>
            </el-menu-item>
            <the-submenu v-else :key="menu.menu_title" :subMenu="menu"></the-submenu>
          </template>
        </el-menu>
    </aside>
</template>

<script>
import SideBar from '@/views/home/SideBar'

export default {
    name:'homeSidebar',
    props: ['openNav','menuList'],
    computed:{
      selectMenu () {
        //console.log(this.$route.path);
       return this.$route.path
      },
    },
    components:{
      'the-submenu': SideBar
    }
}
</script>

<style scoped lang="scss">
.sidebar {
  float: left;
  width: 240px;
  height: 100%;
  border-right: 1px solid #e6e6e6;
  overflow: auto;

  .sidebar-menu {
    border: none;
    height: 100%;
  }
}

.sidebar-hide {
  width: 65px;
}
</style>