import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';
import store from "./store.js"

const app = createApp(App)
app.use(ElementPlus)
app.use(store)
app.mount('#app')
