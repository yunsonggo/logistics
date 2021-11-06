import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

// types
const types = {
  SET_MENUDATA: "SET_MENUDATA",
  SET_MANAGER: "SET_MANAGER",
  SET_CATEGROY: "SET_CATEGROY",
  SET_GIVEDATA: "SET_GIVEDATA",
  SET_SALEPROMOTE: "SET_SALEPROMOTE",
}

// state
const state = {
  menuData: [],   // 后台菜单
  manager: {},    // 登录用户数据
  categroy:[],  // 大类 类别
  giveData:[], // 提供服务数据
  salePromote:[],// 促销活动
}

// getters
const getters = {
  menuData: state => state.menuData,
  manager: state => state.manager,
  categroy: state => state.categroy,
  giveData: state => state.giveData,
  salePromote: state => state.salePromote,
}

// mutations
const mutations = {
  [types.SET_MENUDATA](state,menuData) {
    if (menuData) {
      state.menuData = menuData
    } else {
      state.menuData = []
    }
  },
  [types.SET_MANAGER](state,manager) {
    if (manager) {
      state.manager = manager
    } else {
      state.manager = {}
    }
  },
  [types.SET_CATEGROY](state,categroy) {
    if (categroy) {
      state.categroy = categroy
    } else {
      state.categroy = []
    }
  },
  [types.SET_GIVEDATA](state,giveData) {
    if(giveData) {
      state.giveData = giveData
    } else {
      state.giveData = []
    }
  },
  [types.SET_SALEPROMOTE](state,salePromote) {
    if (salePromote) {
      state.salePromote = salePromote
    } else {
      state.salePromote = []
    }
  },
}

// actions
const actions = {
  setMenuData: ({commit},menuData) => {
    commit(types.SET_MENUDATA,menuData)
  },
  setManager: ({commit},manager) => {
    commit(types.SET_MANAGER,manager)
  },
  setCategroy:({commit},categroy) => {
    commit(types.SET_CATEGROY,categroy)
  },
  setGiveData:({commit},giveData) => {
    commit(types.SET_GIVEDATA,giveData)
  },
  setSalePromote:({commit},salePromote) => {
    commit(types.SET_SALEPROMOTE,salePromote)
  }
}

export default new Vuex.Store ({
  state,
  getters,
  mutations,
  actions
})
