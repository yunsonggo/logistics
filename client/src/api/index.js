import ajax from './ajax'
const BASE_URL = 'http://192.168.1.102:8080/api'


//请求 获取验证码
export const reqSendCode = (phone,tpl_id,key) => ajax(BASE_URL + '/sms/send', {phone,tpl_id,key},'POST')
//请求 登录
export const reqLogin = (phone,verifyCode) => ajax(BASE_URL + '/sms/login',{phone,verifyCode},'POST')
//获取 用户信息
export const reqUserInfo = (user_id) => ajax(BASE_URL + '/auth/user/info',{user_id},'POST') 
//获取 用户地址信息
export const reqUserAddress = (user_id) => ajax(BASE_URL + '/auth/user/address/list',{user_id},'POST')
//发送 新增地址 保存入库
export const reqAddUserAddress = (user_id,addressInfo) => ajax(BASE_URL + '/auth/user/address/insertone',{user_id,addressInfo},'POST')
//发送 编辑地址 保存入库
export const reaEditUserAddress = (user_id,addressInfo) => ajax(BASE_URL + '/auth/user/address/editone',{user_id,addressInfo},'POST')
//发送 删除地址 删除数据
export const reqDelUserAddress = (user_id,addressInfo) => ajax(BASE_URL + '/auth/user/address/removeone',{user_id,addressInfo},'POST')



//获取 基础城市信息
export const reqCitiesData = () => ajax(BASE_URL + '/cities')
//获取 首页商品数据
export const reqProfileShoping = () => ajax(BASE_URL + '/profile/shopping')
//获取 首页类别数据
export const reqCategroy = () => ajax(BASE_URL + '/profile/category')
//获取 菜单
export const reqMenu = () => ajax(BASE_URL + '/tab/menu')
//获取 首页商家
export const reqProShop = (keys,page,count,condition) => ajax(BASE_URL + '/profile/supplier',{keys,page,count,condition},'POST')
//搜索 商家或商品
export const reqSearch = (keyword) => ajax(BASE_URL + '/profile/search',{keyword},'POST')
//获取 关键词相关词
export const reqWithKeyword = (area,code,k,q) => ajax('http://192.168.1.102:8080/tmsug',{area,code,k,q})
//获取 商家信息 通过ID
export const reqShopInfo = (id) => ajax(BASE_URL+'/shop/info',{id})
//获取 商家 参与服务
export const reqShopGive = (shop_id,give_ids) => ajax(BASE_URL + '/shop/give',{shop_id,give_ids},'POST')
//获取 商家 参与活动
export const reqShopSale = (shop_id,sale_ids) => ajax(BASE_URL + '/shop/sale',{shop_id,sale_ids},'POST')
//获取 商家 所有在售商品
export const reqShopAllGoods = (shop_id) => ajax(BASE_URL + '/shop/goods',{shop_id},'POST')
//获取 商家 商品分类
export const reqShopCates = (shop_id) => ajax(BASE_URL + '/shop/shopcate',{shop_id},'POST')
//获取 商品 评论
export const reqGoodEva = (id_name,id_val,page,count) => ajax(BASE_URL + '/good/evaluation',{id_name,id_val,page,count},'POST')
//获取 商家 根据分类ID
export const reqShopByCate = (cate_id) => ajax(BASE_URL + '/profile/shop/cate',{cate_id})
//提交 订单
export const reqOrderSend = (user_id,userid,order_list) => ajax(BASE_URL + '/auth/order/send',{user_id,userid,order_list},'POST')
//获取 用户所有订单
export const reqOrdersData = (user_id,userid) => ajax(BASE_URL + '/auth/order/all',{user_id,userid},'POST')