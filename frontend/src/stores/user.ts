import {defineStore} from 'pinia'
import type {ModelsUser} from "@/api";

export const useUserStore = defineStore('user', {
    state: (): { user: ModelsUser | null } => ({
        user: null,
    }),

    actions: {
        setUser(user: ModelsUser) {
            this.user = user;
        },
    },
});