import { useEffect, useMemo, useState } from 'react'
import { X } from 'lucide-react'
import { useQueryClient } from '@tanstack/react-query'
import { jobsService } from '@/services/jobsService'
import type { CompanyResponse } from '@/models/companies'
import type { CreateJobRequest } from '@/models/jobs'
import { Button } from '@/components/ui/button'

interface Props {
  isOpen: boolean
  onClose: () => void
  companies: CompanyResponse[] | undefined
}

const EMPLOYMENT_TYPES = [
  { value: 'full-time', label: 'Tiempo completo' },
  { value: 'part-time', label: 'Medio tiempo' },
  { value: 'contract', label: 'Contrato' },
  { value: 'temporary', label: 'Temporal' },
  { value: 'intern', label: 'Pasantía' },
  { value: 'volunteer', label: 'Voluntariado' },
  { value: 'other', label: 'Otro' },
]

export const CreateJobModal = ({ isOpen, onClose, companies }: Props) => {
  const queryClient = useQueryClient()
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [location, setLocation] = useState('')
  const [employmentType, setEmploymentType] = useState(EMPLOYMENT_TYPES[0].value)
  const [salaryMin, setSalaryMin] = useState<number | undefined>(undefined)
  const [salaryMax, setSalaryMax] = useState<number | undefined>(undefined)
  const [companyId, setCompanyId] = useState<string | undefined>(companies?.[0]?.id)
  const [isSaving, setIsSaving] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const hasCompanies = Boolean(companies && companies.length > 0)

  useEffect(() => {
    if (isOpen) {
      document.body.style.overflow = 'hidden'
      setError(null)
      setTitle('')
      setDescription('')
      setLocation('')
      setSalaryMin(undefined)
      setSalaryMax(undefined)
      setEmploymentType(EMPLOYMENT_TYPES[0].value)
      setCompanyId(companies?.[0]?.id)
    } else {
      document.body.style.overflow = 'unset'
    }
    return () => {
      document.body.style.overflow = 'unset'
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isOpen])

  useEffect(() => {
    if (companies && companies.length > 0) {
      setCompanyId(companies[0].id)
    }
  }, [companies])

  const canPublish = useMemo(() => {
    return (
      hasCompanies &&
      title.trim().length > 0 &&
      description.trim().length > 0 &&
      location.trim().length > 0 &&
      !!companyId
    )
  }, [hasCompanies, title, description, location, companyId])

  const handleSubmit = async () => {
    if (!canPublish || !companyId) return
    setIsSaving(true)
    setError(null)

    const payload: CreateJobRequest = {
      title: title.trim(),
      description: description.trim(),
      location: location.trim(),
      company_id: companyId,
      employment_type: employmentType,
      salary_min: salaryMin,
      salary_max: salaryMax,
    }

    try {
      await jobsService.createJob(payload)
      await queryClient.invalidateQueries({ queryKey: ['jobs'] })
      onClose()
    } catch (err) {
      setError('No se pudo crear el empleo. Por favor intenta nuevamente.')
      console.error(err)
    } finally {
      setIsSaving(false)
    }
  }

  if (!isOpen) return null

  return (
    <div className="fixed inset-0 bg-black/60 z-50 flex items-center justify-center p-4 animate-fade-in">
      <div className="bg-white rounded-xl shadow-2xl w-full max-w-2xl mx-auto relative flex flex-col max-h-[90vh] overflow-hidden animate-scale-in">
        <div className="flex items-center justify-between p-4 border-b border-gray-200">
          <div>
            <h2 className="text-lg font-semibold text-gray-900">Publicar Empleo</h2>
            <p className="text-xs text-gray-500">Completa los datos y publica un nuevo puesto.</p>
          </div>
          <button
            onClick={onClose}
            className="p-2 bg-gray-100 hover:bg-gray-200 rounded-full text-gray-500 transition"
            aria-label="Cerrar modal"
          >
            <X size={18} />
          </button>
        </div>

        <div className="p-4 flex-1 overflow-y-auto space-y-4">
          {!hasCompanies ? (
            <div className="rounded-lg border border-yellow-200 bg-yellow-50 p-4">
              <p className="text-sm text-yellow-800">
                Para publicar un empleo necesitas una empresa asociada. Crea tu empresa en la sección de perfil.
              </p>
            </div>
          ) : (
            <>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <label className="flex flex-col text-sm font-medium text-gray-600">
                  Título del puesto
                  <input
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    className="mt-1 rounded-md border border-gray-200 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary"
                    placeholder="Ej. Desarrollador Frontend"
                  />
                </label>

                <label className="flex flex-col text-sm font-medium text-gray-600">
                  Ubicación
                  <input
                    value={location}
                    onChange={(e) => setLocation(e.target.value)}
                    className="mt-1 rounded-md border border-gray-200 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary"
                    placeholder="Ej. Remoto / Ciudad"
                  />
                </label>

                <label className="flex flex-col text-sm font-medium text-gray-600">
                  Tipo de empleo
                  <select
                    value={employmentType}
                    onChange={(e) => setEmploymentType(e.target.value)}
                    className="mt-1 rounded-md border border-gray-200 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary"
                  >
                    {EMPLOYMENT_TYPES.map((opt) => (
                      <option key={opt.value} value={opt.value}>
                        {opt.label}
                      </option>
                    ))}
                  </select>
                </label>

                <label className="flex flex-col text-sm font-medium text-gray-600">
                  Empresa
                  <select
                    value={companyId}
                    onChange={(e) => setCompanyId(e.target.value)}
                    className="mt-1 rounded-md border border-gray-200 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary"
                  >
                    {companies?.map((company) => (
                      <option key={company.id} value={company.id}>
                        {company.name}
                      </option>
                    ))}
                  </select>
                </label>

                <div className="flex flex-col gap-2">
                  <label className="flex flex-col text-sm font-medium text-gray-600">
                    Salario mínimo
                    <input
                      type="number"
                      min={0}
                      value={salaryMin ?? ''}
                      onChange={(e) => setSalaryMin(e.target.value ? Number(e.target.value) : undefined)}
                      className="mt-1 rounded-md border border-gray-200 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary"
                      placeholder="Ej. 500"
                    />
                  </label>

                  <label className="flex flex-col text-sm font-medium text-gray-600">
                    Salario máximo
                    <input
                      type="number"
                      min={0}
                      value={salaryMax ?? ''}
                      onChange={(e) => setSalaryMax(e.target.value ? Number(e.target.value) : undefined)}
                      className="mt-1 rounded-md border border-gray-200 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary"
                      placeholder="Ej. 1000"
                    />
                  </label>
                </div>
              </div>

              <label className="flex flex-col text-sm font-medium text-gray-600">
                Descripción
                <textarea
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                  className="mt-1 min-h-[150px] resize-none rounded-md border border-gray-200 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary"
                  placeholder="Descripción del puesto, responsabilidades y requisitos"
                />
              </label>

              {error && (
                <p className="text-sm text-red-600">{error}</p>
              )}
            </>
          )}
        </div>

        <div className="p-4 border-t border-gray-200 flex items-center justify-end gap-2">
          <Button variant="secondary" onClick={onClose} disabled={isSaving}>
            Cancelar
          </Button>
          <Button
            onClick={handleSubmit}
            disabled={!canPublish || isSaving}
          >
            {isSaving ? 'Publicando...' : 'Publicar empleo'}
          </Button>
        </div>
      </div>
    </div>
  )
}
