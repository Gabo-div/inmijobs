import { Service } from "./service";
import type { APIResponse } from "@/models/APIResponse";
import type { CompanyResponse, CreateCompanyRequest, UpdateCompanyRequest, UpdateCompanyResponse } from "@/models/companies";

class CompaniesService extends Service {
  createCompany(data: CreateCompanyRequest): Promise<APIResponse<CompanyResponse>> {
    return this.client.post("/", { json: data }).json();
  }

  getCompanies(): Promise<APIResponse<CompanyResponse[]>> {
    return this.client.get("/").json();
  }

  getCompanyById(id: string): Promise<APIResponse<CompanyResponse>> {
    return this.client.get(`/${id}`).json();
  }

  updateCompanyById(id: string, data: UpdateCompanyRequest): Promise<APIResponse<UpdateCompanyResponse>> {
    return this.client.put(`/${id}`, { json: data }).json();
  }
}

export const companiesService = new CompaniesService("/api/companies")
