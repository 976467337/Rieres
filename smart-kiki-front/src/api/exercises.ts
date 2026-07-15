import { api } from '@/lib/api'
import type { Exercise } from '@/types/exercise'

export async function listExercises(muscleGroup?: string) {
  const { data } = await api.get<Exercise[]>('/exercises', {
    params: muscleGroup ? { muscle_group: muscleGroup } : undefined,
  })
  return data
}
