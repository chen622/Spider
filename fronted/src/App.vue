<template>
    <el-container id="app">
        <el-header>
            <el-menu
                    :default-active="activeIndex"
                    mode="horizontal"
                    text-color="#409EFF"
                    active-text-color="#67C23A"
                    @select="handleSelect" class="top-nav">
                <el-menu-item class="logo-title">CCM</el-menu-item>
                <el-menu-item class="nav-title" index="4"><a href="https://www.ele.me" target="_blank">订单管理</a>
                </el-menu-item>
                <el-menu-item class="nav-title" index="3" disabled>消息中心</el-menu-item>

                <el-submenu class="nav-title" index="2">
                    <template slot="title">我的工作台</template>
                    <el-menu-item index="2-1">选项1</el-menu-item>
                    <el-menu-item index="2-2">选项2</el-menu-item>
                    <el-menu-item index="2-3">选项3</el-menu-item>
                    <el-submenu index="2-4">
                        <template slot="title">选项4</template>
                        <el-menu-item index="2-4-1">选项1</el-menu-item>
                        <el-menu-item index="2-4-2">选项2</el-menu-item>
                        <el-menu-item index="2-4-3">选项3</el-menu-item>
                    </el-submenu>
                </el-submenu>
                <el-menu-item class="nav-title" index="1">首页</el-menu-item>

            </el-menu>
        </el-header>

        <el-main>

        </el-main>

        <el-footer>

        </el-footer>
    </el-container>
</template>

<script>

    export default {
        name: 'app',
        data () {
            return {
                activeIndex: '1',
            }
        },
        methods: {
            handleSelect (key, keyPath) {
                console.log(key, keyPath);
            }
        },
        components: {}
    }
</script>

<style class="scss">
    body {
        margin: 0 60px;
    }

    .top-nav{
    }

    .logo-title {
        height: 80px!important;
        line-height: 80px!important;
        float: left !important;
        font-size: 30px!important;
    }

    .nav-title {
        height: 80px!important;
        line-height: 80px!important;
        float: right !important;
    }


</style>
