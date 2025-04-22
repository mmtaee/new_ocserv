import {defineStore} from 'pinia'
import {PanelApi} from "../api";
import router from "../plugins/router.ts";

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
            console.log("setup: ", this.setup)
            if (!this.setup) {
                await router.push("/setup")
                return
            }
        },
        setSetup: function (bool: boolean) {
            this.setup = bool
        }
    },
})