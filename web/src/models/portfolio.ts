import * as z from "zod/v4";

export const portfolioFileSchema = z.object({
  id: z.string(),
  original_name: z.string(),
  file_type: z.string(),
  file_size: z.number(),
  download_url: z.string(),
  created_at: z.number(),
});

export type PortfolioFile = z.infer<typeof portfolioFileSchema>;

export const uploadPortfolioResponseSchema = z.object({
  message: z.string(),
  file: portfolioFileSchema,
});

export type UploadPortfolioResponse = z.infer<typeof uploadPortfolioResponseSchema>;