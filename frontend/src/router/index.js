import { createRouter, createWebHistory } from "vue-router";

const routes = [
    {
        path: '/',
        name: 'Login',
        component: () => import('../views/Login.vue')
    }, {
        path: '/register',
        name: 'Register',
        component: () => import('../views/Register.vue')
    }, {
        path: '/check',
        name: 'Check',
        component: () => import('../views/Check.vue')
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router;