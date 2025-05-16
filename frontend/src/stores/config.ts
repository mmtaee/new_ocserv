import {defineStore} from 'pinia'
import {PanelApi} from "@/api";
import router from "@/plugins/router.ts";


interface ConfigState {
    setup: boolean
    googleCaptchaSiteKey: string
}

export const useConfigStore = defineStore('config', {
    state: (): ConfigState => ({
        setup: false,
        googleCaptchaSiteKey: "",
    }),

    actions: {
        async fetchConfig() {
            const api = new PanelApi()
            await api.panelInitGet().then((response) => {
                this.setup = response.data.setup
                this.googleCaptchaSiteKey = response.data?.google_captcha_site_key || ''
                if (!this.setup) {
                    router.push('/setup')
                }
            })
        },
        setSetup: function (bool: boolean) {
            this.setup = bool
        }
    },

    getters: {
        config(state): ConfigState {
            return {
                setup: state.setup,
                googleCaptchaSiteKey: state.googleCaptchaSiteKey,
            }
        },
    },
})