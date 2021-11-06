<template>
  <el-row class="page">
    <el-col :span="24" style="position: absolute">
      <home-header
        :open-nav="openNav"
        :infoError="infoError"
        :managerInfo="managerInfo"
        @toggle-open="toggleOpen"
      ></home-header>
    </el-col>
    <el-col :span="24" class="page-main">
      <home-sidebar :openNav="openNav" :menuList="menuList"></home-sidebar>
      <section
        class="page-content"
        :class="{ 'page-content-hide-aside': !openNav }"
      >
        <home-main></home-main>
        <home-footer></home-footer>
      </section>
    </el-col>
  </el-row>
</template>

<script>
import HomeHeader from "@/views/home/HomeHeader";
import HomeSidebar from "@/views/home/HomeSidebar";
import HomeMain from "@/views/home/HomeMain";
import HomeFooter from "@/views/home/HomeFooter";
import {
  reqManagerInfo,
  reqMenuList,
  reqCategroy,
  reqGiveList,
  reqSalePromote,
} from "../api/index";

export default {
  name: "home",
  data() {
    return {
      openNav: true,
      infoError: "",
      managerInfo: {},
      menuList: [],
    };
  },
  // beforeRouteEnter(to,from,next) {
  //   next(vm => {
  //     vm.getManagerInfo()
  //     vm.getMenuList()
  //   })
  // },
  mounted() {
    this.getManagerInfo();
    this.getMenuList();
    this.getCateGroupData();
    this.getGiveData();
    this.getSalePromoteData();
  },
  methods: {
    toggleOpen() {
      this.openNav = !this.openNav;
    },
    // 获取用户数据
    async getManagerInfo() {
      let managerToken = localStorage.getItem("manager_token");
      let managerId = localStorage.getItem("manager_id");
      const result = await reqManagerInfo(managerToken, managerId);
      if (result.code === 1) {
        this.infoError = result.msg;
        this.$message.error(result.msg);
        if (managerToken) {
          localStorage.removeItem("manager_id");
          localStorage.removeItem("manager_token");
        }
        if (result.msg == "需要登录") {
          this.$router.push("/yunsong/login");
        } else if (result.msg == "需要注册") {
          this.$router.push("/yunsong/signup");
        }
      } else if (result.code === 0) {
        this.managerInfo = result.data;
        this.$store.dispatch("setManager", this.managerInfo);
        //console.log(this.managerInfo);
      } else {
        (this.infoError = ""), (this.managerInfo = {});
      }
    },
    // 获取菜单数据
    async getMenuList() {
      let managerToken = localStorage.getItem("manager_token");
      const result = await reqMenuList(managerToken);
      if (result.code === 1) {
        this.menuList = [];
        return;
      }
      result.data.forEach((ele) => {
        if (ele.parent_id === 0) {
          switch (ele.manager_power) {
            case "777":
              if (this.managerInfo.power == 777) {
                this.menuList.push(ele);
              }
              break;
            case "77":
              if (
                this.managerInfo.power == 777 ||
                this.managerInfo.power == 77
              ) {
                this.menuList.push(ele);
              }
              break;
            case "7":
              this.menuList.push(ele);
              break;
          }
        }
      });
      if (this.menuList.length > 0) {
        this.menuList.forEach((ele) => {
          ele.children = [];
          result.data.forEach((item) => {
            if (item.parent_id === ele.id) {
              ele.children.push(item);
            }
          });
        });
      }
      this.$store.dispatch("setMenuData", this.menuList);
      //console.log(this.menuList);
    },
    // 获取大类分类数据
    async getCateGroupData() {
      const result = await reqCategroy();
      if (result.code === 0) {
        this.$store.dispatch("setCategroy", result.data);
      }
    },
    // 获取所有服务数据
    async getGiveData() {
      let managerToken = localStorage.getItem("manager_token");
      const result = await reqGiveList(managerToken);
      if (result.code === 0) {
        this.$store.dispatch("setGiveData", result.data);
      }
    },
    // 获取所有活动数据
    async getSalePromoteData() {
      let managerToken = localStorage.getItem("manager_token");
      const result = await reqSalePromote(managerToken);
      if (result.code == 0) {
        this.$store.dispatch("setSalePromote", result.data);
      }
    },
  },
  components: {
    "home-header": HomeHeader,
    "home-sidebar": HomeSidebar,
    "home-footer": HomeFooter,
    "home-main": HomeMain,
  },
};
</script>


<style scoped>
.page {
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  width: 100%;
}
.page .page-main {
  box-sizing: border-box;
  padding-top: 60px;
  height: 100%;
}
.page .page-content {
  overflow: auto;
  margin-left: 240px;
  height: 100%;
  background-color: #ebeef5;
}
.page .page-content-hide-aside {
  margin-left: 65px;
}
</style>
