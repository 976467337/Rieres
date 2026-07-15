import { api } from '@/lib/api'
import type { CreateWorkoutInput, Workout } from '@/types/workout'

export async function createWorkout(input: CreateWorkoutInput) {
  const { data } = await api.post<Workout>('/workouts', input)
  return data
}

export async function listWorkouts(studentId?: string) {
  const { data } = await api.get<Workout[]>('/workouts', {
    params: studentId ? { student_id: studentId } : undefined,
  })
  return data
}

export async function getWorkout(id: string) {
  const { data } = await api.get<Workout>(`/workouts/${id}`)
  return data
}

export async function deleteWorkout(id: string) {
  await api.delete(`/workouts/${id}`)
}

export async function completeWorkout(id: string) {
  await api.post(`/workouts/${id}/complete`)
}
