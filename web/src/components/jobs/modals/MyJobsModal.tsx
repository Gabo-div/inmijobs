import { useEffect } from 'react'
import { useNavigate } from '@tanstack/react-router'
import { X } from 'lucide-react'
import { authClient } from '@/lib/auth'
import { useJobs } from '@/hooks/jobs'
import { JobCard } from '@/components/jobs/JobCard'
import { Button } from '@/components/ui/button'

interface Props {
  isOpen: boolean
  onClose: () => void
}

export const MyJobsModal = ({ isOpen, onClose }: Props) => {
  const navigate = useNavigate()
  const { data: session } = authClient.useSession()
  const userId = session?.user.id

  const { data, isLoading, isError } = useJobs(
    { userId: userId ?? '', page: 1, pageSize: 20 },
    { enabled: Boolean(userId) },
  )

  useEffect(() => {
    if (!isOpen) return
    document.body.style.overflow = 'hidden'
    return () => {
      document.body.style.overflow = 'unset'
    }
  }, [isOpen])

  if (!isOpen) return null

  const jobs = data?.data.jobs ?? []

  const handleNavigateToApplications = (jobId: string) => {
    onClose()
    navigate({ to: '/empleos/$empleoId/aplicaciones' as any, params: { empleoId: jobId } as any })
  }

  return (
    <div className="fixed inset-0 bg-black/60 z-50 flex items-center justify-center p-4 animate-fade-in">
      <div className="bg-white rounded-xl shadow-2xl w-full max-w-4xl mx-auto relative flex flex-col max-h-[90vh] overflow-hidden animate-scale-in">
        <div className="flex items-center justify-between p-4 border-b border-gray-200">
          <div>
            <h2 className="text-lg font-semibold text-gray-900">Mis publicaciones</h2>
            <p className="text-xs text-gray-500">Ver y administrar tus empleos publicados.</p>
          </div>
          <button
            onClick={onClose}
            className="p-2 bg-gray-100 hover:bg-gray-200 rounded-full text-gray-500 transition"
            aria-label="Cerrar modal"
          >
            <X size={18} />
          </button>
        </div>

        <div className="p-4 flex-1 overflow-y-auto">
          {isLoading ? (
            <p className="text-center text-sm text-gray-500">Cargando tus empleos...</p>
          ) : isError ? (
            <p className="text-center text-sm text-red-500">Error al cargar tus publicaciones. Intenta de nuevo.</p>
          ) : jobs.length === 0 ? (
            <p className="text-center text-sm text-gray-500">No has publicado ningún empleo aún.</p>
          ) : (
            <div className="space-y-2">
              {jobs.map((job) => (
                <button
                  key={job.id}
                  onClick={() => handleNavigateToApplications(job.id)}
                  className="w-full text-left"
                >
                  <JobCard job={job} onSelect={() => handleNavigateToApplications(job.id)} selected={false} />
                </button>
              ))}
            </div>
          )}
        </div>

        <div className="p-4 border-t border-gray-200 flex justify-end">
          <Button variant="secondary" onClick={onClose}>
            Cerrar
          </Button>
        </div>
      </div>
    </div>
  )
}
