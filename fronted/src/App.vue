<template>
    <el-container id="app">
        <!--header-->
        <el-row type="flex" justify="center">
            <el-col :lg="18">
                <el-header>
                    <el-menu v-if="!mobile"
                             :default-active="activeIndex"
                             mode="horizontal"
                             text-color="#409EFF"
                             active-text-color="#67C23A"
                             class="top-nav">
                        <el-menu-item class="logo-title">Cloud Messaging</el-menu-item>
                        <el-menu-item class="nav-title" index="4"><a href="https://github.com/chen622/Spider"
                                                                     target="_blank">GitHub</a>
                        </el-menu-item>
                        <el-menu-item class="nav-title" index="3" disabled>消息中心</el-menu-item>

                        <el-submenu class="nav-title" index="2">
                            <template slot="title">工作台</template>
                            <el-menu-item index="2-1">登陆</el-menu-item>
                            <el-menu-item index="2-2">注册</el-menu-item>
                            <el-menu-item index="2-3">试用</el-menu-item>
                        </el-submenu>
                        <el-menu-item class="nav-title" index="1">首页</el-menu-item>
                    </el-menu>
                </el-header>

                <!--main contain-->
                <el-main>
                    <div class="home-slogan">
                        <el-row type="flex" justify="center">
                            <el-col :span="24">
                                <h1>CCM 就是 CCM Cloud Messaging!</h1>
                                <h2>尝试让安卓的使用体验变得更好，即使在中国</h2>
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
                    <div class="home-slogan">
                        <el-row type="flex" justify="center">
                            <el-col :span="24">
                                <h1>项目时间线</h1>
                            </el-col>
                        </el-row>
                    </div>
                    <div class="block home-timeline">
                        <el-timeline>
                            <el-timeline-item timestamp="2018" placement="top">
                                <el-card>
                                    <h4>受 <a href="https://www.coolapk.com/apk/com.ruguoapp.jike">即刻 App</a>
                                        启发编写一个能够推送聚合消息的应用
                                    </h4>
                                    <p>晨茗 <a href="https://github.com/chen622">GitHub</a> 于 2018</p>
                                </el-card>
                            </el-timeline-item>
                            <el-timeline-item timestamp="2019.3.10" placement="top">
                                <el-card>
                                    <h4>使用 Golang 编写该项目以熟悉 Golang</h4>
                                    <p>晨茗 <a href="https://github.com/chen622">GitHub</a> 于 2019.3.10</p>
                                </el-card>
                            </el-timeline-item>
                            <el-timeline-item timestamp="2019/3/17" placement="top">
                                <el-card>
                                    <h4>项目启动并创建仓库</h4>
                                    <p>晨茗 <a href="https://github.com/chen622">GitHub</a> 提交于 2019/3/17 </p>
                                </el-card>
                            </el-timeline-item>
                        </el-timeline>
                    </div>
                </el-main>
            </el-col>
        </el-row>

        <!--footer-->
        <el-footer class="footer">
            <el-row :gutter="20">
                <el-col :lg="6">
                    <h3>链接</h3>
                    <p><a href="https://vuejs.org/" class="link">Vue.js</a></p>
                    <p><a href="http://element-cn.eleme.io" class="link">Element-UI</a></p>
                    <p><a href="https://golang.org/" class="link">Golang</a></p>
                    <p><a href="https://green-android.org/" class="link">安卓绿色应用规约</a></p>
                </el-col>
                <el-col :lg="6">
                    <h3>联系</h3>
                    <p><a href="mailto:ccm@ccm.ink" class="link">ccm@ccm.ink</a></p>
                </el-col>
                <el-col class="footer-title" :offset="6" :lg="6">
                    <h3>CCM Alpha</h3>
                    <p><a href="https://github.com/chen622"><img class="footer-icon" src="../public/github-icon.png"></a></p>
                    <p>©2019 Chenming C.<br> All rights reserved.</p>
                </el-col>
            </el-row>
        </el-footer>

    </el-container>
</template>

<script>
    export default {
        name: 'app',
        data () {
            return {
                mobile: false,
                activeIndex: '1',
                isCollapse: true
            }
        },
        methods: {},
        components: {}
    }
</script>

<style class="scss">
    body {
        margin: 0px;
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
        border: none !important;
    }

    .home-image {
        width: 100%;
    }

    .home-slogan {
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

    .el-main {
        overflow: inherit !important;
    }

    .footer {
        margin-top: 40px;
        height: auto !important;
        background: #f7fbfd;
        padding: 0px 180px 60px 180px !important;
    }

    .footer .el-col {
        text-align: left;
    }

    .footer .h3 {
        color: #333;
    }

    .footer-title {
        text-align: right !important;
    }
    .footer-title p {
        color: #888
    }

    .footer-icon {
        height: 3em;
        width: 3em;
    }

    .link {
        color: #888;
        text-decoration: none
    }

    a:link {
        text-decoration: none;
    }

    a:active {
    }

    a:visited {
        text-decoration: none;
    }

    a:hover {
        color: #67C23A;
        text-decoration: underline;
    }
</style>
