import {defineStore} from 'pinia'
import {PanelApi} from "../api";

export const useConfigStore = defineStore('config', {
    state: () => ({
        setup: false,
        googleCaptchaSiteKey: "",
    }),

    actions: {
        fetchConfig: function () {
            const api = new PanelApi()
            api.panelGet().then((res) => {
                this.setup = res.data.setup
                this.googleCaptchaSiteKey = res.data?.google_captcha_secret_key || ""
            })
        },
        setSetupComplete: async function () {
            this.setup = true;
        }
    },
})