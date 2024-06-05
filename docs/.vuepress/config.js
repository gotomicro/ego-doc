const moment = require("moment");
const path = require('path')
module.exports = {
    title: "EGO",
    description: "最简单的GO微服务框架",
    head: [
        ["link", {rel: "icon", href: "/logo.png"}],
        [
            "meta",
            {
                name: "keywords",
                content: "微服务,框架,ego,微服务框架,Go微服务框架,golang,micro service,gRPC",
            },
        ],
        ["script",{},`
var _hmt = _hmt || [];
(function() {
  var hm = document.createElement("script");
  hm.src = "https://hm.baidu.com/hm.js?5b74978e3772cd939e423b6c55896b6d";
  var s = document.getElementsByTagName("script")[0]; 
  s.parentNode.insertBefore(hm, s);
})();
        `]
    ],
    configureWebpack: () => {
        const NODE_ENV = process.env.NODE_ENV
        //判断是否是生产环境
        if(NODE_ENV === 'production'){
            return {
                output: {
                    publicPath: 'https://ego-org.com/ego-org/'
                },
                resolve: {
                    //配置路径别名
                    alias: {
                        'public': path.resolve(__dirname, './public')
                    }
                }
            }
        }else{
            return {
                resolve: {
                    //配置路径别名
                    alias: {
                        'public': path.resolve(__dirname, './public')
                    }
                }
            }
        }
    },
    markdown: {
        lineNumbers: true, // 代码块显示行号
    },
    themeConfig: {
        nav: [
            {
                text: "首页",
                link: "/",
            },
            {
                text: "框架",
                link: "/frame/",
            },
            {
                text: "最佳实践",
                link: "/practice/",
            },
            {
                text: "微服务",
                link: "/micro/",
            },
            {
                text: "手册",
                link: "/handbook/",
            },
            // {
            //     text: "Awesome",
            //     link: "/awesome/",
            // },
            {
                text: "EGO",
                link: "https://github.com/gotomicro/ego",
            },
        ],
        docsDir: "docs",
        docsBranch: "main",
        editLinks: true,
        editLinkText: "在github.com上编辑此页",
        sidebar: {
            "/summary/": [""], //这样自动生成对应文章
            "/frame/": [
                "quickstart/quickstart",
                "releasenote",
                {
                    title: "核心模块",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "core/build",
                        "core/config",
                        "core/logger",
                        "core/trace",
                        {
                            title: "服务注册与发现",
                            collapsable: false, // 可选的, 默认值是 true,
                            children: [
                                "core/ETCD服务注册与发现使用",
                                "core/ETCD服务注册与发现原理",
                            ],
                        },
                    ],
                },
                {
                    title: "服务模块",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "server/http",
                        "server/grpc",
                        "server/governor",
                    ],
                },
                {
                    title: "任务模块",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "task/job",
                        "task/cron",
                    ],
                },
                {
                    title: "客户端模块",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "client/grpc",
                        "client/http",
                        "client/gorm",
                        "client/redis",
                        "client/mongo",
                        "client/kafka",
                        "client/sentinel",
                        "client/eetcd",
                        "client/ek8s",

                    ],
                },
                {
                    title: "工具",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "tool/tool",
                    ],
                },
                {
                    title: "网关模块",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "gateway/gateway",
                    ],
                },
                {
                    title: "治理模块",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "governor/metric",
                    ],
                },
                {
                    title: "最佳实践",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "bestpractice/metric",
                        "bestpractice/logger",
                    ],
                },
            ], //这样自动生成对应文章
            "/practice/": [
                {
                    title: "第一章 调试和错误处理",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "chapter1/debug",
                        "chapter1/egotrace",
                        "chapter1/error",
                    ],
                },
                {
                    title: "第二章 稳定性",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "sla/overview",
                        "sla/monitor",
                    ],
                },
            ],
            "/micro/": [
                "大纲",
                {
                    title: "第一章 编译和部署",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "chapter1/build",
                        "chapter1/deploy",
                    ],
                },
                {
                    title: "第二章 基础组件",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "chapter2/flag",
                        "chapter2/config",
                        "chapter2/logger",
                        "chapter2/trace",
                    ],
                },
                {
                    title: "第三章 gRPC",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "chapter3/注册中心",
                        "chapter3/服务注册",
                        "chapter3/服务发现",
                    ],
                },
                {
                    title: "第四章 测试",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "chapter4/unittest",
                    ],
                },
                {
                    title: "第八章 治理",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "chapter10/sla",
                    ],
                },
            ],
            "/handbook/": [
                {
                    title: "K8S",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "k8s/docker_hub",
                        "k8s/docker_images",
                        "k8s/k8s_intro",
                        "k8s/k8s_deployment",
                        "k8s/k8s_pod",
                        "k8s/kubectl_install",
                        "k8s/kubectl_cmd",
                    ],
                },
                {
                    title: "WEB",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "web/前端优化",
                    ],
                },
            ],
            // "/awesome/": [
            //     "errors",
            //     "gracefulstop",
            //     "map锁double check",
            //     "egov1.0.3"
            // ]
        },
        sidebarDepth: 2,
        lastUpdated: "上次更新",
        serviceWorker: {
            updatePopup: {
                message: "发现新内容可用",
                buttonText: "刷新",
            },
        },
    },
    plugins: [
        [
            "@vuepress/last-updated",
            {
                transformer: (timestamp, lang) => {
                    // 不要忘了安装 moment
                    const moment = require("moment");
                    moment.locale("zh-cn");
                    return moment(timestamp).format("YYYY-MM-DD HH:mm:ss");
                },

                dateOptions: {
                    hours12: true,
                },
            },
        ],
        [
            '@vssue/vuepress-plugin-vssue',
            {
                platform: 'github', //v3的platform是github，v4的是github-v4
                locale: 'zh', //语言
                // 其他的 Vssue 配置
                owner: 'gotomicro', //github 账户名或组织名
                repo: 'ego-doc', //github 一个项目的名称
                clientId: '601dc4dbe9ae8e87d76f',//注册的 Client ID
                clientSecret: 'de308ea181268f753305b1f7c91b6d7712be694a',//注册的 Client Secret
                autoCreateIssue: true // 自动创建评论，默认是false，最好开启，这样首次进入页面的时候就不用去点击创建评论的按钮了。
            },
        ],
        "@vuepress/back-to-top",
        "@vuepress/active-header-links",
        "@vuepress/medium-zoom",
        "@vuepress/nprogress",
    ],
};
