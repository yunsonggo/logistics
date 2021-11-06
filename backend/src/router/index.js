import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URl,
  linkActiveClass: 'active',
  routes:[
    // 首页
    {
      path: '/yunsong/home',
      redirect: '/yunsong/home/index'
    },
    // 登录
    {
      path: '/yunsong/login',
      name: 'login',
      component: () => import('../views/Login.vue'),
    },
    // 注册
    {
      path: '/yunsong/signup',
      name: 'signup',
      component: () => import('../views/SignUp.vue'),
    },
    // 首页
    {
      path: '/yunsong/home',
      name: 'home',
      component: () => import('../views/Index.vue'),
      children:[
        // 首页
        {
          path:'/yunsong/home/index',
          name:'index',
          component: () => import('../views/pages/PageIndex.vue')
        },
        // 后台用户管理
        {
          path:'/yunsong/home/admin',
          name:'admin',
          component: () => import('../views/pages/Admin/Admin.vue'),
          children: [
            //管理员列表
            {
              path:'/yunsong/home/adminlist',
              name:'adminlist',
              component: () => import('../views/pages/Admin/AdminList.vue')
            },
            //添加管理员
            {
              path:'/yunsong/home/adminadd',
              name:'adminadd',
              component: () => import('../views/pages/Admin/AdminAdd.vue')
            },
            //添加管理员
            {
              path:'/yunsong/home/adminedit',
              name:'adminedit',
              component: () => import('../views/pages/Admin/AdminEdit.vue')
            },
          ]
        },
        // 后台菜单管理
        {
          path:'/yunsong/home/pagemenu',
          name:'pagemenu',
          component: () => import('../views/pages/PageMenu/PageMenu.vue'),
          children:[
            {
              path:'/yunsong/home/menulist',
              name:'menulist',
              component: () => import('../views/pages/PageMenu/MenuList.vue')
            },
            {
              path:'/yunsong/home/addmenu',
              name:'addmenu',
              component: () => import('../views/pages/PageMenu/AddMenu.vue')
            },
            {
              path:'/yunsong/home/editmenu',
              name:'editmenu',
              component: () => import('../views/pages/PageMenu/EditMenu.vue')
            },
          ]
        },
        // 商家管理
        {
          path:'/yunsong/home/supplier',
          name:'supplier',
          component: () => import('../views/pages/Supplier/Supplier.vue'),
          children:[
            {
              path:'/yunsong/home/supplist',
              name:'supplist',
              component: () => import('../views/pages/Supplier/SuppList.vue')
            },
            {
              path:'/yunsong/home/suppadd',
              name:'suppadd',
              component: () => import('../views/pages/Supplier/SuppAdd.vue')
            },
            {
              path:'/yunsong/home/suppedit',
              name:'suppedit',
              component: () => import('../views/pages/Supplier/SuppEdit.vue')
            },
          ]
        },
        // 分类管理
        {
          path:'/yunsong/home/cates',
          name:'cates',
          component: () => import('../views/pages/Cates/Cates.vue'),
          children:[
            {
              path:'/yunsong/home/cateslist',
              name:'cateslist',
              component: () => import('../views/pages/Cates/CatesList.vue')
            },
            {
              path:'/yunsong/home/catesedit',
              name:'catesedit',
              component: () => import('../views/pages/Cates/CatesEdit.vue')
            },
            {
              path:'/yunsong/home/catesadd',
              name:'catesadd',
              component: () => import('../views/pages/Cates/CatesAdd.vue')
            },
          ]
        },
        // 服务管理
        {
          path:'/yunsong/home/gives',
          name:'gives',
          component: () => import('../views/pages/Gives/Gives.vue'),
          children:[
            {
              path:'/yunsong/home/giveslist',
              name:'giveslist',
              component: () => import('../views/pages/Gives/GivesList.vue')
            },
            {
              path:'/yunsong/home/givesadd',
              name:'givesadd',
              component: () => import('../views/pages/Gives/GivesAdd.vue')
            },
            {
              path:'/yunsong/home/givesedit',
              name:'givesedit',
              component: () => import('../views/pages/Gives/GivesEdit.vue')
            },
          ]
        },
        // 活动管理
        {
          path:'/yunsong/home/salepromote',
          name:'salepromote',
          component: () => import('../views/pages/SalePromote/SalePromote.vue'),
          children:[
            {
              path:'/yunsong/home/salepromotelist',
              name:'salepromotelist',
              component: () => import('../views/pages/SalePromote/SalePromoteList.vue')
            },
            {
              path:'/yunsong/home/salepromoteadd',
              name:'salepromoteadd',
              component: () => import('../views/pages/SalePromote/SalePromoteAdd.vue')
            },
            {
              path:'/yunsong/home/salepromoteedit',
              name:'salepromoteedit',
              component: () => import('../views/pages/SalePromote/SalePromoteEdit.vue')
            },
          ]
        },

        // 客户管理
        {
          path:'/yunsong/home/user',
          name:'user',
          component: () => import('../views/pages/User/User.vue'),
          children:[
            {
              path:'/yunsong/home/userlist',
              name:'userlist',
              component: () => import('../views/pages/User/UserList.vue')
            },
            {
              path:'/yunsong/home/useredit',
              name:'useredit',
              component: () => import('../views/pages/User/UserEdit.vue')
            }
          ]
        },
        // 订单管理
        {
          path:'/yunsong/home/order',
          name:'order',
          component: () => import('../views/pages/Order/Order.vue'),
          children:[
            {
              path:'/yunsong/home/orderlist',
              name:'orderlist',
              component: () => import('../views/pages/Order/OrderList.vue')
            }
          ]
        },
      ]
    },
  ],
})

// 路由守卫
router.beforeEach((to, from, next) => {
    const isLogin = localStorage.manager_token ? true : false;
    if (to.path === '/yunsong/login' || to.path === '/yunsong/signup') {
        next()
    }else{
        // 是否在登录状态下
        isLogin ? next() : next('/yunsong/login');
    }
});

export default router
