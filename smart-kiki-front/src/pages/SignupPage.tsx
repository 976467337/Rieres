import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useMutation } from '@tanstack/react-query'
import { cn } from '@/lib/utils'
import { registerUser } from '@/api/auth'
import { useAuthStore } from '@/stores/auth-store'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import type { Role } from '@/types/user'

export function SignupPage() {
  const navigate = useNavigate()
  const setAuth = useAuthStore((s) => s.setAuth)

  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [role, setRole] = useState<Role>('client')

  const mutation = useMutation({
    mutationFn: registerUser,
    onSuccess: (data) => {
      setAuth(data.token, data.user)
      navigate('/app')
    },
  })

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    mutation.mutate({ name, email, password, role })
  }

  return (
    <div className="flex min-h-svh flex-col items-center justify-center gap-6 px-6 py-10">
      <div className="text-center">
        <h1 className="font-heading text-xl font-extrabold">Criar sua conta</h1>
        <p className="mt-1 text-sm text-muted-foreground">Comece sua jornada com a Kiki</p>
      </div>

      <form onSubmit={handleSubmit} className="flex w-full max-w-sm flex-col gap-3">
        <div className="flex rounded-full border border-border bg-card p-0.5">
          <button
            type="button"
            onClick={() => setRole('client')}
            className={cn(
              'flex-1 rounded-full py-2 text-xs font-bold transition-colors',
              role === 'client' ? 'bg-primary text-primary-foreground' : 'text-muted-foreground',
            )}
          >
            Sou aluno(a)
          </button>
          <button
            type="button"
            onClick={() => setRole('trainer')}
            className={cn(
              'flex-1 rounded-full py-2 text-xs font-bold transition-colors',
              role === 'trainer' ? 'bg-primary text-primary-foreground' : 'text-muted-foreground',
            )}
          >
            Sou personal
          </button>
        </div>

        <Input
          placeholder="Nome completo"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <Input
          type="email"
          placeholder="E-mail"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <Input
          type="password"
          placeholder="Senha"
          minLength={8}
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />

        {mutation.isError && (
          <p className="text-sm text-destructive">Não foi possível criar sua conta. Verifique os dados e tente novamente.</p>
        )}

        <Button type="submit" disabled={mutation.isPending} className="mt-2">
          {mutation.isPending ? 'Criando conta...' : 'Criar conta'}
        </Button>

        <Link to="/login" className="mt-2 text-center text-sm text-muted-foreground underline">
          Já tenho conta — Entrar
        </Link>
      </form>
    </div>
  )
}
