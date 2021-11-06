<template>
    <div class="addAddress">
        <hader :isLeft="true" :title="title"></hader>
        <!-- add  -->
        <div class="viewbody">
            <div class="content">
                <formBlock v-model="addressInfo.name" label="联系人:" placeholder="姓名" :tags="gender" :selectTage="addressInfo.gender" @checkGender="checkGender"></formBlock>
                <formBlock v-model="addressInfo.phone" label="电 话:" placeholder="手机号码"></formBlock>
                <formBlock textarea="textarea" v-model="addressInfo.address" @click="showSearch = true " label="地 址:" placeholder="街道/小区/写字楼/学校" icon="angle-right"></formBlock>
                <formBlock v-model="addressInfo.bottom" label="门牌号:" placeholder="1号楼1单元101" icon="edit" textarea="textarea"></formBlock>
                <div class="formblock">
                    <div class="label-wrap">常用:</div>
                    <tabTage :tages="tages" :selectTage="addressInfo.tage" @checkTage="checkTage"></tabTage>
                </div>
                <formBlock label="自定义:" placeholder="自定义标签" v-model="addressInfo.tage"></formBlock>
                <div class="formblock">
                  <div class="label-wrap">默认:</div>
                  <mt-switch  v-model="addressInfo.position" @change="turn"></mt-switch>
                </div>
                
            </div>
            <!-- 确定按钮 -->
            <div class="form-button-wrap">
              <button class="form-button" @click="handleSaveAddr()">
                确定
              </button>
            </div>
        </div>
        <!-- search -->
        <addressSearch :addressInfo="addressInfo" :showSearch="showSearch" @closeSearch="showSearch = false"></addressSearch>
    </div>
</template>

<script>
import hader from "../../components/header"
import formBlock from '@/components/orders/formBlock'
import tabTage from '@/components/orders/tabTage'
import addressSearch from '@/components/orders/addressSearch'
import { Toast} from 'mint-ui'
import {reqAddUserAddress,reaEditUserAddress} from '../../api/index'


export default {
    name: "addAddress",
    data () {
        return {
            title:"添加地址",
            tages:["总店","分店","中心店","新店"],
            gender:["先生","女士"],
            userData:{},
            sendData:{},
            userId:Number,
            addressInfo:{
              userId:Number,
              tage:"",
              gender:"",
              address:"",
              lng:Number,
              lat:Number,
              name:"",
              phone:"",
              bottom:"",
              addrName:"",
              position:false
            },
            showSearch:false
        }
    },
    beforeRouteEnter (to, from, next) {
      next(vm => {
        if (to.params.userData) {
        vm.userData = to.params.userData
      }
      if (to.query.id) {
        vm.userId = to.query.id
      }
        vm.addressInfo.id = to.params.addressInfo.id
        vm.addressInfo.userId = to.params.addressInfo.user_id
        vm.addressInfo.tage = to.params.addressInfo.tage
        vm.addressInfo.gender = to.params.addressInfo.user_gender
        vm.addressInfo.address = to.params.addressInfo.address
        vm.addressInfo.lng = to.params.addressInfo.lng
        vm.addressInfo.lat = to.params.addressInfo.lat
        vm.addressInfo.name = to.params.addressInfo.user_name
        vm.addressInfo.phone = to.params.addressInfo.mobile
        vm.addressInfo.bottom = to.params.addressInfo.bottom
        if (to.params.addressInfo.position === 0){
           vm.addressInfo.position = true
        } else {
          vm.addressInfo.position = false
        }
        vm.title = to.params.title
      })
      // ...
    },
    methods:{
      turn() {
        if (this.addressInfo.position === true) {
          this.sendData.position = 0
        }else {
          this.sendData.position = 1
        }
      },
      checkTage(item) {
        this.addressInfo.tage = item
      },
      checkGender(item) {
        this.addressInfo.gender = item
      },
      handleSaveAddr() {
        if (this.userId) {
           this.addressInfo.userId = this.userId
        } else if (this.$route.query.id) {
           this.addressInfo.userId = this.$route.query.id
        } else {
           this.addressInfo.userId = this.$route.params.userData.id
        }

        if (!this.addressInfo.name) {
          this.showMsg("请输入联系人")
          return
        }
        if (!this.addressInfo.phone) {
          this.showMsg("请输入电话号码")
          return
        }
        if (!this.addressInfo.address) {
          this.showMsg("请输入或选择地址")
          return
        }
        //存储数据
        if (this.title == "添加地址") {
          this.saveToAddress()
        } else {
          this.editToAddress()
        }
        
      },
      async saveToAddress() {
        let user_id = localStorage.getItem("user_login")
        this.sendData.user_id = parseInt(this.addressInfo.userId) 
        this.sendData.user_name = this.addressInfo.name
        this.sendData.user_gender = this.addressInfo.gender
        this.sendData.name = this.addressInfo.addrName
        this.sendData.mobile = this.addressInfo.phone
        this.sendData.address = this.addressInfo.address
        this.sendData.bottom = this.addressInfo.bottom
        this.sendData.lng = this.addressInfo.lng
        this.sendData.lat = this.addressInfo.lat
        this.sendData.tage = this.addressInfo.tage
        if (this.addressInfo.position == true) {
          this.sendData.position = 0
        } else {
          this.sendData.position = 1
        }
        const result = await reqAddUserAddress(user_id,this.sendData)
        if (result.code === 0) {
          if (!this.$store.getters.selectedAddress) {
            this.$store.dispatch("setSelectedAddress",this.sendData)
          }
          this.showMsg("添加成功")
          this.$router.push({
              name:"myAddress",
              query:{
                id:this.addressInfo.userId
              }
            })
        } else {
          this.showMsg("result.msg")
          return
        }
      },
      async editToAddress() {
        let user_id = localStorage.getItem("user_login")
        this.sendData.user_id = parseInt(this.addressInfo.userId) 
        this.sendData.user_name = this.addressInfo.name
        this.sendData.user_gender = this.addressInfo.gender
        this.sendData.name = this.addressInfo.addrName
        this.sendData.mobile = this.addressInfo.phone
        this.sendData.address = this.addressInfo.address
        this.sendData.bottom = this.addressInfo.bottom
        this.sendData.lng = this.addressInfo.lng
        this.sendData.lat = this.addressInfo.lat
        this.sendData.tage = this.addressInfo.tage
        this.sendData.id = this.addressInfo.id
        console.log(this.sendData);
        const result = await reaEditUserAddress(user_id,this.sendData)
        if (result.code === 0) {
          this.showMsg("修改成功")
          this.$router.push({
              name:"myAddress",
              query:{
                id:this.addressInfo.userId
              }
            })
          } else {
            this.showMsg("result.msg")
            return
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
    },
    components:{
        hader,
        formBlock,
        tabTage,
        addressSearch,
        reqAddUserAddress,
    },
}
</script>

<style scoped>
.addAddress {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  padding-top: 45px;
}
.viewbody {
  width: 100%;
  height: 100%;
  overflow: auto;
  box-sizing: border-box;
  background-color: #f5f5f5;
}

.content {
  padding-left: 4vw;
  background: #fff;
  font-size: 0.94rem;
}

.formblock {
  background-color: #fff;
  border-bottom: 1px solid #eee;
  display: flex;
}
.formblock .label-wrap {
  flex-basis: 17.333333vw;
  padding: 3.733333vw 0;
  line-height: 4.8vw;
  color: #333;
  font-weight: 700;
}

/* 确定按钮 */
.form-button-wrap {
  padding: 5.333333vw 4vw;
  display: flex;
}
.form-button-wrap .form-button {
  background: #00d762;
  text-align: center;
  border-radius: 0.533333vw;
  flex: 1;
  font-size: 1.1rem;
  line-height: 5.066667vw;
  color: #fff;
  padding: 3.333333vw 0;
  border: none;
  font-weight: 500;
}
</style>
