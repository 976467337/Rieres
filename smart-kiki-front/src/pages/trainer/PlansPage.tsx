import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { getSubscription, changePlan } from '@/api/trainer'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { cn } from '@/lib/utils'
import type { Plan } from '@/types/subscription'

const PLANS: { key: Plan; label: string; price: string; limit: string; visibility: string }[] = [
  { key: 'free', label: 'Grátis', price: 'R$ 0/mês', limit: 'até 2 alunos', visibility: 'não aparece na lista de busca' },
  { key: 'plus', label: 'Plus', price: 'R$ 10/mês', limit: 'até 10 alunos', visibility: 'aparece na busca nos primeiros 7 dias de cada ciclo' },
  { key: 'pro', label: 'Pro', price: 'R$ 20/mês', limit: 'alunos ilimitados', visibility: 'sempre aparece na lista de busca' },
]

export function PlansPage() {
  const queryClient = useQueryClient()

  const { data: subscription } = useQuery({
    queryKey: ['subscription'],
    queryFn: getSubscription,
  })

  const mutation = useMutation({
    mutationFn: changePlan,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['subscription'] })
    },
  })

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <div>
        <h1 className="font-heading text-xl font-extrabold">Planos</h1>
        <p className="mt-0.5 text-sm text-muted-foreground">
          Escolha quantos alunos você pode ter e se seu perfil aparece na busca de personais.
        </p>
      </div>

      <div className="flex flex-col gap-3">
        {PLANS.map((plan) => {
          const isCurrent = subscription?.plan === plan.key
          return (
            <Card key={plan.key} className={cn(isCurrent && 'ring-2 ring-primary')}>
              <CardContent className="flex items-center justify-between gap-3 p-4">
                <div>
                  <div className="font-heading text-base font-extrabold">{plan.label}</div>
                  <div className="mt-0.5 text-sm text-muted-foreground">{plan.price}</div>
                  <div className="mt-1.5 text-xs text-muted-foreground">{plan.limit}</div>
                  <div className="mt-0.5 text-xs text-muted-foreground">{plan.visibility}</div>
                </div>
                <Button
                  variant={isCurrent ? 'secondary' : 'default'}
                  disabled={isCurrent || mutation.isPending}
                  onClick={() => mutation.mutate(plan.key)}
                >
                  {isCurrent ? 'Plano atual' : 'Assinar'}
                </Button>
              </CardContent>
            </Card>
          )
        })}
      </div>
    </div>
  )
}
