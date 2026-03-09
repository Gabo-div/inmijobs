import { useMemo, useState } from 'react'
import { useMutation } from '@tanstack/react-query'
import { companiesService } from '@/services/companiesService'
import type { CreateCompanyRequest } from '@/models/companies'

export type CompanyFormValues = {
  name: string
  description: string
  industry: string
  website: string
  location: string
}

export type CompanyFormErrors = Partial<Record<keyof CompanyFormValues, string>>

export function useCompanyForm(onSuccess?: () => void) {
  const [values, setValues] = useState<CompanyFormValues>({
    name: '',
    description: '',
    industry: '',
    website: '',
    location: '',
  })

  const [errors, setErrors] = useState<CompanyFormErrors>({})
  const [submissionError, setSubmissionError] = useState<string | null>(null)

  const mutation = useMutation({
    mutationFn: (payload: CreateCompanyRequest) => companiesService.createCompany(payload),
    onSuccess: () => {
      onSuccess?.()
    },
    onError: (error) => {
      setSubmissionError(
        (error as any)?.message ||
          'No se pudo crear la compañía. Intenta de nuevo más tarde.',
      )
    },
  })

  const validate = (values: CompanyFormValues): CompanyFormErrors => {
    const nextErrors: CompanyFormErrors = {}

    if (!values.name.trim()) {
      nextErrors.name = 'El nombre de la compañía es obligatorio.'
    }

    if (!values.description.trim()) {
      nextErrors.description = 'La descripción es obligatoria.'
    }

    if (!values.industry.trim()) {
      nextErrors.industry = 'La industria es obligatoria.'
    }

    if (!values.website.trim()) {
      nextErrors.website = 'El sitio web es obligatorio.'
    } else {
      try {
        new URL(values.website)
      } catch {
        nextErrors.website = 'Ingresa una URL válida.'
      }
    }

    if (!values.location.trim()) {
      nextErrors.location = 'La ubicación es obligatoria.'
    }

    return nextErrors
  }

  const isSubmitting = mutation.isLoading

  const isValid = useMemo(() => {
    const validation = validate(values)
    return Object.keys(validation).length === 0
  }, [values])

  const setField = (field: keyof CompanyFormValues, value: string) => {
    setValues((prev) => ({ ...prev, [field]: value }))
    setErrors((prev) => ({ ...prev, [field]: undefined }))
  }

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    setSubmissionError(null)

    const validationErrors = validate(values)
    setErrors(validationErrors)

    if (Object.keys(validationErrors).length > 0) {
      return
    }
console.log('Submitting company with values:', values)
    mutation.mutate({
      name: values.name.trim(),
      description: values.description.trim(),
      sector: values.industry.trim(),
      weblink: values.website.trim(),
      // location is not part of the API model, but keep it around for local validation.
    })
  }

  return {
    values,
    setField,
    errors,
    isSubmitting,
    submissionError,
    isValid,
    handleSubmit,
  }
}
