import {createApp} from 'vue'
import App from './App.vue'
import {createPinia} from "pinia";
import router from "./plugins/router.ts";
import i18n from "./plugins/i18n.ts";
import vuetify from "./plugins/vuetify.ts";
import {useConfigStore} from "./stores/config.ts";
import {useUserStore} from "@/stores/user.ts";

const app = createApp(App)

app.use(createPinia())


;(async () => {
    const configStore = useConfigStore()
    await configStore.fetchConfig()

    if (localStorage.getItem("token")) {
        const userStore = useUserStore()
        await userStore.fetchUser()
    }

    app.use(vuetify)
    app.use(i18n)
    app.use(router)
    app.mount('#app')
})()
