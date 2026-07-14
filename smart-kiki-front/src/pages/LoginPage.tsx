import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useMutation } from '@tanstack/react-query'
import { Moon, Sun } from 'lucide-react'
import { loginUser } from '@/api/auth'
import { useAuthStore } from '@/stores/auth-store'
import { useThemeStore } from '@/stores/theme-store'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import kikiMascot from '@/assets/kiki-mascot.png'

export function LoginPage() {
  const navigate = useNavigate()
  const setAuth = useAuthStore((s) => s.setAuth)
  const { theme, toggleTheme } = useThemeStore()

  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const mutation = useMutation({
    mutationFn: loginUser,
    onSuccess: (data) => {
      setAuth(data.token, data.user)
      navigate('/app')
    },
  })

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    mutation.mutate({ email, password })
  }

  return (
    <div className="relative flex min-h-svh flex-col items-center justify-center gap-6 px-6 py-10">
      <button
        type="button"
        onClick={toggleTheme}
        aria-label="Alternar tema"
        className="absolute right-4 top-4 flex size-9 items-center justify-center rounded-full border border-border bg-card text-foreground"
      >
        {theme === 'dark' ? <Moon size={16} /> : <Sun size={16} />}
      </button>

      <img
        src={kikiMascot}
        alt="Kiki"
        className="h-44 w-auto animate-[kiki-bob_2.6s_ease-in-out_infinite]"
      />

      <div className="text-center">
        <h1 className="font-heading text-2xl font-extrabold">Smart Kiki</h1>
        <p className="mt-1 text-sm text-muted-foreground">Seu personal, sempre com você</p>
      </div>

      <form onSubmit={handleSubmit} className="flex w-full max-w-sm flex-col gap-3">
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
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />

        {mutation.isError && (
          <p className="text-sm text-destructive">E-mail ou senha inválidos.</p>
        )}

        <Button type="submit" disabled={mutation.isPending} className="mt-2">
          {mutation.isPending ? 'Entrando...' : 'Entrar'}
        </Button>

        <Link
          to="/signup"
          className="mt-2 text-center text-sm font-bold text-primary"
        >
          Não tem conta? Criar conta
        </Link>
      </form>
    </div>
  )
}
