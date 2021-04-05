import { createWebHistory, createRouter } from "vue-router";
import Home from "@/Home";

const routes = [
    { path: '/', component: Home },
];

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
});

export default router;