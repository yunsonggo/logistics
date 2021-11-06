<template>
  <el-header class="header el-button--primary">
    <router-link to="/yunsong/home">
      <div class="logo" :class="{ 'logo-hide': !openNav }">
        <img src="../../assets/images/homelogo.png" class="image" />
        <span class="text">后台控制中心</span>
      </div>
    </router-link>
    <!-- 头部 折叠图标 -->
    <div class="content">
      <i
        class="el-icon-s-fold toggle"
        @click="navOpenToggle"
        v-show="openNav"
      ></i>
      <i
        class="el-icon-s-unfold toggle"
        @click="navOpenToggle"
        v-show="!openNav"
      ></i>
    </div>
    <div class="right max-right">
      <!-- 头部 信件 -->
      <div class="right-item">
        <i class="el-icon-message" style="font-size: 18px"></i>
        <el-badge :value="1" class="item" />
      </div>
      <!-- 用户类别 -->
      <div class="right-item" v-if="managerInfo">
        <p v-if="managerInfo.power == '7' ">管理员:</p>
        <p v-else-if="managerInfo.power == '77' ">系统管理员:</p>
        <p v-else>超级管理员:</p>
      </div>
      <!-- 用户 信息 -->
      <div class="right-item">
        <el-dropdown triggen="click">
          <p class="user-info" v-if="infoError">
            {{infoError}} <i class="el-icon-s-custom" style="margin-left: 10px"></i>
          </p>
          <p class="user-info" v-else-if="managerInfo.usernick">
            {{managerInfo.usernick}} <i class="el-icon-s-custom" style="margin-left: 10px"></i>
          </p>
          <p class="user-info" v-else>
            {{managerInfo.username}} <i class="el-icon-s-custom" style="margin-left: 10px"></i>
          </p>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item>
              <router-link to="">
                <el-link :underline="false">个人中心</el-link>
              </router-link>
            </el-dropdown-item>
            <el-dropdown-item divided>
              <el-button type="text" @click="signOut">退出登录</el-button>
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
  </el-header>
</template>

<script>
export default {
  name: "homeHeader",
  data() {
    return {};
  },
  props: ["openNav","infoError","managerInfo"],
  methods: {
    navOpenToggle() {
      this.$emit("toggle-open");
    },
    // 退出登陆
    signOut() {
      localStorage.removeItem("manager_id")
      localStorage.removeItem("manager_token")
      this.$router.push("/yunsong/login")
    }
  },
};
</script>


<style scoped lang="scss">
.header {
  color: #ffffff;
  line-height: 60px;
  user-select: none;

  div {
    display: inline-block;
  }

  .logo {
    width: 240px;
    border-right: 1px solid #c0c4cc;
    margin-left: -20px;
    text-align: center;
    font-size: 25px;
    cursor: pointer;

    .image {
      width: 40px;
      height: 40px;
      vertical-align: middle;
    }

    .text {
      color: #ffffff;
    }
  }

  .logo-hide {
    width: 65px;

    .text {
      display: none;
    }
  }

  .content {
    padding: 0 20px;

    .toggle {
      font-size: 22px;
      cursor: pointer;
    }
  }

  .right {
    float: right;

    .right-item {
      display: inline-block;
      padding: 0 10px;
      min-width: 60px;
      text-align: center;
      font-size: 14px;
      cursor: pointer;

      .drop-icon {
        transition: transform 0.2s;
      }
    }

    .right-item:hover {
      background-color: rgba(255, 255, 255, 0.3);
    }

    .user-info {
      color: #ffffff;
    }
  }

  .min-right {
    display: none;
  }
}

@media (max-width: 768px) {
  .header {
    .logo {
      border: none;
      display: contents;

      .text {
        display: inline-block !important;
      }
    }

    .content {
      float: left;
      margin-left: -20px;
    }
  }

  .max-right {
    display: none !important;
  }

  .min-right {
    display: inline-block !important;
  }
}
</style>
