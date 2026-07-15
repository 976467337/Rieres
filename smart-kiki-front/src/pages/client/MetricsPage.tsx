import { useQuery } from '@tanstack/react-query'
import { getMyAssessments } from '@/api/assessments'
import { useAuthStore } from '@/stores/auth-store'
import { Card, CardContent } from '@/components/ui/card'
import { ProgressPhotoGallery } from '@/components/ProgressPhotoGallery'

export function MetricsPage() {
  const userId = useAuthStore((s) => s.user?.id)

  const { data: assessments } = useQuery({
    queryKey: ['my-assessments'],
    queryFn: getMyAssessments,
  })

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <h1 className="font-heading text-xl font-extrabold">Evolução</h1>

      {userId && <ProgressPhotoGallery studentId={userId} />}

      <div className="flex flex-col gap-2.5">
        {assessments?.length === 0 && (
          <p className="text-sm text-muted-foreground">
            Seu personal ainda não registrou nenhuma avaliação física.
          </p>
        )}
        {assessments?.map((assessment) => (
          <Card key={assessment.id}>
            <CardContent className="p-3.5">
              <div className="flex items-baseline justify-between">
                <div className="text-sm font-semibold">
                  {new Date(assessment.recorded_at).toLocaleDateString('pt-BR')}
                </div>
                {assessment.weight_kg != null && (
                  <div className="font-heading text-lg font-extrabold text-primary">
                    {assessment.weight_kg} kg
                  </div>
                )}
              </div>
              {assessment.notes && (
                <div className="mt-0.5 text-xs text-muted-foreground">{assessment.notes}</div>
              )}
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  )
}
