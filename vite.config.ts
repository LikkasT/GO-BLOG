import { defineConfig } from 'vite';

export default defineConfig({
  // 项目根目录
  root: '.',
  // 构建输出配置
  build: {
    outDir: 'dist', // 输出目录
    sourcemap: true, // 生成sourcemap
  },
  // 开发服务器配置
  server: {
    port: 3000, // 开发服务器端口
    open: true, // 自动打开浏览器
    // API代理配置，将/api请求转发到后端服务
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '') // 根据实际情况调整
      }
    }
  }
});