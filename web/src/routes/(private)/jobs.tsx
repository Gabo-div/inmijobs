import { createFileRoute } from '@tanstack/react-router'
import { useEffect, useState } from 'react';
import type { JobFilters } from '@/models/jobs';
import { JobCard } from '@/components/jobs/JobCard';
import JobDetailView from '@/components/jobs/JobDetailView';
import { useJobs } from '@/hooks/jobs';
import { ScrollArea } from '@/components/ui/scroll-area';
import JobsSearchFilter from '@/components/jobs/JobsSearchFilter';
import { Button } from '@/components/ui/button';

export const Route = createFileRoute('/(private)/jobs')({
  component: RouteComponent,
})

function RouteComponent() {
  const [filters, setFilters] = useState<JobFilters>({})
  const [page, setPage] = useState(1);
  const { data, isLoading, isError } = useJobs({ ...filters, page, pageSize: 10 })

  const [selectedId, setSelectedId] = useState<string | null>(null);
  const activeJob = data?.data.jobs.find(job => job.id === selectedId);

  useEffect(() => {
    if (data?.data.jobs.length) {
      setSelectedId(data.data.jobs[0]?.id || null);
    }
  }, [data])

  return (
    <div className="bg-linear-to-br from-[#FFF3E6] to-[#F3E8FF] h-[calc(100vh-64px)] p-8 gap-4 flex flex-col">
      <JobsSearchFilter jobFilters={filters} onChange={(newFiltes) => {
        setFilters(newFiltes);
        setPage(1);
      }}
      />

      <div className="max-w-7xl flex w-full flex-1 mx-auto rounded-lg border border-[#E5E7EB] bg-white overflow-hidden">        {data?.data.jobs.length ? (<>
        <div className="h-[calc(100%-64px)]">
          <ScrollArea className="h-full border-r border-[#E5E7EB]">
            <div className="flex flex-col gap-px bg-[#E5E7EB]">
              {data.data.jobs.map(job => (
                <JobCard
                  key={job.id}
                  job={job}
                  onSelect={setSelectedId}
                  selected={job.id === selectedId}
                />
              ))}
            </div>
          </ScrollArea>
          <div className="h-16 border-t border-[#E5E7EB] flex items-center p-2 gap-2">
            <Button className="flex-1 cursor-pointer" variant="secondary" disabled={page === 1} onClick={() => setPage(prev => Math.max(prev - 1, 1))}>
              Anterior
            </Button>
            <div className="flex-2 text-center text-sm text-gray-500">
              Página {page} de {data.data.totalPages}
            </div>
            <Button className="flex-1 cursor-pointer" variant="secondary" disabled={page === data.data.totalPages} onClick={() => setPage(prev => Math.min(prev + 1, data.data.totalPages))}>
              Siguiente
            </Button>
          </div>
        </div>
        <JobDetailView job={activeJob} />
      </>
      ) :
        <div className="flex-1 bg-white flex items-center justify-center">
          <p className="text-gray-500">
            {isLoading ? 'Loading...' : isError ? 'Error loading jobs' : 'No jobs found'}
          </p>
        </div>
      }
      </div>
    </div>)
}
