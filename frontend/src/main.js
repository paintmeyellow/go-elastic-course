import { createApp } from 'vue'
import Home from './Home.vue'
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';
import store from "./store.js"
import router from "./router.js"

const app = createApp(Home)
app.use(ElementPlus)
app.use(store)
app.use(router)
app.mount('#app')
