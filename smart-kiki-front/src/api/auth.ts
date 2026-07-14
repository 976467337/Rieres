import { api } from '@/lib/api'
import type { AuthResponse, Role, User } from '@/types/user'

export interface RegisterPayload {
  name: string
  email: string
  password: string
  role: Role
}

export interface LoginPayload {
  email: string
  password: string
}

export async function registerUser(payload: RegisterPayload) {
  const { data } = await api.post<AuthResponse>('/auth/register', payload)
  return data
}

export async function loginUser(payload: LoginPayload) {
  const { data } = await api.post<AuthResponse>('/auth/login', payload)
  return data
}

export async function fetchMe() {
  const { data } = await api.get<User>('/users/me')
  return data
}
