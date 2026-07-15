export type SessionStatus = 'agendada' | 'cancelada'

export interface Session {
  id: string
  trainer_id: string
  student_id: string
  scheduled_at: string
  duration_minutes: number
  notes: string
  status: SessionStatus
  created_at: string
}

export interface CreateSessionInput {
  student_id: string
  scheduled_at: string
  duration_minutes?: number
  notes?: string
}
