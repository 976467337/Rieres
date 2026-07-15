import { api } from '@/lib/api'
import type { CreateSessionInput, Session } from '@/types/session'

export async function createSession(input: CreateSessionInput) {
  const { data } = await api.post<Session>('/sessions', input)
  return data
}

export async function listSessions() {
  const { data } = await api.get<Session[]>('/sessions')
  return data
}

export async function cancelSession(id: string) {
  await api.post(`/sessions/${id}/cancel`)
}
