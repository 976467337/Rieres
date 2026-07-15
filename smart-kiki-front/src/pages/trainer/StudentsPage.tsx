import { useState } from 'react'
import { Link } from 'react-router-dom'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { isAxiosError } from 'axios'
import { addStudent, getSubscription, listStudents } from '@/api/trainer'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'

export function StudentsPage() {
  const queryClient = useQueryClient()
  const [email, setEmail] = useState('')

  const { data: subscription } = useQuery({
    queryKey: ['subscription'],
    queryFn: getSubscription,
  })

  const { data: students } = useQuery({
    queryKey: ['students'],
    queryFn: listStudents,
  })

  const mutation = useMutation({
    mutationFn: addStudent,
    onSuccess: () => {
      setEmail('')
      queryClient.invalidateQueries({ queryKey: ['students'] })
      queryClient.invalidateQueries({ queryKey: ['subscription'] })
    },
  })

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    mutation.mutate(email)
  }

  const limitReached =
    isAxiosError(mutation.error) && mutation.error.response?.status === 403

  const limitLabel = subscription
    ? `${subscription.students_count} / ${subscription.student_limit ?? '∞'} alunos`
    : '—'

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <div>
        <h1 className="font-heading text-xl font-extrabold">Alunos</h1>
        <div className="mt-1 flex items-center gap-2 text-sm text-muted-foreground">
          <span>{limitLabel}</span>
          <span>·</span>
          <Link to="/app/plans" className="font-semibold text-primary">
            Ver planos
          </Link>
        </div>
      </div>

      <form onSubmit={handleSubmit} className="flex gap-2">
        <Input
          type="email"
          placeholder="E-mail do aluno"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <Button type="submit" disabled={mutation.isPending}>
          Adicionar
        </Button>
      </form>

      {mutation.isError && (
        <p className="text-sm text-destructive">
          {limitReached ? (
            <>
              Limite de alunos do plano atual atingido.{' '}
              <Link to="/app/plans" className="font-semibold underline">
                Fazer upgrade
              </Link>
            </>
          ) : (
            'Não foi possível adicionar esse aluno. Verifique o e-mail e tente novamente.'
          )}
        </p>
      )}

      <div className="flex flex-col gap-2.5">
        {students?.length === 0 && (
          <p className="text-sm text-muted-foreground">Você ainda não tem alunos cadastrados.</p>
        )}
        {students?.map((student) => (
          <Link key={student.id} to={`/app/students/${student.id}`}>
            <Card>
              <CardContent className="p-3.5">
                <div className="text-sm font-semibold">{student.name}</div>
                <div className="text-xs text-muted-foreground">{student.email}</div>
              </CardContent>
            </Card>
          </Link>
        ))}
      </div>
    </div>
  )
}
