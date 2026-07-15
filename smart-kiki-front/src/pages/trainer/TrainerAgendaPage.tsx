import { useState } from 'react'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { listStudents } from '@/api/trainer'
import { cancelSession, createSession, listSessions } from '@/api/sessions'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'

export function TrainerAgendaPage() {
  const queryClient = useQueryClient()
  const [open, setOpen] = useState(false)
  const [studentId, setStudentId] = useState('')
  const [scheduledAt, setScheduledAt] = useState('')
  const [duration, setDuration] = useState(60)

  const { data: students } = useQuery({ queryKey: ['students'], queryFn: listStudents })
  const { data: sessions } = useQuery({ queryKey: ['sessions'], queryFn: listSessions })

  const createMutation = useMutation({
    mutationFn: createSession,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['sessions'] })
      setOpen(false)
      setStudentId('')
      setScheduledAt('')
    },
  })

  const cancelMutation = useMutation({
    mutationFn: cancelSession,
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['sessions'] }),
  })

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    if (!studentId || !scheduledAt) return
    createMutation.mutate({
      student_id: studentId,
      scheduled_at: new Date(scheduledAt).toISOString(),
      duration_minutes: duration,
    })
  }

  const studentName = (id: string) => students?.find((s) => s.id === id)?.name ?? 'Aluno'

  const upcoming = sessions
    ?.filter((s) => s.status === 'agendada')
    .sort((a, b) => a.scheduled_at.localeCompare(b.scheduled_at))

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <div className="flex items-center justify-between">
        <h1 className="font-heading text-xl font-extrabold">Agenda</h1>
        <Dialog open={open} onOpenChange={setOpen}>
          <DialogTrigger asChild>
            <Button size="sm">Agendar sessão</Button>
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Nova sessão</DialogTitle>
            </DialogHeader>
            <form onSubmit={handleSubmit} className="flex flex-col gap-3">
              <select
                value={studentId}
                onChange={(e) => setStudentId(e.target.value)}
                required
                className="rounded-lg border border-border bg-card px-3 py-2 text-sm"
              >
                <option value="">Selecione o aluno</option>
                {students?.map((s) => (
                  <option key={s.id} value={s.id}>
                    {s.name}
                  </option>
                ))}
              </select>
              <Input
                type="datetime-local"
                value={scheduledAt}
                onChange={(e) => setScheduledAt(e.target.value)}
                required
              />
              <Input
                type="number"
                min={15}
                step={15}
                value={duration}
                onChange={(e) => setDuration(Number(e.target.value))}
                placeholder="Duração (min)"
              />
              <Button type="submit" disabled={createMutation.isPending}>
                Salvar
              </Button>
            </form>
          </DialogContent>
        </Dialog>
      </div>

      <div className="flex flex-col gap-2.5">
        {upcoming?.length === 0 && (
          <p className="text-sm text-muted-foreground">Nenhuma sessão agendada.</p>
        )}
        {upcoming?.map((session) => (
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
                  {studentName(session.student_id)} · {session.duration_minutes} min
                </div>
              </div>
              <Button
                type="button"
                variant="ghost"
                size="sm"
                onClick={() => cancelMutation.mutate(session.id)}
              >
                Cancelar
              </Button>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  )
}
