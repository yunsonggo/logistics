<template>
    <div style="width:100%">
        <el-card>
            <span>编辑菜单</span>
        </el-card>
        <el-card v-if="showForm" style="margin-top:20px; width: 600px">
            <el-form :model="menuInfo" :rules="rules" ref="menuInfo" label-width="100px">
                <el-form-item label="菜单名称" prop="menu_title">
                    <el-input v-model="menuInfo.menu_title"></el-input>
                </el-form-item>
                <el-form-item label="上级菜单" prop="parent_id">
                    <el-select v-model="menuInfo.parent_id" placeholder="一级菜单">
                        <el-option label="一级菜单" value= "0" ></el-option>
                        <el-option v-for="(item,index) in menuData" :key="index" :label="item.menu_title" :value="item.id"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="菜单路径" prop="menu_path">
                    <el-input v-model="menuInfo.menu_path"></el-input>
                </el-form-item>
                <el-form-item label="菜单图标" prop="icon">
                    <el-input v-model="menuInfo.icon"></el-input>
                </el-form-item>
                 <el-form-item label="菜单权限" prop="manager_power">
                    <el-radio-group v-model="menuInfo.manager_power">
                        <el-radio label="7">管理员</el-radio>
                        <el-radio label="77">系统管理员</el-radio>
                        <el-radio label="777">超级管理员</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="是否启用" prop="is_show">
                    <el-switch v-model="menuInfo.is_show"></el-switch>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="submitForm('menuInfo')">立即修改</el-button>
                    <el-button @click="resetForm()">取消</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card v-if="!showForm" style="margin-top:20px; width: 600px">
            请从 "菜单列表" 页选择菜单编辑
        </el-card>
    </div>
</template>

<script>
import {reqMenuEdit} from '../../../api/index'

export default {
    name:'editmenu',
    data() {
        return {
            menuInfo:{},
            rules: {
                menu_title: {required: true, message: '请输入菜单名称', trigger: 'blur'},
                menu_path: {required: true, message: '请输入菜单路径', trigger: 'blur'},
            },
            showForm:false
        }
    },
    beforeRouteEnter(to,from,next) {
        next(vm => {
            if (to.params.menuInfo) {
                vm.menuInfo = to.params.menuInfo
                console.log(vm.menuInfo.parent_id);
                if (vm.menuInfo.is_show == 1) {
                    vm.menuInfo.is_show = false
                } else {
                    vm.menuInfo.is_show = true
                }
                if (vm.menuInfo.parent_id == 0) {
                    vm.menuInfo.parent_id = ""
                }
                vm.showForm = true
            } else {
                vm.showForm = false
            }
        })
    },
    computed: {
        menuData() {
            return this.$store.getters.menuData
        }
    },
    methods:{
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    if (this.menuInfo.is_show) {
                        this.menuInfo.is_show = 2
                    } else {
                        this.menuInfo.is_show = 1
                    }
                    if (this.menuInfo.parent_id == "" || this.menuInfo.parent_id == "0") {
                        this.menuInfo.parent_id = 0
                    }
                    if (this.menuInfo.manager_power == "") {
                        this.menuInfo.manager_power = "7"
                    }
                    this.$delete(this.menuInfo,'children')
                    // 发送修改数据
                    this.editMenu()
                } else {
                    return false;
                }
            });
        },
        resetForm() {
            this.$router.push('/yunsong/home/menulist')
            return
        },
        // 发送编辑菜单数据
        async editMenu() {
            let manager_token = localStorage.getItem("manager_token")
            const result = await reqMenuEdit(manager_token,this.menuInfo)
            if (result.code === 0) {
                this.$message({
                    message: '菜单修改成功',
                    type: 'success'
                })
                this.$router.push('/yunsong/home/menulist')
                return
            } else {
                this.$message.error(result.msg)
                return
            }
        }
    }
}
</script>