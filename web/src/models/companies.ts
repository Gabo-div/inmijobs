import * as z from "zod/v4";

export const createCompanyRequestSchema = z.object({
  name: z.string().min(1, "El nombre es requerido"),
  weblink: z.string().optional(),
  linkedinUrl: z.string().optional(),
  number: z.string().optional(),
  description: z.string().optional(),
  sector: z.string().optional(),
  foundation: z.string().optional(),
  size: z.string().optional(),
  locations: z.array(z.object({
    address: z.string(),
    city: z.string(),
    country: z.string(),
    isHq: z.boolean(),
  })).optional(),
})

export type CreateCompanyRequest = z.infer<typeof createCompanyRequestSchema>;

export const companyResponseSchema = z.object({
  id: z.string(),
  name: z.string(),
  weblink: z.string(),
  linkedinUrl: z.string(),
  number: z.string(),
  description: z.string(),
  sector: z.string(),
  foundation: z.string(),
  size: z.string(),
  logo: z.string().nullable(),
  banner: z.string().nullable(),
  createdAt: z.number(),
  updatedAt: z.number(),
  userId: z.string(),
  locations: z.array(z.object({
    id: z.string(),
    address: z.string(),
    city: z.string(),
    country: z.string(),
    isHq: z.boolean(),
  })),
})

export type CompanyResponse = z.infer<typeof companyResponseSchema>;

export const updateCompanyRequestSchema = z.object({
  name: z.string().optional(),
  weblink: z.string().optional(),
  linkedinUrl: z.string().optional(),
  number: z.string().optional(),
  description: z.string().optional(),
  sector: z.string().optional(),
  foundation: z.string().optional(),
  size: z.string().optional(),
  logo: z.string().optional(),
  banner: z.string().optional(),
})

export type UpdateCompanyRequest = z.infer<typeof updateCompanyRequestSchema>;


export const updateCompanyResponseSchema = z.object({
  message: z.string(),
})

export type UpdateCompanyResponse = z.infer<typeof updateCompanyResponseSchema>;
