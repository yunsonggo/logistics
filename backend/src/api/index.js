import ajax from './ajax'
const BASE_URL = 'http://192.168.1.102:8081/backend'
const WEB_URL = 'http://192.168.1.102:8081/api'

// 请求 验证 验证码
export const reqManagerVerifyCaptcha = (captcha_token) => ajax(BASE_URL + '/captcha/verify',{captcha_token},'POST') 
// 登录 请求验证
export const reaManagerLogin = (user_name,user_password) => ajax(BASE_URL + '/login',{user_name,user_password},'POST')
// 注册 新用户
export const reqManagerSignUp = (user_name,user_password) => ajax(BASE_URL + '/sign/up',{user_name,user_password},'POST')
// TOKEN+ID 请求用户数据
export const reqManagerInfo = (manager_token,manager_id) => ajax(BASE_URL + '/auth/manager/info',{manager_token,manager_id},'POST')
// 请求菜单数据
export const reqMenuList = (manager_token) => ajax(BASE_URL + '/auth/manager/menu/list',{manager_token},'POST')
// 添加菜单
export const reqMenuAdd = (manager_token,add_menudata) => ajax(BASE_URL + '/auth/manager/menu/add',{manager_token,add_menudata},'POST')
// 修改菜单
export const reqMenuEdit = (manager_token,add_menudata) => ajax(BASE_URL + '/auth/manager/menu/edit',{manager_token,add_menudata},'POST') 
// 工作人员列表
export const reqAdminList = (manager_token) => ajax(BASE_URL + '/auth/manager/admin/list',{manager_token},'POST')
// 后台添加工作人员
export const reqAddAdmin = (manager_token,manager_data) => ajax(BASE_URL + '/auth/manager/admin/add',{manager_token,manager_data},'POST')
// 编辑工作人员基本信息
export const reqEditAdminInfo = (manager_token,manager_data) => ajax(BASE_URL + '/auth/manager/admin/edit',{manager_token,manager_data},'POST')
// 删除非超级管理员账号
export const reqDelAdminById = (manager_token,manager_data) => ajax(BASE_URL + '/auth/manager/admin/del',{manager_token,manager_data},'POST')
// 获取商家列表
export const reqSuppList = (manager_token) => ajax(BASE_URL + '/auth/manager/supp/list',{manager_token},'POST')
// 提交 商家编辑数据
export const reqSuppEdit = (manager_token,supplier_data) => ajax(BASE_URL + '/auth/manager/supp/edit',{manager_token,supplier_data},'POST')
// 提交 添加商家数据
export const reqSuppAdd = (manager_token,supplier_data) => ajax(BASE_URL + '/auth/manager/supp/add',{manager_token,supplier_data},'POST')
// 提交 修改商家类别
export const reqCategoryEdit = (manager_token,category_data) => ajax(BASE_URL + '/auth/manager/category/edit',{manager_token,category_data},'POST')
// 提交 添加商家类别
export const reqCategoryAdd = (manager_token,category_data) => ajax(BASE_URL + '/auth/manager/category/add',{manager_token,category_data},'POST')
// 获取所有服务数据
export const reqGiveList = (manager_token) => ajax(BASE_URL + '/auth/manager/give/list',{manager_token},'POST')
// 提交更新服务数据
export const reqGiveEdit = (manager_token,give_data) => ajax(BASE_URL + '/auth/manager/give/edit',{manager_token,give_data},'POST')
// 提交添加服务数据
export const reqGiveAdd = (manager_token,give_data) => ajax(BASE_URL + '/auth/manager/give/add',{manager_token,give_data},'POST')
// 获取所有活动数据
export const reqSalePromote = (manager_token) => ajax(BASE_URL + '/auth/manager/salepromote/list',{manager_token},'POST')
// 提交修改活动数据
export const reqSalePromoteEdit = (manager_token,promote_data) => ajax(BASE_URL + '/auth/manager/salepromote/edit',{manager_token,promote_data},'POST')
// 提交 添加活动数据
export const reqSalePromoteAdd = (manager_token,promote_data) => ajax(BASE_URL + '/auth/manager/salepromote/add',{manager_token,promote_data},'POST')

// 获取所有用户数据
export const reqUserList = (manager_token) => ajax(BASE_URL + '/auth/manager/user/list',{manager_token},'POST')
// 调用前台接口 所有商家类别数据 
export const reqCategroy = () => ajax(WEB_URL + '/profile/category')

// 获取所有订单数据
export const reqOrderList = (manager_token) => ajax(BASE_URL + '/auth/manager/order/list',{manager_token},'POST')

// 根据订单获取所有 已售商品数据
export const reqSaleGoodsByOrderId = (manager_token,order_info) => ajax(BASE_URL + '/auth/manager/order/salegood',{manager_token,order_info},'POST')