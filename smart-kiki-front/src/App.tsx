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
import { LibraryPage } from '@/pages/trainer/LibraryPage'
import { StudentDetailPage } from '@/pages/trainer/StudentDetailPage'
import { WorkoutBuilderPage } from '@/pages/trainer/WorkoutBuilderPage'
import { WorkoutDetailPage } from '@/pages/client/WorkoutDetailPage'
import { RoleAgendaPage } from '@/pages/app/RoleAgendaPage'
import { RoleChatListPage } from '@/pages/app/RoleChatListPage'
import { ChatConversationPage } from '@/pages/app/ChatConversationPage'
import { MarketplacePage } from '@/pages/client/MarketplacePage'
import { RoleProfilePage } from '@/pages/app/RoleProfilePage'
import { NutritionPage } from '@/pages/client/NutritionPage'
import { MetricsPage } from '@/pages/client/MetricsPage'

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
              <Route path="agenda" element={<RoleAgendaPage />} />
              <Route path="marketplace" element={<MarketplacePage />} />
              <Route path="metrics" element={<MetricsPage />} />
              <Route path="challenges" element={<ComingSoon title="Desafios" />} />
              <Route path="chat" element={<RoleChatListPage />} />
              <Route path="chat/:userId" element={<ChatConversationPage />} />
              <Route path="nutrition" element={<NutritionPage />} />
              <Route path="clients" element={<StudentsPage />} />
              <Route path="students/:studentId" element={<StudentDetailPage />} />
              <Route path="students/:studentId/workouts/new" element={<WorkoutBuilderPage />} />
              <Route path="workouts/:id" element={<WorkoutDetailPage />} />
              <Route path="plans" element={<PlansPage />} />
              <Route path="library" element={<LibraryPage />} />
              <Route path="profile" element={<RoleProfilePage />} />
            </Route>
          </Route>

          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </BrowserRouter>
    </QueryClientProvider>
  )
}
