import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { LoginPage } from '@/pages/LoginPage'
import { SignupPage } from '@/pages/SignupPage'
import { AppShell } from '@/components/AppShell'
import { ProtectedRoute } from '@/components/ProtectedRoute'
import { ComingSoon } from '@/components/ComingSoon'
import { RoleHomePage } from '@/pages/app/RoleHomePage'
import { StudentsPage } from '@/pages/trainer/StudentsPage'
import { PlansPage } from '@/pages/trainer/PlansPage'

const queryClient = new QueryClient()

export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Navigate to="/app/home" replace />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/signup" element={<SignupPage />} />

          <Route element={<ProtectedRoute />}>
            <Route path="/app" element={<AppShell />}>
              <Route path="home" element={<RoleHomePage />} />
              <Route path="agenda" element={<ComingSoon title="Agenda" />} />
              <Route path="marketplace" element={<ComingSoon title="Buscar personal" />} />
              <Route path="metrics" element={<ComingSoon title="Evolução" />} />
              <Route path="challenges" element={<ComingSoon title="Desafios" />} />
              <Route path="chat" element={<ComingSoon title="Chat" />} />
              <Route path="nutrition" element={<ComingSoon title="Nutrição" />} />
              <Route path="clients" element={<StudentsPage />} />
              <Route path="plans" element={<PlansPage />} />
              <Route path="library" element={<ComingSoon title="Biblioteca" />} />
              <Route path="profile" element={<ComingSoon title="Perfil" />} />
            </Route>
          </Route>

          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </BrowserRouter>
    </QueryClientProvider>
  )
}
