# Vue 3 + TypeScript + Vite + ElementPlus


# 使用说明
### cd web-vue

### npm install
 
### npm  run  dev





# 开发笔记

# 一、创建项目
npm create vite@latest   按提示创建项目

# 二、引入elementPlus
 npm install element-plus --save

 // main.ts
import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'

const app = createApp(App)

app.use(ElementPlus)
app.mount('#app')

# 解决热更新问题


# 引入 axios,封装API


# 联调测试

