<template>
    <el-container id="app">
        <el-header>
            <el-menu
                    :default-active="activeIndex"
                    mode="horizontal"
                    text-color="#409EFF"
                    active-text-color="#67C23A"
                    class="top-nav">
                <el-menu-item class="logo-title">Cloud Messaging</el-menu-item>
                <el-menu-item class="nav-title" index="4"><a href="" target="_blank">GitHub</a>
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
            <div class="home-slogan">
                <el-row type="flex" justify="center">
                    <el-col :span="24">
                        <h1>CCM Is CCM Cloud Messaging!</h1>
                        <h2>Make Android Better By Pushing Message To Android Phone In China!</h2>
                    </el-col>
                </el-row>
            </div>
            <el-row type="flex" justify="center">
                <el-col :span="18">
                    <el-card shadow="hover" class="home-image-card">
                        <img class="home-image" src="./assets/img/android.jpg">
                    </el-card>
                </el-col>
            </el-row>
            <div class="block home-timeline">
                <el-timeline>
                    <el-timeline-item timestamp="2018/4/12" placement="top">
                        <el-card>
                            <h4>更新 Github 模板</h4>
                            <p>王小虎 提交于 2018/4/12 20:46</p>
                        </el-card>
                    </el-timeline-item>
                    <el-timeline-item timestamp="2018/4/3" placement="top">
                        <el-card>
                            <h4>更新 Github 模板</h4>
                            <p>王小虎 提交于 2018/4/3 20:46</p>
                        </el-card>
                    </el-timeline-item>
                    <el-timeline-item timestamp="2018/4/2" placement="top">
                        <el-card>
                            <h4>更新 Github 模板</h4>
                            <p>王小虎 提交于 2018/4/2 20:46</p>
                        </el-card>
                    </el-timeline-item>
                </el-timeline>
            </div>
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
        methods: {},
        components: {}
    }
</script>

<style class="scss">
    body {
        margin: 0 80px;
        text-align: center;
    }

    .top-nav {
        padding-top: 10px;
    }

    .logo-title {
        /*height: 80px !important;*/
        /*line-height: 80px !important;*/
        float: left !important;
        font-size: 30px !important;
        font-weight: bold;
    }

    .nav-title {
        height: 65px !important;
        /*line-height: 80px !important;*/
        float: right !important;
    }

    .home-image-card {
        border: none!important;
    }

    .home-image {
        width: 100%;
    }

    .home-slogan{
        margin: 30px 0 50px 0;
    }

    .home-slogan h1 {
        font-size: xx-large;
        color: #555;
        margin: 0px;
    }

    .home-slogan h2 {
        margin-top: 10px;
        font-size: large;
        color: #888;
    }

    .home-timeline {
        margin-top: 40px;
    }
</style>
