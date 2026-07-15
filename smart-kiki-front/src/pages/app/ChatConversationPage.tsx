import { useState } from 'react'
import { useParams } from 'react-router-dom'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { listConversation, sendMessage } from '@/api/messages'
import { useAuthStore } from '@/stores/auth-store'
import { cn } from '@/lib/utils'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'

export function ChatConversationPage() {
  const { userId } = useParams<{ userId: string }>()
  const queryClient = useQueryClient()
  const currentUserId = useAuthStore((s) => s.user?.id)
  const [body, setBody] = useState('')

  const { data: messages } = useQuery({
    queryKey: ['messages', userId],
    queryFn: () => listConversation(userId!),
    enabled: !!userId,
    refetchInterval: 5000,
  })

  const mutation = useMutation({
    mutationFn: () => sendMessage(userId!, body),
    onSuccess: () => {
      setBody('')
      queryClient.invalidateQueries({ queryKey: ['messages', userId] })
    },
  })

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    if (!body.trim()) return
    mutation.mutate()
  }

  return (
    <div className="flex flex-1 flex-col">
      <div className="flex flex-1 flex-col gap-2 overflow-y-auto px-4 py-5">
        {messages?.map((message) => {
          const isMine = message.sender_id === currentUserId
          return (
            <div
              key={message.id}
              className={cn(
                'max-w-[75%] rounded-2xl px-3.5 py-2 text-sm',
                isMine ? 'self-end bg-primary text-primary-foreground' : 'self-start bg-muted',
              )}
            >
              {message.body}
            </div>
          )
        })}
      </div>

      <form onSubmit={handleSubmit} className="flex gap-2 border-t border-border p-3">
        <Input
          placeholder="Escreva uma mensagem..."
          value={body}
          onChange={(e) => setBody(e.target.value)}
        />
        <Button type="submit" disabled={mutation.isPending || !body.trim()}>
          Enviar
        </Button>
      </form>
    </div>
  )
}
