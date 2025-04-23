import {defineStore} from 'pinia'
import {PanelApi} from "../api";
import router from "../plugins/router.ts";


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
            try {
                const res = await api.panelGet()
                this.setup = res.data.setup
                this.googleCaptchaSiteKey = res.data?.google_captcha_secret_key || ''
                console.log('setup:', this.setup)

                if (!this.setup) {
                    await router.push('/setup')
                }
            } catch (error) {
                console.error('Failed to fetch config:', error)
            }
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