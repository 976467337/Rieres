import { useRef } from 'react'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { apiOrigin } from '@/lib/api'
import { listProgressPhotos, uploadProgressPhoto } from '@/api/progress-photos'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'

interface ProgressPhotoGalleryProps {
  studentId: string
}

function formatDate(iso: string) {
  return new Date(iso).toLocaleDateString('pt-BR')
}

export function ProgressPhotoGallery({ studentId }: ProgressPhotoGalleryProps) {
  const queryClient = useQueryClient()
  const fileInputRef = useRef<HTMLInputElement>(null)

  const { data: photos } = useQuery({
    queryKey: ['progress-photos', studentId],
    queryFn: () => listProgressPhotos(studentId),
    enabled: !!studentId,
  })

  const mutation = useMutation({
    mutationFn: (file: File) => uploadProgressPhoto(studentId, file),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['progress-photos', studentId] }),
  })

  function handleFileChange(e: React.ChangeEvent<HTMLInputElement>) {
    const file = e.target.files?.[0]
    if (file) mutation.mutate(file)
    e.target.value = ''
  }

  const sorted = photos?.slice().sort((a, b) => a.recorded_at.localeCompare(b.recorded_at))
  const oldest = sorted?.[0]
  const newest = sorted && sorted.length > 1 ? sorted[sorted.length - 1] : undefined

  return (
    <div>
      <div className="mb-2 text-sm font-bold">Fotos de progresso</div>
      <Card>
        <CardContent className="flex flex-col gap-3 p-3.5">
          <input
            type="file"
            accept="image/*"
            ref={fileInputRef}
            onChange={handleFileChange}
            className="hidden"
          />
          <Button
            type="button"
            size="sm"
            variant="outline"
            onClick={() => fileInputRef.current?.click()}
            disabled={mutation.isPending}
          >
            {mutation.isPending ? 'Enviando...' : 'Enviar foto'}
          </Button>

          {mutation.isError && (
            <p className="text-xs text-destructive">Não foi possível enviar a foto. Verifique o arquivo e tente novamente.</p>
          )}

          {oldest && newest ? (
            <div>
              <div className="mb-1 text-xs font-semibold text-muted-foreground">Antes / Depois</div>
              <div className="grid grid-cols-2 gap-2">
                <div>
                  <img
                    src={`${apiOrigin}${oldest.image_path}`}
                    alt="Antes"
                    className="aspect-square w-full rounded-lg object-cover"
                  />
                  <div className="mt-1 text-center text-xs text-muted-foreground">{formatDate(oldest.recorded_at)}</div>
                </div>
                <div>
                  <img
                    src={`${apiOrigin}${newest.image_path}`}
                    alt="Depois"
                    className="aspect-square w-full rounded-lg object-cover"
                  />
                  <div className="mt-1 text-center text-xs text-muted-foreground">{formatDate(newest.recorded_at)}</div>
                </div>
              </div>
            </div>
          ) : oldest ? (
            <img
              src={`${apiOrigin}${oldest.image_path}`}
              alt="Foto de progresso"
              className="aspect-square w-full rounded-lg object-cover"
            />
          ) : (
            <p className="text-sm text-muted-foreground">Nenhuma foto ainda.</p>
          )}

          {sorted && sorted.length > 2 && (
            <div className="flex gap-2 overflow-x-auto">
              {sorted
                .slice()
                .reverse()
                .map((photo) => (
                  <img
                    key={photo.id}
                    src={`${apiOrigin}${photo.image_path}`}
                    alt={formatDate(photo.recorded_at)}
                    className="h-14 w-14 flex-shrink-0 rounded-md object-cover"
                  />
                ))}
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  )
}
