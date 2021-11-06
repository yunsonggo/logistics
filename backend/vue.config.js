module.exports = {
    devServer: {
        open: true,
        //host: 'localhost',
        host: '0.0.0.0',
        port: 8081,
        https: false,
        hotOnly: false,
        proxy: {
          // 配置跨域
          '/backend': {
            target: 'http://192.168.1.102:8090/manager',
            ws: true,
            changOrigin: true,
            pathRewrite: {
              '^/backend': ''
            }
          },
          '/api': {
            target: 'http://192.168.1.102:8090/api',
            ws: true,
            changOrigin: true,
            pathRewrite: {
              '^/api': ''
            }
          },
          // '/tmsug': {
          //   target: 'https://suggest.taobao.com/sug',
          //   secure: false,
          //   ws: true,
          //   changOrigin: true,
          //   pathRewrite: {
          //     '^/tmsug': ''
          //   }
          // }
        },
        before: app => {}
      }
}