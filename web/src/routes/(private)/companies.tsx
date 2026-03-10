import { useMemo } from 'react'
import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useCompanyForm } from '@/hooks/useCompanyForm'

export const Route = createFileRoute('/(private)/companies')({
  component: RouteComponent,
})

function RouteComponent() {
  const navigate = useNavigate()
  const {
    values,
    setField,
    errors,
    isSubmitting,
    submissionError,
    isValid,
    handleSubmit,
  } = useCompanyForm(() => navigate({ to: '/jobs' }))

  const inputClass = useMemo(
    () =>
      'border-[#E5E7EB] focus:border-[#F97316] ' +
      'focus:ring-primary focus:ring-1',
    [],
  )

  const errorClass = 'border-destructive focus:border-destructive focus:ring-destructive/50'

  return (
    <div className="bg-linear-to-br from-[#FFF3E6] to-[#F3E8FF] h-[calc(100vh-64px)] p-8 gap-4 flex flex-col">
      <div className="max-w-2xl mx-auto w-full">
        <Card className="border-2 border-[#E5E7EB] shadow-lg">
          <CardHeader>
            <CardTitle className="text-xl font-bold text-[#1F2937]">Crear Nueva Compañía</CardTitle>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="name">Nombre de la compañía</Label>
                <Input
                  id="name"
                  value={values.name}
                  onChange={(event) => setField('name', event.target.value)}
                  placeholder="Ej: TechCorp"
                  className={`${inputClass} ${errors.name ? errorClass : ''}`}
                  aria-invalid={Boolean(errors.name)}
                />
                {errors.name && <p className="text-xs text-destructive">{errors.name}</p>}
              </div>

              <div className="space-y-2">
                <Label htmlFor="description">Descripción</Label>
                <Input
                  id="description"
                  value={values.description}
                  onChange={(event) => setField('description', event.target.value)}
                  placeholder="Breve descripción de la empresa"
                  className={`${inputClass} ${errors.description ? errorClass : ''}`}
                  aria-invalid={Boolean(errors.description)}
                />
                {errors.description && <p className="text-xs text-destructive">{errors.description}</p>}
              </div>

              <div className="space-y-2">
                <Label htmlFor="industry">Industria</Label>
                <Input
                  id="industry"
                  value={values.industry}
                  onChange={(event) => setField('industry', event.target.value)}
                  placeholder="Ej: Tecnología, Salud, Educación..."
                  className={`${inputClass} ${errors.industry ? errorClass : ''}`}
                  aria-invalid={Boolean(errors.industry)}
                />
                {errors.industry && <p className="text-xs text-destructive">{errors.industry}</p>}
              </div>

              <div className="space-y-2">
                <Label htmlFor="website">Sitio web</Label>
                <Input
                  id="website"
                  type="url"
                  value={values.website}
                  onChange={(event) => setField('website', event.target.value)}
                  placeholder="https://ejemplo.com"
                  className={`${inputClass} ${errors.website ? errorClass : ''}`}
                  aria-invalid={Boolean(errors.website)}
                />
                {errors.website && <p className="text-xs text-destructive">{errors.website}</p>}
              </div>

              <div className="space-y-2">
                <Label htmlFor="location">Ubicación</Label>
                <Input
                  id="location"
                  value={values.location}
                  onChange={(event) => setField('location', event.target.value)}
                  placeholder="Ciudad, País"
                  className={`${inputClass} ${errors.location ? errorClass : ''}`}
                  aria-invalid={Boolean(errors.location)}
                />
                {errors.location && <p className="text-xs text-destructive">{errors.location}</p>}
              </div>

              {submissionError && (
                <div className="rounded-md bg-destructive/10 px-4 py-3 text-sm text-destructive">
                  {submissionError}
                </div>
              )}

              <Button
                type="submit"
                className="w-full bg-[#F97316] hover:bg-[#EA580C] text-white font-semibold"
                disabled={isSubmitting || !isValid}
              >
                {isSubmitting ? 'Creando...' : 'Crear Compañía'}
              </Button>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
