import { useNavigate } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { useAuthStore } from '@/stores/auth-store'
import { listWorkouts } from '@/api/workouts'

export function ClientHomePage() {
  const user = useAuthStore((s) => s.user)
  const navigate = useNavigate()
  const firstName = user?.name?.split(' ')[0] ?? 'aluno(a)'

  const { data: workouts } = useQuery({
    queryKey: ['workouts'],
    queryFn: () => listWorkouts(),
  })
  const latestWorkout = workouts?.[0]

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <div>
        <h1 className="font-heading text-xl font-extrabold">Olá, {firstName}</h1>
        <p className="mt-0.5 text-sm text-muted-foreground">Pronta para o treino de hoje?</p>
      </div>

      <Card className="border-none bg-accent">
        <CardContent className="p-4">
          <div className="text-[11px] font-bold uppercase tracking-wide text-primary">Treino de hoje</div>
          {latestWorkout ? (
            <>
              <div className="mt-1.5 font-heading text-lg font-extrabold">{latestWorkout.name}</div>
              <Button className="mt-3.5" onClick={() => navigate(`/app/workouts/${latestWorkout.id}`)}>
                Iniciar treino
              </Button>
            </>
          ) : (
            <div className="mt-1.5 text-sm text-muted-foreground">
              Nenhum treino atribuído ainda — seu personal vai montar em breve.
            </div>
          )}
        </CardContent>
      </Card>

      <div className="flex gap-3">
        <Card className="flex-1">
          <CardContent className="p-3.5">
            <div className="font-heading text-2xl font-extrabold text-success">4</div>
            <div className="mt-0.5 text-xs text-muted-foreground">dias seguidos</div>
          </CardContent>
        </Card>
        <Card className="flex-1">
          <CardContent className="p-3.5">
            <div className="font-heading text-2xl font-extrabold text-success">3/4</div>
            <div className="mt-0.5 text-xs text-muted-foreground">treinos na semana</div>
          </CardContent>
        </Card>
      </div>

      <Card>
        <CardContent className="p-3.5">
          <p className="text-sm leading-snug">
            Dica da Kiki: beba água antes do treino — você está a 1 dia de bater seu recorde de sequência!
          </p>
        </CardContent>
      </Card>

      <div>
        <div className="mb-2 text-sm font-bold">Próxima sessão</div>
        <Card>
          <CardContent className="flex items-center gap-3 p-3.5">
            <div className="flex size-10 items-center justify-center rounded-xl bg-accent text-xs font-extrabold text-primary">
              QUI
            </div>
            <div className="text-sm font-semibold">Avaliação física com Ana</div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
