import { useAuthStore } from '@/stores/auth-store'
import { TrainerAgendaPage } from '@/pages/trainer/TrainerAgendaPage'
import { ClientAgendaPage } from '@/pages/client/ClientAgendaPage'

export function RoleAgendaPage() {
  const role = useAuthStore((s) => s.user?.role)
  return role === 'trainer' ? <TrainerAgendaPage /> : <ClientAgendaPage />
}
