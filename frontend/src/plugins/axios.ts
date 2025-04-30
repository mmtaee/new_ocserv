import axios, {
    type AxiosInstance,
    type AxiosRequestConfig, type AxiosResponse
} from 'axios'

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
                return Promise.resolve(response) // still return response so caller can handle it
            }
            if (response.status === 400) {
                console.error('Unable to connect to the server')
            }
        }
        return Promise.reject(error)
    }
)

export default api
