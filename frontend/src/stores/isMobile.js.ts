// src/stores/isMobile.js
import {defineStore} from 'pinia';

export const useIsMobileStore = defineStore('isMobile', {
    state: () => ({
        isMobile: false,
        isTablet: false,
    }),
    actions: {
        setIsMobile(val: boolean) {
            this.isMobile = val;
        },
    }
});
