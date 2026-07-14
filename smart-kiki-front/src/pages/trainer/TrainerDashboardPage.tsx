import { Card, CardContent } from '@/components/ui/card'
import { useAuthStore } from '@/stores/auth-store'

const todaySessions = [
  { time: '09:00', label: 'Marina Souza — Pernas & Glúteos' },
  { time: '11:30', label: 'João Pedro — Avaliação física' },
  { time: '18:00', label: 'Marina Souza — Avaliação física' },
]

export function TrainerDashboardPage() {
  const user = useAuthStore((s) => s.user)
  const firstName = user?.name ?? 'Personal'

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <div>
        <h1 className="font-heading text-xl font-extrabold">Olá, {firstName}</h1>
        <p className="mt-0.5 text-sm text-muted-foreground">Você tem {todaySessions.length} sessões hoje</p>
      </div>

      <div className="flex gap-3">
        <Card className="flex-1">
          <CardContent className="p-3.5">
            <div className="font-heading text-2xl font-extrabold text-primary">18</div>
            <div className="mt-0.5 text-xs text-muted-foreground">alunos ativos</div>
          </CardContent>
        </Card>
        <Card className="flex-1">
          <CardContent className="p-3.5">
            <div className="font-heading text-2xl font-extrabold text-primary">{todaySessions.length}</div>
            <div className="mt-0.5 text-xs text-muted-foreground">sessões hoje</div>
          </CardContent>
        </Card>
      </div>

      <div>
        <div className="mb-2 text-sm font-bold">Agenda de hoje</div>
        <div className="flex flex-col gap-2.5">
          {todaySessions.map((session) => (
            <Card key={session.time}>
              <CardContent className="flex items-center gap-3 p-3">
                <div className="w-11 text-xs font-extrabold text-primary">{session.time}</div>
                <div className="text-sm">{session.label}</div>
              </CardContent>
            </Card>
          ))}
        </div>
      </div>

      <Card className="bg-accent">
        <CardContent className="p-3.5">
          <p className="text-sm leading-snug">
            Kiki: 2 alunos estão sem treino atualizado esta semana — dá uma olhada na aba Alunos.
          </p>
        </CardContent>
      </Card>
    </div>
  )
}
