<template>
    <div style="width: 100%">
        <el-card>添加菜单</el-card>
        <el-card style="margin-top:20px; width: 600px">
            <el-form :model="addMenuForm" :rules="rules" ref="addMenuForm" label-width="100px">
                <el-form-item label="菜单名称" prop="menu_title">
                    <el-input v-model="addMenuForm.menu_title"></el-input>
                </el-form-item>
                <el-form-item label="上级菜单" prop="parent_id">
                    <el-select v-model="addMenuForm.parent_id" placeholder="一级菜单">
                        <el-option label="一级菜单" value= "0" ></el-option>
                        <el-option v-for="(item,index) in menuData" :key="index" :label="item.menu_title" :value="item.id"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="菜单路径" prop="menu_path">
                    <el-input v-model="addMenuForm.menu_path"></el-input>
                </el-form-item>
                <el-form-item label="菜单图标" prop="icon">
                    <el-input v-model="addMenuForm.icon"></el-input>
                </el-form-item>
                 <el-form-item label="菜单权限" prop="manager_power">
                    <el-radio-group v-model="addMenuForm.manager_power">
                        <el-radio label="7">管理员</el-radio>
                        <el-radio label="77">系统管理员</el-radio>
                        <el-radio label="777">超级管理员</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="是否启用" prop="is_show">
                    <el-switch v-model="addMenuForm.is_show"></el-switch>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="submitForm('addMenuForm')">立即创建</el-button>
                    <el-button @click="resetForm('ruleForm')">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
    </div>
</template>

<script>
import {reqMenuAdd} from '../../../api/index'

export default {
    name:'addmenu',
    data() {
        return {
            addMenuForm: {
                menu_title :'',
                parent_id : '',
                is_show : true,
                manager_power:'',
                menu_path: '',
                icon: ''
            },
            rules: {
                menu_title: {required: true, message: '请输入菜单名称', trigger: 'blur'},
                menu_path: {required: true, message: '请输入菜单路径', trigger: 'blur'},
            }
        }
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
                    if (this.addMenuForm.is_show) {
                        this.addMenuForm.is_show = 2
                    } else {
                        this.addMenuForm.is_show = 1
                    }
                    if (this.addMenuForm.parent_id == "" || this.addMenuForm.parent_id == "0") {
                        this.addMenuForm.parent_id = 0
                    }
                    if (this.addMenuForm.manager_power == "") {
                        this.addMenuForm.manager_power = "7"
                    }
                    // 发送添加数据
                    this.AddMenu()
                    console.log(this.addMenuForm);
                } else {
                    return false;
                }
            });
        },
        resetForm(formName) {
            this.$refs[formName].resetFields();
        },
        // 发送添加数据
        async AddMenu() {
            let manager_token = localStorage.getItem("manager_token") 
            const result = await reqMenuAdd(manager_token,this.addMenuForm)
            if (result.code === 0) {
                let menuData = this.$store.getters.menuData
                menuData.push(result.data)
                this.$store.dispatch('setMenuData',menuData)
                this.$message({
                    message: '添加菜单成功!',
                    type: 'success'
                });
                this.$router.push('/yunsong/home')
                return
            } else {
                this.$message.error(result.msg)
                return
            }
        }
    }
}
</script>
