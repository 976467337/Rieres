import { useParams } from 'react-router-dom'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { completeWorkout, getWorkout } from '@/api/workouts'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'

export function WorkoutDetailPage() {
  const { id } = useParams<{ id: string }>()
  const queryClient = useQueryClient()

  const { data: workout } = useQuery({
    queryKey: ['workout', id],
    queryFn: () => getWorkout(id!),
    enabled: !!id,
  })

  const mutation = useMutation({
    mutationFn: () => completeWorkout(id!),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['workouts'] })
    },
  })

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <div>
        <h1 className="font-heading text-xl font-extrabold">{workout?.name}</h1>
        {workout?.notes && <p className="mt-0.5 text-sm text-muted-foreground">{workout.notes}</p>}
      </div>

      <div className="flex flex-col gap-2.5">
        {workout?.exercises?.map((item) => (
          <Card key={item.id}>
            <CardContent className="p-3.5">
              <div className="text-sm font-semibold">{item.exercise.name}</div>
              <div className="mt-0.5 text-xs text-muted-foreground">
                {item.sets} séries × {item.reps} repetições
                {item.rest_seconds > 0 && ` · ${item.rest_seconds}s descanso`}
              </div>
            </CardContent>
          </Card>
        ))}
      </div>

      <Button onClick={() => mutation.mutate()} disabled={mutation.isPending || mutation.isSuccess}>
        {mutation.isSuccess ? 'Treino concluído!' : 'Concluir treino'}
      </Button>
    </div>
  )
}
