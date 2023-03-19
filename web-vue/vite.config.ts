import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    /* 是否开启 $ref , vue3.2 的语法糖 */
    vue({
      refSugar: true,
        reactivityTransform: true,
    }),
],



server: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:3000',
        changeOrigin: true,
        rewrite: path => path.replace(/^\/api/, '')
      }
    }
  },



base: './', //打包相对路径
})


