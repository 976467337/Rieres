import { api } from '@/lib/api'
import type { Plan } from '@/types/subscription'
import type { User } from '@/types/user'

export interface MarketplaceTrainer extends User {
  plan: Plan
}

export async function listVisibleTrainers() {
  const { data } = await api.get<MarketplaceTrainer[]>('/marketplace/trainers')
  return data
}

export async function requestTrainer(trainerId: string) {
  await api.post(`/marketplace/trainers/${trainerId}/request`)
}
