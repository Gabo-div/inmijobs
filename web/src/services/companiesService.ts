import { Service } from "./service";
import type { APIResponse } from "@/models/APIResponse";
import type { CompanyResponse, CreateCompanyRequest, UpdateCompanyRequest, UpdateCompanyResponse } from "@/models/companies";

type GetCompaniesFilters = {
  userId?: string
  page?: number
  limit?: number
  name?: string
}

class CompaniesService extends Service {
  createCompany(data: CreateCompanyRequest): Promise<APIResponse<CompanyResponse>> {
    return this.client.post("", { json: data }).json();
  }

  getCompanies(filters: GetCompaniesFilters = {}): Promise<APIResponse<{ data: CompanyResponse[] }>> {
    const searchParams: Record<string, any> = {}
    if (filters.userId) searchParams.userId = filters.userId
    if (filters.page) searchParams.page = filters.page
    if (filters.limit) searchParams.limit = filters.limit
    if (filters.name) searchParams.name = filters.name

    return this.client.get("", { searchParams }).json();
  }

  getCompanyById(id: string): Promise<APIResponse<CompanyResponse>> {
    return this.client.get(id).json();
  }

  updateCompanyById(id: string, data: UpdateCompanyRequest): Promise<APIResponse<UpdateCompanyResponse>> {
    return this.client.put(id, { json: data }).json();
  }
}

export const companiesService = new CompaniesService("/api/companies")
