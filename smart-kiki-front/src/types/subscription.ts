export type Plan = 'free' | 'plus' | 'pro'

export interface Subscription {
  plan: Plan
  price: number
  student_limit: number | null
  students_count: number
  visible_in_marketplace: boolean
  current_period_start: string
}
