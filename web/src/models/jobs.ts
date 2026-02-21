import * as z from "zod/v4";

export const jobSchema = z.object({
  id: z.string(),
  title: z.string(),
  description: z.string(),
  location: z.string(),
  salaryMin: z.number().optional(),
  salaryMax: z.number().optional(),
  employmentType: z.string(),
  isActive: z.boolean(),
  createdAt: z.coerce.date(),
  updatedAt: z.coerce.date(),
});

export type Job = z.infer<typeof jobSchema>;

export const updateJobRequestSchema = z.object({
  id: z.string(),
  title: z.string(),
  description: z.string(),
  location: z.string(),
  salaryMin: z.number().optional(),
  salaryMax: z.number().optional(),
  employmentType: z.string(),
  isActive: z.boolean(),
});

export type UpdateJobRequest = z.infer<typeof updateJobRequestSchema>;

export enum EmploymentType {
  FULL_TIME = "full-time",
  PART_TIME = "part-time",
  CONTRACT = "contract",
  TEMPORARY = "temporary",
  INTERN = "intern",
  VOLUNTEER = "volunteer",
  OTHER = "other",
}

export const jobFiltersSchema = z.object({
  page: z.coerce.number(),
  pageSize: z.coerce.number(),
  location: z.string().optional(),
  employmentType: z.enum(EmploymentType).optional(),
  minSalary: z.number().optional(),
  maxSalary: z.number().optional(),
  sector: z.string().optional(),
}).partial();

export type JobFilters = z.infer<typeof jobFiltersSchema>;

export const paginatedJobsSchema = z.object({
  jobs: z.array(jobSchema),
  total: z.coerce.number(),
  page: z.coerce.number(),
  limit: z.coerce.number(),
  totalPages: z.coerce.number(),
});

export type PaginatedJobs = z.infer<typeof paginatedJobsSchema>;
