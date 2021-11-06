<template>
    <div v-if="sellerInfo">
        <div class="seller" >
            <section>
                <img :src="static_url + sellerInfo.banner" alt="">
                <h3>{{sellerInfo.name}}</h3>
                <p>{{sellerInfo.content}}</p>
                <div>查看品牌故事...</div>
            </section>
        </div>
    </div>
    <div v-else>
        暂未获得商家数据...
    </div>
</template>

<script>
import  {reqShopInfo} from '../../api/index'
import global from '../../components/global'

export default {
    name:"seller",
    data () {
        return {
            static_url:global.static_url,
            sellerInfo:null
        }
    },
    created() {
        this.getData()
    },
    methods:{
        async getData() {
            const result = await reqShopInfo (this.$route.query.id)
            if (result.code === 0) {
                this.sellerInfo = result.data
            } 
        }
    },
    components:{
        reqShopInfo,
        global
    }
}
</script>

<style scoped>
.seller section {
  margin-bottom: 2.666667vw;
  padding: 4.266667vw 4vw 4vw;
  font-size: 0.9rem;
  background-color: #fff;
  color: #666;
  border-bottom: 1px solid #eee;
}
section > img {
  width: 92vw;
  height: 52vw;
  margin-bottom: 4.266667vw;
}
section > h3 {
  color: #333;
  font-weight: 700;
  font-size: 1rem;
  margin-bottom: 1.066667vw;
}
section > p {
  height: 40px;
  line-height: 1.4;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}
section > div {
  margin: 4vw 0 -4vw;
  text-align: center;
  font-size: 0.346667rem;
  padding: 4vw 0;
}
</style>
