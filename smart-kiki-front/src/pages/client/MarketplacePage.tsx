import { useState } from 'react'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { isAxiosError } from 'axios'
import { listVisibleTrainers, requestTrainer } from '@/api/marketplace'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Badge } from '@/components/ui/badge'

export function MarketplacePage() {
  const queryClient = useQueryClient()
  const [search, setSearch] = useState('')
  const [requestedIds, setRequestedIds] = useState<string[]>([])
  const [errorId, setErrorId] = useState<string | null>(null)

  const { data: trainers } = useQuery({
    queryKey: ['marketplace-trainers'],
    queryFn: listVisibleTrainers,
  })

  const mutation = useMutation({
    mutationFn: requestTrainer,
    onSuccess: (_data, trainerId) => {
      setRequestedIds((prev) => [...prev, trainerId])
      setErrorId(null)
      queryClient.invalidateQueries({ queryKey: ['my-trainers'] })
    },
    onError: (error, trainerId) => {
      setErrorId(trainerId)
      if (isAxiosError(error) && error.response?.status === 409) {
        setRequestedIds((prev) => [...prev, trainerId])
      }
    },
  })

  const filtered = trainers?.filter((t) =>
    t.name.toLowerCase().includes(search.toLowerCase()),
  )

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <div>
        <h1 className="font-heading text-xl font-extrabold">Buscar personal</h1>
        <p className="mt-0.5 text-sm text-muted-foreground">Encontre um personal trainer pra te acompanhar.</p>
      </div>

      <Input
        placeholder="Buscar por nome..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />

      <div className="flex flex-col gap-2.5">
        {filtered?.length === 0 && (
          <p className="text-sm text-muted-foreground">Nenhum personal disponível no momento.</p>
        )}
        {filtered?.map((trainer) => {
          const requested = requestedIds.includes(trainer.id)
          return (
            <Card key={trainer.id}>
              <CardContent className="flex items-center gap-3 p-3.5">
                <div className="flex-1">
                  <div className="flex items-center gap-1.5">
                    <span className="text-sm font-semibold">{trainer.name}</span>
                    {trainer.plan === 'pro' && <Badge>Pro</Badge>}
                  </div>
                  <div className="text-xs text-muted-foreground">{trainer.email}</div>
                  {errorId === trainer.id && !requested && (
                    <div className="mt-1 text-xs text-destructive">
                      Não foi possível solicitar esse personal agora.
                    </div>
                  )}
                </div>
                <Button
                  type="button"
                  size="sm"
                  variant={requested ? 'secondary' : 'default'}
                  disabled={requested || mutation.isPending}
                  onClick={() => mutation.mutate(trainer.id)}
                >
                  {requested ? 'Solicitado' : 'Solicitar'}
                </Button>
              </CardContent>
            </Card>
          )
        })}
      </div>
    </div>
  )
}
