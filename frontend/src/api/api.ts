import axios from 'axios'
import router from '@/router'
import { userStorage } from '@/composables/useUserStorage.ts'

const baseURL = import.meta.env.VITE_API_URL
const api = axios.create({
  baseURL: baseURL,
  headers: {
    'Content-Type': 'application/json',
  },
})
api.interceptors.request.use((config) => {
  config.withCredentials = true
  return config
})
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response.status === 401) {
      try {
        const res = await axios.post(
          `${baseURL}/auth/refresh`,
          {},
          { withCredentials: true },
        )
        if (res.status === 200) {
          return api(error.config)
        }
      } catch (error) {
        if (axios.isAxiosError(error) && error.response?.status === 401) {
          userStorage.value = {
            id: null,
            username: null,
            role: null,
          }
          router.push('/login')
        }
      }
    }
    return Promise.reject(error)
  },
)
export default api
