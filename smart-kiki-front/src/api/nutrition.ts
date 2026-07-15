import { api } from '@/lib/api'
import type { NutritionPlan } from '@/types/nutrition'

export async function getNutritionPlanForStudent(studentId: string) {
  const { data } = await api.get<NutritionPlan>(`/trainer/students/${studentId}/nutrition-plan`)
  return data
}

export async function upsertNutritionPlan(studentId: string, content: string) {
  const { data } = await api.put<NutritionPlan>(`/trainer/students/${studentId}/nutrition-plan`, { content })
  return data
}

export async function getMyNutritionPlan() {
  const { data } = await api.get<NutritionPlan>('/students/me/nutrition-plan')
  return data
}
