<template>
    <div class="myAddress">
        <hader :isLeft="true" :title="title"></hader>
        <!-- 显示地址列表 -->
        <div class="address-view">
          <div class="address-card" v-for="(item,index) in allAddress" :key="index">
              <div class="address-card-select">
                <i class="fa fa-check-circle" v-if="selectedIndex == index"></i>
              </div>
              <div class="address-card-body" @click="setAddressInfo(item,index)">
                <p class="address-card-title">
                  <span v-if="item.user_name" class="username">{{item.user_name}}</span>
                  <span v-else class="username">收货人</span>
                  <span v-if="item.user_gender" class="gender">{{item.user_gender}}</span>
                  <span class="phone">{{item.mobile}}</span>
                </p>
                <p class="address-card-address">
                  <span v-if="item.tage" class="tag">{{item.tage}}</span>
                  <span class="address-text">{{item.address}}{{item.name}}{{item.bottom}}</span>
                </p>
              </div>
              <!-- 操作按钮 -->
              <div class="address-card-edit">
                <i class="fa fa-edit" @click="handleEdit(item)"></i>
                <i class="fa fa-close" @click="handleClose(item)"></i>
              </div>
          </div>
        </div>
        <!-- 新增收货地址 -->
        <div class="address-view-bottom" @click="addAddress">
            <i class="fa fa-plus-circle"></i>
            <span>新增收货地址</span>
        </div>
    </div>
</template>

<script>
import hader from "../../components/header"
import {reqUserAddress,reqDelUserAddress} from "../../api/index"
import { Toast} from 'mint-ui'

export default {
    name: "myAddress",
    data () {
        return {
            title:"我的地址",
            userData :{},
            userId:Number,
            allAddress:[],
            selectedIndex:0
        }
    },
    created () {
      if (this.$route.params.userData) {
        this.userData = this.$route.params.userData
      }
      if (this.$route.query.id) {
        this.userId = this.$route.query.id
      }
    },
    components:{
        hader,
        reqDelUserAddress,
        reqUserAddress,
        Toast
    },
    beforeRouteEnter (to, from, next) {
      // 请求地址数据
      next(vm => vm.getData())
    },
    methods:{
        addAddress() {
            this.$router.push({
              name:"addAddress",
              params:{
                  title: "添加地址",
                  userData:this.userData,
                  addressInfo: {
                    id:Number,
                    userId:Number,
                    tage:"",
                    gender:"",
                    address:"",
                    lng:Number,
                    lat:Number,
                    name:"",
                    phone:"",
                    bottom:""
                  }
                },
                query:{
                  id:this.userId
                }
              })
        },
        async getData () {
          let user_id = localStorage.getItem("user_login")
          const result = await reqUserAddress(user_id)
          if (result.code === 0) {
            this.allAddress = result.data
          }
          //console.log(this.allAddress);
        },
        async delData(item,index) {
          let user_id = localStorage.getItem("user_login")
          const result = await reqDelUserAddress(user_id,item)
          if (result.code === 0) {
            this.showMsg("删除地址成功")
            this.allAddress.splice(index,1)
          } else {
            this.showMsg("删除地址失败")
            return
          }
        },
        handleEdit(item) {
          this.$router.push({
            name:"addAddress",
            query:{
              id:this.userId 
            },
            params:{
              title:"编辑收货地址",
              addressInfo:item
            }
          })
        },
        handleClose(item,index) {
          this.delData(item,index)        
        },
        setAddressInfo(item,index) {
          this.selectedIndex = index
          // 选择的address存入 vuex
          localStorage.removeItem("selectedAddress")
          this.$store.dispatch("setSelectedAddress",item)
          localStorage.setItem("selectedAddress",JSON.stringify(item))
          if (!this.$store.getters.orderInfo || !localStorage.getItem("orderInfo")) {
            this.$router.push({
              name:"me"
            })
          }else {
            this.$router.push({
              name:"settlement"
            })
          }
          
        },
         // 输入框检查 mini UI 的 toast
        showMsg(msg) {
          Toast({
            message: msg,
            position: 'middle',
            duration: 2000
          });
        }
    }
}
</script>

<style scoped>
.myAddress {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  padding-top: 45px;
}

.address-view-bottom {
  position: fixed;
  height: 13.866667vw;
  bottom: 0;
  width: 100%;
  background: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
  border-top: 0.266667vw solid #ddd;
  color: #3190e8;
  font-size: 1rem;
}
.address-view-bottom > i {
  font-size: 1.3rem;
  margin-right: 8px;
}

.address-view {
  height: 100%;
  overflow-y: auto;
  padding-bottom: 13.866667vw;
}
.address-card {
  background-color: #fff;
  padding: 4vw;
  border-bottom: 1px solid #ddd;
  display: flex;
  min-height: 18.133333vw;
}
.address-card-body {
  flex: 1;
  overflow: hidden;
}
.address-card-title {
  font-size: 1rem;
  color: #666;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  padding-bottom: 1.066667vw;
}
.address-card-title .username {
  color: #333;
  font-weight: 700;
}
.address-card-title .gender {
  padding: 0 1.6vw 0 0.8vw;
}
.address-card-address {
  color: #666;
  font-size: 0.84rem;
  display: flex;
  align-items: center;
}
.address-card-address .tag {
  display: inline-block;
  position: relative;
  margin-right: 0.8vw;
  padding: 0.533333vw;
  color: #ff5722;
  white-space: nowrap;
  border: 1px solid #ff5722;
  border-radius: 0.133333vw;
  font-size: 0.6rem !important;
  line-height: 2.666667vw;
}
.address-text {
  line-height: 4.533333vw;
}

/* 编辑和删除 */
.address-card-edit {
  flex-basis: 13.066667vw;
  display: flex;
  justify-content: space-around;
  align-items: center;
}
.address-card-edit > i {
  color:steelblue;
  font-size: 1.2rem;
}

/* .address-card-close {
  flex-basis: 13.066667vw;
  display: flex;
  justify-content: space-around;
  align-items: center;
}
.address-card-close > i {
  color: #aaa;
  font-size: 1.2rem;
} */

/*  选中图标 */
.address-card-select {
  flex-basis: 7.333333vw;
  min-width: 7.333333vw;
  display: flex;
  align-items: center;
}
.address-card-select > i {
  font-size: 1.5rem;
  color: #4cd964;
}
</style>
