import { useState } from "react";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import type { JobFilters } from "@/models/jobs";
import { EmploymentType } from "@/models/jobs"
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"

type Props = {
  jobFilters: JobFilters
  onChange: (filters: JobFilters) => void
}

const employmentTypeOptions: Record<EmploymentType, string> = {
  [EmploymentType.FULL_TIME]: "Tiempo completo",
  [EmploymentType.PART_TIME]: "Medio tiempo",
  [EmploymentType.CONTRACT]: "Contrato",
  [EmploymentType.TEMPORARY]: "Temporal",
  [EmploymentType.INTERN]: "Pasantía",
  [EmploymentType.VOLUNTEER]: "Voluntariado",
  [EmploymentType.OTHER]: "Otro",
}

export default function JobsSearchFilter({ jobFilters, onChange }: Props) {
  const [filters, setFilters] = useState<JobFilters>(jobFilters)

  return (
    <div className="max-w-7xl w-full flex mx-auto rounded-lg bg-white p-4 border border-[#E5E7EB] gap-4">
      <Select>
        <SelectTrigger className="w-45">
          <SelectValue placeholder="Tipo de Empleo" />
        </SelectTrigger>
        <SelectContent>
          <SelectGroup>
            {Object.keys(EmploymentType).map(key => (
              <SelectItem
                key={key}
                value={EmploymentType[key as keyof typeof EmploymentType]}
                onSelect={() => setFilters({ ...filters, employmentType: EmploymentType[key as keyof typeof EmploymentType] })}
              >
                {employmentTypeOptions[EmploymentType[key as keyof typeof EmploymentType]]}
              </SelectItem>
            ))}
          </SelectGroup>
        </SelectContent>
      </Select>
      <Input placeholder="Salario Minimo" type="number" value={filters.minSalary || ""} onChange={e => setFilters({ ...jobFilters, minSalary: e.target.value ? parseInt(e.target.value) : undefined })} />
      <Input placeholder="Salario Maximo" type="number" value={filters.maxSalary || ""} onChange={e => setFilters({ ...jobFilters, maxSalary: e.target.value ? parseInt(e.target.value) : undefined })} />
      <Button onClick={() => onChange(filters)} className="cursor-pointer bg-linear-to-r from-[#F97316] to-[#8B5CF6] text-white hover:from-[#EA580C] hover:to-[#7C3AED] shadow-lg transition-all duration-200">
        Aplicar Filtros
      </Button>
    </div>
  )
}

