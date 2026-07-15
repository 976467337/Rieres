import { api } from '@/lib/api'
import type { ProgressPhoto } from '@/types/progress-photo'

export async function uploadProgressPhoto(studentId: string, file: File) {
  const form = new FormData()
  form.append('student_id', studentId)
  form.append('photo', file)

  const { data } = await api.post<ProgressPhoto>('/progress-photos', form, {
    headers: { 'Content-Type': 'multipart/form-data' },
  })
  return data
}

export async function listProgressPhotos(studentId?: string) {
  const { data } = await api.get<ProgressPhoto[]>('/progress-photos', {
    params: studentId ? { student_id: studentId } : undefined,
  })
  return data
}
