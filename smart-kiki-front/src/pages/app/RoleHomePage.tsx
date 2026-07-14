import { useAuthStore } from '@/stores/auth-store'
import { ClientHomePage } from '@/pages/client/ClientHomePage'
import { TrainerDashboardPage } from '@/pages/trainer/TrainerDashboardPage'

export function RoleHomePage() {
  const role = useAuthStore((s) => s.user?.role)
  return role === 'trainer' ? <TrainerDashboardPage /> : <ClientHomePage />
}
