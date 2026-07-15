import { Link } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { listStudents } from '@/api/trainer'
import { listMyTrainers } from '@/api/messages'
import { useAuthStore } from '@/stores/auth-store'
import { Card, CardContent } from '@/components/ui/card'

export function RoleChatListPage() {
  const role = useAuthStore((s) => s.user?.role)
  const isTrainer = role === 'trainer'

  const { data: contacts } = useQuery({
    queryKey: isTrainer ? ['students'] : ['my-trainers'],
    queryFn: isTrainer ? listStudents : listMyTrainers,
  })

  return (
    <div className="flex flex-col gap-4 px-4 py-5">
      <h1 className="font-heading text-xl font-extrabold">Chat</h1>

      <div className="flex flex-col gap-2.5">
        {contacts?.length === 0 && (
          <p className="text-sm text-muted-foreground">
            {isTrainer ? 'Você ainda não tem alunos pra conversar.' : 'Você ainda não tem um personal vinculado.'}
          </p>
        )}
        {contacts?.map((contact) => (
          <Link key={contact.id} to={`/app/chat/${contact.id}`}>
            <Card>
              <CardContent className="p-3.5">
                <div className="text-sm font-semibold">{contact.name}</div>
                <div className="text-xs text-muted-foreground">{contact.email}</div>
              </CardContent>
            </Card>
          </Link>
        ))}
      </div>
    </div>
  )
}
