import type { Exercise } from '@/types/exercise'

export interface WorkoutExercise {
  id: string
  exercise_id: string
  position: number
  sets: number
  reps: number
  load: string
  rest_seconds: number
  notes: string
  exercise: Exercise
}

export interface Workout {
  id: string
  trainer_id: string
  student_id: string
  name: string
  notes: string
  created_at: string
  exercises?: WorkoutExercise[]
}

export interface WorkoutExerciseInput {
  exercise_id: string
  sets: number
  reps: number
  load?: string
  rest_seconds?: number
  notes?: string
}

export interface CreateWorkoutInput {
  student_id: string
  name: string
  notes?: string
  exercises: WorkoutExerciseInput[]
}
