const moment = require("moment");
module.exports = {
    title: "EGO",
    description: "最简单的GO框架",
    head: [
        ["link", { rel: "icon", href: "/logo.png" }],
        [
            "meta",
            {
                name: "keywords",
                content: "Go,golang,ego,micro service,gRPC",
            },
        ],
    ],

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
                {
                    title: "核心模块",
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        "core/build",
                        "core/config",
                    ],
                },
            ], //这样自动生成对应文章
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
        "@vuepress/back-to-top",
        "@vuepress/active-header-links",
        "@vuepress/medium-zoom",
        "@vuepress/nprogress",
    ],
};