import { useQuery } from "@tanstack/react-query"
import type { JobFilters } from "@/models/jobs";
import { jobsService } from "@/services/jobsService";

export const useJobs = (filters: JobFilters, options?: { enabled?: boolean }) => {
  return useQuery({
    queryKey: ["jobs", filters],
    enabled: options?.enabled ?? true,
    queryFn: async () => {
      return jobsService.getJobs(filters);
    },
  });
}
