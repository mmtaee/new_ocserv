import {defineStore} from 'pinia'

export type SnackbarColor = 'success' | 'error' | 'info' | 'warning'

export type SnackbarItem = {
    id: number
    message: string
    color: SnackbarColor
    timeout: number
}


export const useSnackbarStore = defineStore('snackbar', {
    state: () => ({
        snackbars: [] as SnackbarItem[],
    }),

    actions: {
        show(input: SnackbarItem | SnackbarItem[]) {
            const items = Array.isArray(input) ? input : [input]
            items.forEach(({id, message, color, timeout}) => {
                this.snackbars.push({
                    id: id + Math.floor(Math.random() * 1000000),
                    message,
                    color: color || 'info',
                    timeout: timeout || 3000,
                })
            })
        },

        remove(id: number) {
            this.snackbars = this.snackbars.filter(snack => snack.id !== id)
        },

    },
})
