export interface Assessment {
  id: string
  trainer_id: string
  student_id: string
  recorded_at: string
  weight_kg: number | null
  notes: string
  created_at: string
}

export interface CreateAssessmentInput {
  recorded_at: string
  weight_kg?: number
  notes?: string
}
