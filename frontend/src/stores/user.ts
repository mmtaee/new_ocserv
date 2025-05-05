import {defineStore} from 'pinia'
import {type ModelsUser, UserApi} from "@/api";
import router from "@/plugins/router.ts";

export const useUserStore = defineStore('user', {
    state: (): { user: ModelsUser | null } => ({
        user: null,
    }),

    actions: {
        setUser(user: ModelsUser) {
            this.user = user;
        },
        async fetchUser() {
            const api = new UserApi();
            const response = await api.userProfileGet();
            this.setUser(response.data);
        },
        async logout() {
            this.user = null
            localStorage.removeItem("token")
            await router.push("/login")
        }
    },

    getters: {
        getUser(state): ModelsUser | null {
            return state.user;
        },
    },
});