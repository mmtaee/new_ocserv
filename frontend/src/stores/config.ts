import {defineStore} from 'pinia'
import {PanelApi} from "../api";

export const useConfigStore = defineStore('config', {
    state: () => ({
        setup: false,
        googleCaptchaSiteKey: "",
    }),

    actions: {
        fetchConfig: async function () {
            const api = new PanelApi()
            const res = await api.panelGet()
            this.setup = res.data.setup
            this.googleCaptchaSiteKey = res.data?.google_captcha_secret_key || ""
            console.log(res.data)
        },
    },
})