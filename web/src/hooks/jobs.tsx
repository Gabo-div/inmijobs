import { useQuery } from "@tanstack/react-query"
import type { JobFilters } from "@/models/jobs";
import { jobsService } from "@/services/jobsService";

export const useJobs = (filters: JobFilters) => {
  return useQuery({
    queryKey: ["jobs", filters],
    queryFn: async () => {
      return jobsService.getJobs(filters);
    },
  });
}
