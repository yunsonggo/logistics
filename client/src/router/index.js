import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  linkActiveClass: 'active',
  routes: [
    {
      path: '/',
      //name: 'index',
      component: () => import('@/views/index.vue'),
      children: [
        {
          path:'',
          redirect:'/home'
        },
        {
          path:'/home',
          name:'home',
          component: () => import('@/views/home.vue'),
        },
        {
          path:'/order',
          name:'order',
          component: () => import('@/views/order.vue'),
        },
        {
          path:'/me',
          name:'me',
          component: () => import('@/views/me.vue'),
        },
        {
          path:'/address',
          name:'theAddress',
          component: () => import('@/views/address.vue'),
        },
        {
          path:'/city',
          name:'city',
          component: () => import('@/views/city.vue'),
        },
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/login.vue')
    },
    {
      path: '/search',
      name: 'search',
      component: () => import('@/views/search.vue')
    },
    {
      path:'/shop',
      name:'shop',
      redirect: '/goods',
      component:() => import('@/views/shops/shop.vue'),
      children:[
        {
          path: '/goods',
          name: 'goods',
          component: () => import('@/views/shops/goods')
        },
        {
          path: '/comments',
          name: 'comments',
          component: () => import('@/views/shops/comments')
        },
        {
          path: '/seller',
          name: 'seller',
          component: () => import('@/views/shops/seller')
        },
      ]
    },
    {
      path: '/myAddress/:id',
      name: 'myAddress',
      component: () => import('@/views/order/myAddress.vue')
    },
    {
      path: '/addAddress/:id',
      name: 'addAddress',
      component: () => import('@/views/order/addAddress.vue')
    },
    {
      path: '/settlement',
      name: 'settlement',
      component: () => import('@/views/order/settlement.vue')
    },
    {
      path: '/remark',
      name: 'remark',
      component: () => import('@/views/order/remark.vue')
    },
    {
      path: '/pay',
      name: 'pay',
      component: () => import('@/views/order/pay.vue')
    },
    {
      path: '/orderinfo',
      name: 'orderinfo',
      component: () => import('@/views/order/orderInfo.vue')
    },
  ]
});

// 路由守卫
router.beforeEach((to, from, next) => {
  const isLogin = localStorage.user_login ? true : false;
  if (to.path === '/myAddress' || to.path === '/addAddress' || to.path === '/settlement' || to.path==='/pay'||to.path==='/order'||to.path==='/orderinfo') {
    // 是否在登录状态下
    isLogin ? next() : next('/login');
  } else {
    next()
  }
});

export default router;
