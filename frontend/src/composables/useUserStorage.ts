import { useLocalStorage } from '@vueuse/core'

export const userStorage = useLocalStorage('user', {
  id: null as string | null,
  username: null as string | null,
  role: null as 'user' | 'admin' | null,
})
