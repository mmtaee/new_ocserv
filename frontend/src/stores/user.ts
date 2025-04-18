import {defineStore} from 'pinia'

export const useUserStore = defineStore('user', {
    state: () => ({
        uid: "",
        username: "",
        isAdmin: false,
        lastLogin: "",
    }),

    actions: {
        setUser: async function (username: string, isAdmin: boolean, lastLogin?: string, uid?: string) {
            this.uid = uid || ""
            this.username = username
            this.isAdmin = isAdmin
            this.lastLogin = lastLogin || ""
        },
    },
})