import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { useQuery } from '@tanstack/react-query'
import { jobsService } from '@/services/jobsService'
import { Button } from '@/components/ui/button'
import JobDetailView from '@/components/jobs/JobDetailView'

export const Route = createFileRoute('/(private)/empleos/$empleoId/aplicaciones')({
  component: RouteComponent,
})

function RouteComponent() {
  const { empleoId } = Route.useParams()
  const navigate = useNavigate()

  const {
    data: jobResponse,
    isLoading: isJobLoading,
    isError: isJobError,
  } = useQuery({
    queryKey: ['job', empleoId],
    queryFn: () => jobsService.getJobById(empleoId),
    enabled: Boolean(empleoId),
  })
// ['job', empleoId],
// () => jobsService.getJobById(empleoId),
// { enabled: Boolean(empleoId) },

  const {
    data: applicationsResponse,
    isLoading: isApplicationsLoading,
    isError: isApplicationsError,
  } = useQuery({
    queryKey: ['jobApplications', empleoId],
    queryFn: () => jobsService.getApplicationsForJob(empleoId),
    enabled: Boolean(empleoId),
  })

  const job = jobResponse?.data
  const applications = applicationsResponse?.data ?? []

  return (
    <div className="min-h-screen bg-slate-50 p-6">
      <div className="max-w-6xl mx-auto">
        <div className="flex items-center justify-between mb-6">
          <div>
            <h1 className="text-2xl font-bold text-gray-900">Aplicaciones</h1>
            <p className="text-sm text-gray-500">Revisa las solicitudes para este puesto.</p>
          </div>
          <Button variant="secondary" onClick={() => navigate({ to: '/jobs' as any })}>
            Volver a empleos
          </Button>
        </div>

        {isJobLoading ? (
          <div className="rounded-lg bg-white p-10 text-center text-gray-500">Cargando detalles del empleo...</div>
        ) : isJobError || !job ? (
          <div className="rounded-lg bg-white p-10 text-center text-red-500">No se pudo cargar la información del empleo.</div>
        ) : (
          <div className="space-y-6">
            <div className="rounded-lg bg-white shadow-sm border border-gray-100">
              <JobDetailView job={job} />
            </div>

            <div className="rounded-lg bg-white shadow-sm border border-gray-100 p-6">
              <div className="flex items-center justify-between mb-4">
                <h2 className="text-lg font-semibold text-gray-900">Solicitudes recibidas</h2>
                <span className="text-sm text-gray-500">{applications.length} aplicac{applications.length === 1 ? 'ión' : 'iones'}</span>
              </div>

              {isApplicationsLoading ? (
                <div className="text-sm text-gray-500">Cargando aplicaciones...</div>
              ) : applications.length === 0 ? (
                <div className="text-sm text-gray-500">Aún no hay aplicaciones para este empleo.</div>
              ) : (
                <div className="space-y-4">
                  {applications.map((app) => (
                    <div key={app.id} className="border border-gray-200 rounded-lg p-4">
                      <div className="flex items-start justify-between gap-4">
                        <div>
                          <p className="font-semibold text-gray-900">{app.user.name}</p>
                          <p className="text-sm text-gray-500">{app.user.email}</p>
                          <p className="mt-2 text-sm text-gray-600 whitespace-pre-line">{app.coverLetter || 'Sin carta de presentación.'}</p>
                        </div>
                        <div className="text-right">
                          <span className="inline-flex items-center rounded-full bg-gray-100 px-3 py-1 text-xs font-medium text-gray-600">
                            {app.status}
                          </span>
                          <p className="text-xs text-gray-400 mt-1">{new Date(app.createdAt).toLocaleString()}</p>
                        </div>
                      </div>
                    </div>
                  ))}
                </div>
              )}
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
