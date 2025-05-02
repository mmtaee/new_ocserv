import axios, {type AxiosInstance, type AxiosRequestConfig, type AxiosResponse} from 'axios'
import {type SnackbarItem, useSnackbarStore} from "@/stores/snackbar.ts";

const config: AxiosRequestConfig = {
    baseURL: `${import.meta.env.VITE_API_BASE_URL}` || 'http://localhost:8080'
}


const api: AxiosInstance = axios.create(config)


// api.interceptors.request.use(
//     (config: InternalAxiosRequestConfig): InternalAxiosRequestConfig => {
//
//         const token = localStorage.getItem('token') || "aaaaaaaa"
//
//         console.log("token", token)
//
//         config.headers = config.headers ?? {} as AxiosRequestHeaders
//         if (token) {
//             config.headers["Authorization"] = `Bearer ${token}`
//         }
//         console.log("config", config)
//         return config
//     },
//     (error) => {
//         return Promise.reject(error)
//     }
// )

// Response interceptor: Handle 401/400
api.interceptors.response.use(
    (response: AxiosResponse): AxiosResponse => {
        return response
    },
    (error) => {
        const {response} = error
        if (response) {
            if (response.status === 401) {
                localStorage.removeItem('token')
                return Promise.resolve(response)
            }
            if (response.status === 400) {
                const snackbar = useSnackbarStore()
                let errorList = response.data.Error

                if (Array.isArray(errorList)) {
                    const messages: SnackbarItem[] = errorList.map((message: string, index: number) => ({
                        id: index + 1,
                        message,
                        color: 'error',
                        timeout: (index + 1) * 1000 + 4000,
                    }))
                    snackbar.show(messages)
                } else {
                    snackbar.show({
                        id: 1,
                        message: typeof errorList === 'string' ? errorList : 'Unknown error',
                        color: 'error',
                        timeout: 4000,
                    })
                }
            }
        }
        return Promise.reject(error)
    }
)

export default api
