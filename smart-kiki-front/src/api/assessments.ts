import { api } from '@/lib/api'
import type { Assessment, CreateAssessmentInput } from '@/types/assessment'

export async function createAssessment(studentId: string, input: CreateAssessmentInput) {
  const { data } = await api.post<Assessment>(`/trainer/students/${studentId}/assessments`, input)
  return data
}

export async function getMyAssessments() {
  const { data } = await api.get<Assessment[]>('/students/me/assessments')
  return data
}
