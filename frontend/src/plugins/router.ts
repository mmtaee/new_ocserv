import type {RouteRecordRaw} from 'vue-router'
import {createRouter, createWebHistory} from 'vue-router'
import HomeView from "../views/HomeView.vue";
import {useConfigStore} from "../stores/config.ts";

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'HomePage',
        component: HomeView,
        meta: {
            title: "Home",
        }
    },
    {
        path: '/setup',
        name: 'SetupPage',
        component: () => import('../views/SetupView.vue'),
        meta: {
            title: "Setup",
        }
    },
    // {
    //     path: '/login',
    //     name: 'LoginPage',
    //     component: () => import('../views/LoginView.vue'),
    // },
    // {
    //     path: '/error',
    //     name: 'ErrorPage',
    //     component: () => import('../views/ErrorView.vue'),
    // },
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: routes,
})


router.beforeEach((to, _from, next) => {
    const configStore = useConfigStore()

    if (to.meta?.title) {
        document.title = to.meta.title as string
    }

    if (!configStore.setup && to.path !== "/setup") {
        next("/setup")
        return;
    }
    if (configStore.setup && to.path === "/setup") {
        next("/")
        return
    }

    if (!configStore.setup && to.path !== '/setup') {
        next('/setup')
        return
    }

    if (!['/login', '/setup'].includes(to.path) && localStorage.getItem('token') === null) {
        next('/login')
        return
    }

    next()
})


export default router