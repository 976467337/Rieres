import { Link, useNavigate } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { useAuthStore } from '@/stores/auth-store'
import { getSubscription } from '@/api/trainer'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'

const PLAN_LABEL: Record<string, string> = {
  free: 'Grátis',
  plus: 'Plus',
  pro: 'Pro',
}

export function RoleProfilePage() {
  const user = useAuthStore((s) => s.user)
  const logout = useAuthStore((s) => s.logout)
  const navigate = useNavigate()
  const isTrainer = user?.role === 'trainer'

  const { data: subscription } = useQuery({
    queryKey: ['subscription'],
    queryFn: getSubscription,
    enabled: isTrainer,
  })

  function handleLogout() {
    logout()
    navigate('/login')
  }

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <h1 className="font-heading text-xl font-extrabold">Perfil</h1>

      <Card>
        <CardContent className="flex flex-col gap-1 p-4">
          <div className="text-base font-semibold">{user?.name}</div>
          <div className="text-sm text-muted-foreground">{user?.email}</div>
          <div className="text-xs text-muted-foreground">{isTrainer ? 'Personal trainer' : 'Aluno(a)'}</div>
        </CardContent>
      </Card>

      {isTrainer && (
        <Link to="/app/plans">
          <Card>
            <CardContent className="p-4">
              <div className="text-xs font-bold uppercase tracking-wide text-primary">Meu plano</div>
              <div className="mt-1 text-base font-semibold">
                {subscription ? PLAN_LABEL[subscription.plan] : '—'}
              </div>
              <div className="mt-0.5 text-xs text-muted-foreground">Toque pra ver ou trocar de plano</div>
            </CardContent>
          </Card>
        </Link>
      )}

      <Button variant="outline" onClick={handleLogout}>
        Sair
      </Button>
    </div>
  )
}
