export type Role = 'client' | 'trainer'

export interface User {
  id: string
  name: string
  email: string
  role: Role
  created_at: string
  updated_at: string
}

export interface AuthResponse {
  token: string
  user: User
}
