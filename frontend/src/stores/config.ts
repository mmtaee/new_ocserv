import {defineStore} from 'pinia'

export const useConfigStore = defineStore('config', {
    state: () => ({
        setup: false,
        configs: {
            googleCaptchaSiteKey: "",
        },
    }),

    actions: {
        fetchConfig: async function () {
            // const data = await config()
            // this.setup = data.setup
            // this.configs = data.config
        },
        setSetupComplete: async function () {
            this.setup = true;
        }
    },
})