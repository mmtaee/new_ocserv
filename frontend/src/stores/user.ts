import {defineStore} from 'pinia'
import {type ModelsUser, UserApi} from "@/api";

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
    },

    getters: {
        getUser(state): ModelsUser | null {
            return state.user;
        },
    },
});