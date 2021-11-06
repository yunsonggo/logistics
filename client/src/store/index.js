import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

//types
const types = {
    SET_LOCATION: "SET_LOCATION",
    SET_ADDRESS: "SET_ADDRESS",
    SET_POSITION:"SET_POSITION",
    ORDER_INFO:"ORDER_INFO",
    SELECTED_ADDRESS:"SELECTED_ADDRESS",
    REMARK_INFO:"REMARK_INFO",   // 配套物料
    DISTANCES_ROUTERS:"DISTANCES_ROUTERS", // 路径规划数据
    DISTANCE_METRE:"DISTANCE_METRE", // 距离
    DELIVERY_PRICE:"DELIVERY_PRICE",
    DELIVERY_DATETIME:"DELIVERY_DATETIME" // 送达时间
    //GOODS_PRICE:"GOODS_PRICE"
}

// state
const state = {
    location: {},
    address: "",
    position: {},
    orderInfo:null,
    selectedAddress:null,
    remarkInfo:{
        ware:'',
        remark:'',
        invoice:''
    },
    distancesRouters:null,
    distanceMetre:"",
    deliveryPrice:Number,
    deliveryDateTime:""
    //goodsPrice:Number
}

//getters
const getters = {
    location: state => state.location,
    address: state => state.address,
    position: state => state.position,
    orderInfo: state => state.orderInfo,
    selectedAddress: state => state.selectedAddress,
    totalPrice : state => {
        let payPrice = 0
        if (!state.orderInfo) {
            if (localStorage.getItem("orderInfo")) {
                state.orderInfo = JSON.parse(localStorage.getItem("orderInfo"))
                localStorage.removeItem("orderInfo")
            }
        }
        const selectedGoods = state.orderInfo.selectedGoods
        selectedGoods.forEach(ele => {
            payPrice += (ele.price * ele.count)
        })
        if (state.orderInfo.shopInfo.is_union === 0) {
            if (state.distanceMetre) {
                payPrice += state.deliveryPrice
            }
        }
        return payPrice
    },
    remarkInfo: state => state.remarkInfo,
    distancesRouters: state => state.distancesRouters,
    distanceMetre: state => state.distanceMetre,
    deliveryPrice: state => state.deliveryPrice,
    goodsPrice: state => {
        let goodPrice = 0
        if (!state.orderInfo) {
            if (localStorage.getItem("orderInfo")) {
                state.orderInfo = JSON.parse(localStorage.getItem("orderInfo"))
                localStorage.removeItem("orderInfo")
            }
        }
        if (state.orderInfo) {
            const selectedGoods = state.orderInfo.selectedGoods
            selectedGoods.forEach(ele => {
                goodPrice += (ele.price * ele.count)
            });
        }
        return goodPrice
    },
    deliveryDateTime:state=>state.deliveryDateTime
}

// mutations
const mutations = {
    [types.SET_LOCATION](state,location) {
        if (location) {
            state.location = location
        } else {
            state.location = null
        }
    },
    [types.SET_ADDRESS](state,address) {
        if (address) {
            state.address = address
        } else {
            state.address = ""
        }
    },
    [types.SET_POSITION](state,position) {
        if (position) {
            state.position = position
        } else {
            state.position = null
        }
    },
    [types.ORDER_INFO](state,orderInfo) {
        if (orderInfo) {
            state.orderInfo = orderInfo
        } else {
            state.orderInfo = null
        }
    },
    [types.SELECTED_ADDRESS](state,selectedAddress) {
        if (selectedAddress) {
            state.selectedAddress = selectedAddress
        } else {
            state.selectedAddress = null
        }
    },
    [types.REMARK_INFO](state,remarkInfo) {
        if (remarkInfo) {
            state.remarkInfo = remarkInfo
        } else {
            state.remarkInfo = null
        }
    }, 
    [types.DISTANCES_ROUTERS](state,distancesRouters) {
        if (distancesRouters) {
            state.distancesRouters = distancesRouters
        } else {
            state.distancesRouters = null
        }
    }, 
    [types.DISTANCE_METRE](state,distanceMetre) {
        if (distanceMetre) {
            state.distanceMetre = distanceMetre
        } else {
            state.distanceMetre = ""
        }
    },
    [types.DELIVERY_PRICE](state,deliveryPrice) {
        if(deliveryPrice) {
            state.deliveryPrice = deliveryPrice
        } else {
            state.deliveryPrice = 0
        }
    },
    [types.DELIVERY_DATETIME](state,deliveryDateTime) {
        if (deliveryDateTime) {
            state.deliveryDateTime = deliveryDateTime
        } else {
            state.deliveryDateTime = ""
        }
    }
}

// actions
const actions = {
    setLocation: ({commit},location) => {
        commit(types.SET_LOCATION,location)
    },
    setAddress: ({commit},address) => {
        commit(types.SET_ADDRESS,address)
    },
    setPosition: ({commit},position) => {
        commit(types.SET_POSITION,position)
    },
    setOrderInfo: ({commit},orderInfo) => {
        commit(types.ORDER_INFO,orderInfo)
    },
    setSelectedAddress: ({commit},selectedAddress) => {
        commit(types.SELECTED_ADDRESS,selectedAddress)
    },
    setRemarkInfo: ({commit},remarkInfo) => {
        commit(types.REMARK_INFO,remarkInfo)
    },
    setDistancesRouters: ({commit},distancesRouters) => {
        commit(types.DISTANCES_ROUTERS,distancesRouters)
    },
    setDistanceMetre: ({commit},distanceMetre) => {
        commit(types.DISTANCE_METRE,distanceMetre)
    },
    setDeliveryPrice: ({commit},deliveryPrice) => {
        commit(types.DELIVERY_PRICE,deliveryPrice)
    },
    setDeliveryDateTime: ({commit},deliveryDateTime) => {
        commit(types.DELIVERY_DATETIME,deliveryDateTime)
    }
}
export default new Vuex.Store ({
    state,
    getters,
    mutations,
    actions
})
