import { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { listStudents } from '@/api/trainer'
import { listWorkouts } from '@/api/workouts'
import { getNutritionPlanForStudent, upsertNutritionPlan } from '@/api/nutrition'
import { createAssessment } from '@/api/assessments'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { ProgressPhotoGallery } from '@/components/ProgressPhotoGallery'

export function StudentDetailPage() {
  const { studentId } = useParams<{ studentId: string }>()
  const queryClient = useQueryClient()

  const { data: students } = useQuery({
    queryKey: ['students'],
    queryFn: listStudents,
  })
  const student = students?.find((s) => s.id === studentId)

  const { data: workouts } = useQuery({
    queryKey: ['workouts', { studentId }],
    queryFn: () => listWorkouts(studentId),
    enabled: !!studentId,
  })

  const { data: nutritionPlan } = useQuery({
    queryKey: ['nutrition-plan', studentId],
    queryFn: () => getNutritionPlanForStudent(studentId!),
    enabled: !!studentId,
    retry: false,
  })

  const [nutritionContent, setNutritionContent] = useState('')
  useEffect(() => {
    if (nutritionPlan) setNutritionContent(nutritionPlan.content)
  }, [nutritionPlan])

  const nutritionMutation = useMutation({
    mutationFn: () => upsertNutritionPlan(studentId!, nutritionContent),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['nutrition-plan', studentId] }),
  })

  const [weight, setWeight] = useState('')
  const [notes, setNotes] = useState('')
  const assessmentMutation = useMutation({
    mutationFn: () =>
      createAssessment(studentId!, {
        recorded_at: new Date().toISOString(),
        weight_kg: weight ? Number(weight) : undefined,
        notes,
      }),
    onSuccess: () => {
      setWeight('')
      setNotes('')
    },
  })

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <div>
        <h1 className="font-heading text-xl font-extrabold">{student?.name ?? 'Aluno'}</h1>
        <p className="mt-0.5 text-sm text-muted-foreground">{student?.email}</p>
      </div>

      <Button asChild>
        <Link to={`/app/students/${studentId}/workouts/new`}>Criar treino</Link>
      </Button>

      <div>
        <div className="mb-2 text-sm font-bold">Treinos atribuídos</div>
        <div className="flex flex-col gap-2.5">
          {workouts?.length === 0 && (
            <p className="text-sm text-muted-foreground">Nenhum treino criado ainda pra esse aluno.</p>
          )}
          {workouts?.map((workout) => (
            <Card key={workout.id}>
              <CardContent className="p-3.5">
                <div className="text-sm font-semibold">{workout.name}</div>
                {workout.notes && <div className="mt-0.5 text-xs text-muted-foreground">{workout.notes}</div>}
              </CardContent>
            </Card>
          ))}
        </div>
      </div>

      {studentId && <ProgressPhotoGallery studentId={studentId} />}

      <div>
        <div className="mb-2 text-sm font-bold">Plano alimentar</div>
        <Card>
          <CardContent className="flex flex-col gap-2.5 p-3.5">
            <textarea
              value={nutritionContent}
              onChange={(e) => setNutritionContent(e.target.value)}
              placeholder="Escreva as orientações de alimentação pra esse aluno..."
              rows={5}
              className="w-full resize-none rounded-lg border border-border bg-card px-3 py-2 text-sm"
            />
            <Button
              type="button"
              size="sm"
              onClick={() => nutritionMutation.mutate()}
              disabled={nutritionMutation.isPending}
            >
              Salvar plano alimentar
            </Button>
          </CardContent>
        </Card>
      </div>

      <div>
        <div className="mb-2 text-sm font-bold">Nova avaliação física</div>
        <Card>
          <CardContent className="flex flex-col gap-2.5 p-3.5">
            <Input
              type="number"
              step="0.1"
              placeholder="Peso (kg)"
              value={weight}
              onChange={(e) => setWeight(e.target.value)}
            />
            <Input
              placeholder="Observações"
              value={notes}
              onChange={(e) => setNotes(e.target.value)}
            />
            <Button
              type="button"
              size="sm"
              onClick={() => assessmentMutation.mutate()}
              disabled={assessmentMutation.isPending}
            >
              {assessmentMutation.isSuccess ? 'Registrado!' : 'Registrar avaliação'}
            </Button>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
