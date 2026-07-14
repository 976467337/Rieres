import { NavLink, Outlet } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import {
  Bell,
  Calendar,
  LayoutDashboard,
  LineChart,
  MessageCircle,
  Moon,
  Salad,
  Search,
  Sun,
  Trophy,
  User,
  Users,
  BookOpen,
} from 'lucide-react'
import { cn } from '@/lib/utils'
import { fetchMe } from '@/api/auth'
import { useAuthStore } from '@/stores/auth-store'
import { useThemeStore } from '@/stores/theme-store'

const clientTabs = [
  { to: '/app/home', label: 'Início', icon: LayoutDashboard },
  { to: '/app/agenda', label: 'Agenda', icon: Calendar },
  { to: '/app/marketplace', label: 'Buscar', icon: Search },
  { to: '/app/metrics', label: 'Evolução', icon: LineChart },
  { to: '/app/challenges', label: 'Desafios', icon: Trophy },
  { to: '/app/chat', label: 'Chat', icon: MessageCircle },
  { to: '/app/nutrition', label: 'Nutrição', icon: Salad },
  { to: '/app/profile', label: 'Perfil', icon: User },
]

const trainerTabs = [
  { to: '/app/home', label: 'Dashboard', icon: LayoutDashboard },
  { to: '/app/agenda', label: 'Agenda', icon: Calendar },
  { to: '/app/clients', label: 'Alunos', icon: Users },
  { to: '/app/library', label: 'Biblioteca', icon: BookOpen },
  { to: '/app/chat', label: 'Chat', icon: MessageCircle },
  { to: '/app/profile', label: 'Perfil', icon: User },
]

export function AppShell() {
  const user = useAuthStore((s) => s.user)
  const setUser = useAuthStore((s) => s.setUser)
  const { theme, toggleTheme } = useThemeStore()

  useQuery({
    queryKey: ['me'],
    queryFn: async () => {
      const me = await fetchMe()
      setUser(me)
      return me
    },
  })

  const tabs = user?.role === 'trainer' ? trainerTabs : clientTabs

  return (
    <div className="flex min-h-svh flex-col">
      <header className="flex flex-shrink-0 items-center justify-between border-b border-border px-4 py-3">
        <div className="flex items-center gap-2">
          <span className="font-heading text-sm font-extrabold">Smart Kiki</span>
        </div>
        <div className="flex items-center gap-2">
          <button
            type="button"
            aria-label="Notificações"
            className="flex size-8 items-center justify-center rounded-full border border-border bg-card"
          >
            <Bell size={15} />
          </button>
          <button
            type="button"
            onClick={toggleTheme}
            aria-label="Alternar tema"
            className="flex size-8 items-center justify-center rounded-full border border-border bg-card"
          >
            {theme === 'dark' ? <Moon size={15} /> : <Sun size={15} />}
          </button>
        </div>
      </header>

      <main className="flex flex-1 flex-col overflow-y-auto">
        <Outlet />
      </main>

      <nav className="flex flex-shrink-0 gap-1 overflow-x-auto border-t border-border bg-card px-2 py-1.5">
        {tabs.map((tab) => (
          <NavLink
            key={tab.to}
            to={tab.to}
            className={({ isActive }) =>
              cn(
                'flex min-w-16 flex-1 flex-col items-center gap-0.5 rounded-lg px-1 py-1.5 text-[11px] font-semibold',
                isActive ? 'text-primary' : 'text-muted-foreground',
              )
            }
          >
            <tab.icon size={18} />
            {tab.label}
          </NavLink>
        ))}
      </nav>
    </div>
  )
}
