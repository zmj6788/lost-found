import { fileURLToPath, URL } from 'node:url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'


export default ({ mode }) => {
  const env = loadEnv(mode, process.cwd())
  const baseUrl = env.VITE_API
  return defineConfig({
    envPrefix: ["VITE_"],  // 需要使用的前缀
    plugins: [vue()],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      },
    },
    server: {
      // host: "127.0.0.1",  // 项目运行的ip
      // port: 80,  // 项目运行的端口
      // proxy: {
      //   "/uploads": {
      //     target: baseUrl
      //   },
      //   "/api": {
      //     target: baseUrl,
      //     changeOrigin: true,
      //   }
      // },
      host: '0.0.0.0', // 监听所有IP地址，包括localhost和自定义域名
      port: 80, // 选择一个非标准端口以避免权限问题
      strictPort: true, // 确保端口被占用时不自动切换
      hmr: { // 热模块重载相关配置，确保跨域安全
        clientPort: 80, // 与server.port保持一致
      },
      // 注意：Vue CLI或Vite的proxy配置主要是针对API请求转发，对于直接通过域名访问应用不适用
      proxy: {
        // 保留原有代理配置，用于开发环境API请求转发
        "/uploads": {
          target: baseUrl
        },
        "/api": {
          target: baseUrl,
          changeOrigin: true,
        },
        "/web/api": {
          //将/web/api替换成/api，实现代理
          target: baseUrl,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/web\/api/, '/api')
        },
        '/web/show/api': {
          target: baseUrl, // 目标地址，这里用'/'表示将'/web/show/api'请求映射到本地'/api'
          rewrite: (path) => path.replace(/^\/web\/show\/api/, '/api'), // 重写路径，移除'/web/show/api'前缀
          changeOrigin: true, // 如果你的目标服务器验证host字段，需要设置为true
          ws: true, // 如果要代理 websockets，记得设置这个选项
        }
        
    }
  }
  })
}