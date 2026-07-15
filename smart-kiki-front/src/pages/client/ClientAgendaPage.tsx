import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { cancelSession, listSessions } from '@/api/sessions'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'

export function ClientAgendaPage() {
  const queryClient = useQueryClient()

  const { data: sessions } = useQuery({ queryKey: ['sessions'], queryFn: listSessions })

  const cancelMutation = useMutation({
    mutationFn: cancelSession,
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['sessions'] }),
  })

  const sorted = sessions?.slice().sort((a, b) => a.scheduled_at.localeCompare(b.scheduled_at))

  function statusLabel(session: NonNullable<typeof sessions>[number]) {
    if (session.status === 'cancelada') return 'Cancelada'
    return new Date(session.scheduled_at) < new Date() ? 'Concluída' : 'Agendada'
  }

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <h1 className="font-heading text-xl font-extrabold">Agenda</h1>

      <div className="flex flex-col gap-2.5">
        {sorted?.length === 0 && (
          <p className="text-sm text-muted-foreground">Nenhuma sessão marcada ainda.</p>
        )}
        {sorted?.map((session) => {
          const label = statusLabel(session)
          const canCancel = label === 'Agendada'
          return (
            <Card key={session.id}>
              <CardContent className="flex items-center gap-3 p-3.5">
                <div className="flex-1">
                  <div className="text-sm font-semibold">
                    {new Date(session.scheduled_at).toLocaleString('pt-BR', {
                      weekday: 'short',
                      day: '2-digit',
                      month: '2-digit',
                      hour: '2-digit',
                      minute: '2-digit',
                    })}
                  </div>
                  <div className="text-xs text-muted-foreground">
                    {label} · {session.duration_minutes} min
                  </div>
                </div>
                {canCancel && (
                  <Button
                    type="button"
                    variant="ghost"
                    size="sm"
                    onClick={() => cancelMutation.mutate(session.id)}
                  >
                    Cancelar
                  </Button>
                )}
              </CardContent>
            </Card>
          )
        })}
      </div>
    </div>
  )
}
