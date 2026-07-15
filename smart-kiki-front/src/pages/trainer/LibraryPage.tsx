import { useState } from 'react'
import { useQuery } from '@tanstack/react-query'
import { listExercises } from '@/api/exercises'
import { Card, CardContent } from '@/components/ui/card'
import { cn } from '@/lib/utils'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import type { Exercise } from '@/types/exercise'

const MUSCLE_GROUPS = ['Peito', 'Costas', 'Pernas', 'Ombro', 'Bíceps', 'Tríceps', 'Core', 'Cardio']

export function LibraryPage() {
  const [muscleGroup, setMuscleGroup] = useState<string | null>(null)
  const [selected, setSelected] = useState<Exercise | null>(null)

  const { data: exercises } = useQuery({
    queryKey: ['exercises', muscleGroup],
    queryFn: () => listExercises(muscleGroup ?? undefined),
  })

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <div>
        <h1 className="font-heading text-xl font-extrabold">Biblioteca</h1>
        <p className="mt-0.5 text-sm text-muted-foreground">Escolha um exercício pra ver as instruções.</p>
      </div>

      <div className="flex gap-2 overflow-x-auto pb-1">
        <button
          type="button"
          onClick={() => setMuscleGroup(null)}
          className={cn(
            'flex-shrink-0 rounded-full border px-3 py-1.5 text-xs font-semibold',
            muscleGroup === null ? 'border-primary bg-primary text-primary-foreground' : 'border-border text-muted-foreground',
          )}
        >
          Todos
        </button>
        {MUSCLE_GROUPS.map((group) => (
          <button
            key={group}
            type="button"
            onClick={() => setMuscleGroup(group)}
            className={cn(
              'flex-shrink-0 rounded-full border px-3 py-1.5 text-xs font-semibold',
              muscleGroup === group ? 'border-primary bg-primary text-primary-foreground' : 'border-border text-muted-foreground',
            )}
          >
            {group}
          </button>
        ))}
      </div>

      <div className="flex flex-col gap-2.5">
        {exercises?.map((exercise) => (
          <Card key={exercise.id} className="cursor-pointer" onClick={() => setSelected(exercise)}>
            <CardContent className="p-3.5">
              <div className="text-sm font-semibold">{exercise.name}</div>
              <div className="mt-0.5 text-xs text-muted-foreground">
                {exercise.muscle_group} · {exercise.equipment}
              </div>
            </CardContent>
          </Card>
        ))}
      </div>

      <Dialog open={!!selected} onOpenChange={(open) => !open && setSelected(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{selected?.name}</DialogTitle>
          </DialogHeader>
          <div className="text-sm text-muted-foreground">
            {selected?.muscle_group} · {selected?.equipment}
          </div>
          <p className="text-sm leading-snug">{selected?.instructions}</p>
        </DialogContent>
      </Dialog>
    </div>
  )
}
