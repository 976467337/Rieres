import { useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { useMutation, useQuery } from '@tanstack/react-query'
import { Plus, X } from 'lucide-react'
import { listExercises } from '@/api/exercises'
import { createWorkout } from '@/api/workouts'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import type { WorkoutExerciseInput } from '@/types/workout'

interface BuilderItem extends WorkoutExerciseInput {
  name: string
}

export function WorkoutBuilderPage() {
  const { studentId } = useParams<{ studentId: string }>()
  const navigate = useNavigate()

  const [name, setName] = useState('')
  const [search, setSearch] = useState('')
  const [items, setItems] = useState<BuilderItem[]>([])

  const { data: exercises } = useQuery({
    queryKey: ['exercises', null],
    queryFn: () => listExercises(),
  })

  const filtered = exercises?.filter((ex) =>
    ex.name.toLowerCase().includes(search.toLowerCase()),
  )

  function addExercise(exerciseId: string, exerciseName: string) {
    if (items.some((item) => item.exercise_id === exerciseId)) return
    setItems((prev) => [...prev, { exercise_id: exerciseId, name: exerciseName, sets: 3, reps: 12 }])
  }

  function removeExercise(exerciseId: string) {
    setItems((prev) => prev.filter((item) => item.exercise_id !== exerciseId))
  }

  function updateItem(exerciseId: string, field: 'sets' | 'reps', value: number) {
    setItems((prev) =>
      prev.map((item) => (item.exercise_id === exerciseId ? { ...item, [field]: value } : item)),
    )
  }

  const mutation = useMutation({
    mutationFn: createWorkout,
    onSuccess: () => navigate(`/app/students/${studentId}`),
  })

  function handleSave() {
    if (!studentId || !name || items.length === 0) return
    mutation.mutate({
      student_id: studentId,
      name,
      exercises: items.map(({ exercise_id, sets, reps }) => ({ exercise_id, sets, reps })),
    })
  }

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <h1 className="font-heading text-xl font-extrabold">Criar treino</h1>

      <Input placeholder="Nome do treino" value={name} onChange={(e) => setName(e.target.value)} />

      {items.length > 0 && (
        <div className="flex flex-col gap-2">
          {items.map((item) => (
            <Card key={item.exercise_id}>
              <CardContent className="flex items-center gap-2 p-3">
                <div className="flex-1 text-sm font-semibold">{item.name}</div>
                <Input
                  type="number"
                  min={1}
                  value={item.sets}
                  onChange={(e) => updateItem(item.exercise_id, 'sets', Number(e.target.value))}
                  className="w-14"
                  aria-label="Séries"
                />
                <span className="text-xs text-muted-foreground">×</span>
                <Input
                  type="number"
                  min={1}
                  value={item.reps}
                  onChange={(e) => updateItem(item.exercise_id, 'reps', Number(e.target.value))}
                  className="w-14"
                  aria-label="Repetições"
                />
                <Button
                  type="button"
                  variant="ghost"
                  size="icon-sm"
                  onClick={() => removeExercise(item.exercise_id)}
                >
                  <X size={14} />
                </Button>
              </CardContent>
            </Card>
          ))}
        </div>
      )}

      <Input
        placeholder="Buscar exercício..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />

      <div className="flex flex-col gap-2">
        {filtered?.map((exercise) => (
          <Card key={exercise.id}>
            <CardContent className="flex items-center gap-2 p-3">
              <div className="flex-1">
                <div className="text-sm font-semibold">{exercise.name}</div>
                <div className="text-xs text-muted-foreground">{exercise.muscle_group}</div>
              </div>
              <Button
                type="button"
                variant="outline"
                size="icon-sm"
                onClick={() => addExercise(exercise.id, exercise.name)}
              >
                <Plus size={14} />
              </Button>
            </CardContent>
          </Card>
        ))}
      </div>

      <Button
        onClick={handleSave}
        disabled={!name || items.length === 0 || mutation.isPending}
      >
        Salvar treino
      </Button>
    </div>
  )
}
