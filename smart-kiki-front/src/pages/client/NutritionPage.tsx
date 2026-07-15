import { useQuery } from '@tanstack/react-query'
import { getMyNutritionPlan } from '@/api/nutrition'
import { Card, CardContent } from '@/components/ui/card'

export function NutritionPage() {
  const { data: plan } = useQuery({
    queryKey: ['my-nutrition-plan'],
    queryFn: getMyNutritionPlan,
    retry: false,
  })

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <h1 className="font-heading text-xl font-extrabold">Nutrição</h1>

      <Card>
        <CardContent className="p-4">
          {plan?.content ? (
            <p className="whitespace-pre-line text-sm leading-relaxed">{plan.content}</p>
          ) : (
            <p className="text-sm text-muted-foreground">
              Seu personal ainda não montou um plano alimentar pra você.
            </p>
          )}
        </CardContent>
      </Card>
    </div>
  )
}
