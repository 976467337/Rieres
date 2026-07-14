import { api } from '@/lib/api'
import type { Plan, Subscription } from '@/types/subscription'
import type { User } from '@/types/user'

export async function getSubscription() {
  const { data } = await api.get<Subscription>('/trainer/subscription')
  return data
}

export async function changePlan(plan: Plan) {
  const { data } = await api.patch<Subscription>('/trainer/subscription', { plan })
  return data
}

export async function listStudents() {
  const { data } = await api.get<User[]>('/trainer/students')
  return data
}

export async function addStudent(email: string) {
  const { data } = await api.post<User>('/trainer/students', { email })
  return data
}
