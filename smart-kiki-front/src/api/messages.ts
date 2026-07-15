import { api } from '@/lib/api'
import type { Message } from '@/types/message'
import type { User } from '@/types/user'

export async function sendMessage(toUserId: string, body: string) {
  const { data } = await api.post<Message>('/messages', { to_user_id: toUserId, body })
  return data
}

export async function listConversation(withUserId: string) {
  const { data } = await api.get<Message[]>('/messages', { params: { with: withUserId } })
  return data
}

export async function listMyTrainers() {
  const { data } = await api.get<User[]>('/students/me/trainers')
  return data
}
