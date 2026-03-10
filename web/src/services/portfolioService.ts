import { Service } from "./service"
import type { APIResponse } from "@/models/APIResponse"
import type { PortfolioFile, UploadPortfolioResponse } from "@/models/portfolio"

class PortfolioService extends Service {
  
  async getMyPortfolio(): Promise<Array<PortfolioFile>> {
    const response = await this.client.get("me").json<APIResponse<Array<PortfolioFile>>>()
    return response.data
  }

  async getUserPortfolio(userID: string): Promise<Array<PortfolioFile>> {
        const response = await this.client.get(`${userID}`).json<APIResponse<Array<PortfolioFile>>>()
        return response.data
    }

  async uploadFile(file: File): Promise<APIResponse<UploadPortfolioResponse>> {
    const formData = new FormData()
    formData.append("file", file)

    console.log("Contenido del FormData:");
    for (const [key, value] of formData.entries()) {
        console.log(`${key}:`, value); 
    }

    const response = await this.client.post("", { 
        body: formData,
        timeout: false,
        onDownloadProgress: (progress) => {
            console.log(`Progreso: ${Math.round(progress.percent * 100)}%`);
        }
    }).json<APIResponse<UploadPortfolioResponse>>()

    return response
  }

  async downloadFile(fileID: string, fileName: string): Promise<void> {
    const blob = await this.client.get(`files/${fileID}/download`).blob()
    
    const url = window.URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = fileName
    document.body.appendChild(a)
    a.click()
    
    window.URL.revokeObjectURL(url)
    document.body.removeChild(a)
  }

  async deleteFile(fileID: string): Promise<void> {
    await this.client.delete(`${fileID}`).json()
  }
}

export const portfolioService = new PortfolioService("/api/portfolio")