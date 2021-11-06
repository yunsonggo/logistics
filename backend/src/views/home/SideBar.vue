<template functional>
  <el-submenu  v-on="listeners" :key="data.key" :index="data.key">
    <template slot="title">
      <i :class="data.attrs.subMenu.icon" v-if="data.attrs.subMenu.icon"></i>
      <span slot="title">{{data.attrs.subMenu.menu_title}}</span>
    </template>
    <template v-for="item in data.attrs.subMenu.children">
      <el-menu-item v-if="!item.children" :key="item.menu_title" :index="item.menu_path">
        <a-icon :type="item.icon" v-if="item.icon"/>
        <span>{{item.menu_title}}</span>
      </el-menu-item>
      <the-submenu v-else :subMenu="item" :key="item.menu_title" ></the-submenu>
    </template>
  </el-submenu>
</template>

