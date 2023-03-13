import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from './App.vue'

if (import.meta.hot) {
    import.meta.hot.accept();
  }



const app = createApp(App)

app.use(ElementPlus)
app.mount('#app')


