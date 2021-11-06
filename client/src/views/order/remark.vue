<template>
    <div class="remark">
        <hader :isLeft="true" title="订单备注"></hader>
        <div class="view-body">
            <section>
                <textarea placeholder="填写备注信息" v-model="text" maxlength="150"></textarea>
                <div class="switch">
                    <span v-for="(item,index) in redioItem" :key="index" 
                          class="switch-item item-line" :class="{'selected':item.select}"
                          @click="handleRadioItem(item)"
                          >
                        {{item.name}}
                    </span>
                </div>
                <div class="switch" v-for="(item,index) in switchItem" :key="index">
                    <span @click="item.select = !item.select" :class="{'selected':item.select}" class="switch-item">{{item.name}}</span>
                </div>
            </section>
            <button @click="submitClick" class="btn-submit">确定</button>
        </div>
    </div>
</template>

<script>
import hader from '../../components/header'

export default {
    name:"remark",
    data () {
        return {
            redioItem:[
                {select:false,name:"现在配送"},
                {select:false,name:"预约配送"},
                {select:false,name:"随时配送"},
            ],
            switchItem:[
                {select:false,name:"调货换货"},
                {select:false,name:"轻拿轻放"},
                {select:false,name:"捆绑打包"},
                {select:false,name:"不好停车"},
                {select:false,name:"提前联系"},
            ],
            text:""
        }
    },
    methods:{
        handleRadioItem(item) {
            this.redioItem.forEach(ele => {
                ele.select = false
            })
            item.select = true
        },
        submitClick() {
            let selectItems = ""
            this.redioItem.forEach(ele => {
                if (ele.select) {
                    selectItems += ele.name + ","
                }
            })
            this.switchItem.forEach(ele => {
                if(ele.select) {
                    selectItems += ele.name + ","
                }
            })
            selectItems += this.text
            // 存储
            localStorage.removeItem("remarkInfo")
            this.$store.dispatch("setRemarkInfo",{
                ware:this.$store.getters.remarkInfo.ware,
                remark:selectItems,
                invoice:this.$store.getters.remarkInfo.invoice
            })
            localStorage.setItem("remarkInfo",JSON.stringify({
                ware:this.$store.getters.remarkInfo.ware,
                remark:selectItems,
                invoice:this.$store.getters.remarkInfo.invoice
            }))
            this.$router.go(-1)
        }
    },
    components:{
        hader
    }
}
</script>


<style scoped>
.remark {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  padding-top: 45px;
}

.view-body {
  width: 100%;
  height: 100%;
  overflow: auto;
  box-sizing: border-box;
  font-size: 0.8rem;
  color: #333;
}
.view-body section {
  margin-top: 2.133333vw;
  padding: 4.266667vw;
  background-color: #fff;
  margin-bottom: 2.133333vw;
}
.view-body section textarea {
  width: 100%;
  height: 29.866667vw;
  margin-bottom: 4.266667vw;
  padding: 4.266667vw;
  border: 1px solid #ccc;
  border-radius: 0.666667vw;
  background-color: #f9f9f9;
  color: #666;
  resize: none;
  box-sizing: border-box;
  outline: none;
}
.switch {
  display: inline-block;
  margin: 0 2.666667vw 2.666667vw 0;
  border: 1px solid #ddd;
  overflow: hidden;
  border-radius: 1.066667vw;
}
.switch-item {
  display: inline-block;
  padding: 0 2.133333vw;
  height: 8vw;
  line-height: 8vw;
  text-align: center;
  color: #666;
}
.item-line::after {
  content: " ";
  display: inline-block;
  height: 4vw;
  width: 1px;
  background: #ddd;
  line-height: 8vw;
  vertical-align: middle;
  left: 2.266667vw;
  position: relative;
}
.btn-submit {
  display: block;
  padding: 3.466667vw 0;
  color: #fff;
  background-color: #00e067;
  border-radius: 0.533333vw;
  opacity: 0.98;
  width: 92vw;
  margin: 3.133333vw auto 0;
  font-size: 1rem;
}

/* 选中样式 */
.switch-item.selected {
  background: #0186ff;
  color: #fff;
  position: relative;
}
</style>