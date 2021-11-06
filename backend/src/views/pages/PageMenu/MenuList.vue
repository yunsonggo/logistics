<template>
    <div style="width:100%">
        <el-card>
            菜单列表
        </el-card>
        <el-card style="margin-top: 20px">
            <el-table
                :data="tableData.slice((currentPage-1)*pageSize,currentPage*pageSize)"
                row-key="id"
                border
                :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
                style="width: 900px">
                <el-table-column prop="id" label="ID" width="85"></el-table-column>
                <el-table-column prop="menu_title" label="菜单名称" width="150"></el-table-column>
                <el-table-column prop="menu_path" label="菜单路径" width="150"></el-table-column>
                <el-table-column prop="icon" label="菜单图标" width="150" ></el-table-column>
                <el-table-column prop="manager_power" label="访问权限" width="150">
                    <template slot-scope="scope">
                        <span v-if="scope.row.manager_power == '7' ">管理员</span>
                        <span v-else-if="scope.row.manager_power == '77' ">系统管理员</span>
                        <span v-else>超级管理员</span>
                    </template>
                </el-table-column>
                <el-table-column label="操作">
                    <template slot-scope="scope">
                        <el-button
                            size="mini"
                            @click="handleEdit(scope.$index, scope.row)">编辑
                        </el-button>
                        <el-button
                            size="mini"
                            type="danger"
                            @click="handleDelete()">删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                background
                :current-page="currentPage"
                :page-sizes="[2, 5, 10, 15]"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="pageTotal">
            </el-pagination>
        </el-card>
    </div>
</template>

<script>
export default {
    name:'menulist',
    data() {
        return {
            tableData: [],
            currentPage: 1,
            pageSize: 5,
            pageTotal: 0,
            menuList:[]
        }
    },
    mounted() {
        this.getTableData(),
        this.onSubmit()
    },
    methods: {
        getTableData() {
            this.tableData = this.$store.getters.menuData
        },
        onSubmit() {
            this.pageTotal = this.tableData.length
        },
        handleEdit (index, row) {
            this.$router.push({name:'editmenu',params:{menuInfo:row}})
        },
        handleDelete () {
            this.$message({
                message: '警告! 后台菜单禁止删除!',
                type: 'warning'
            })
            return
        },
        handleSizeChange (size) {
            this.pageSize = size
        },
        handleCurrentChange (currentPage) {
            this.currentPage = currentPage
        },
    },
}
</script>