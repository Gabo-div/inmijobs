import { useEffect, useState } from 'react'
import { createFileRoute } from '@tanstack/react-router'
import type { PortfolioFile } from "@/models/portfolio"
import { portfolioService } from "@/services/portfolioService"
import { FileCard } from '@/components/portfolio/fileCard'

export const Route = createFileRoute('/(private)/profile/$userId/portfolio')({
  component: UserPortfolio,
})

function UserPortfolio() {
    const { userId } = Route.useParams()
    const [files, setFiles] = useState<Array<PortfolioFile>>([])
    const [isLoading, setIsLoading] = useState(true)

    useEffect(() => {
        const fetchUserPortfolio = async () => {
            try {
                const data = await portfolioService.getUserPortfolio(userId)
                setFiles(data)
            } catch (error) {
                console.error("Error cargando portfolio:", error)
            } finally {
                setIsLoading(false)
            }
        }
        fetchUserPortfolio()
    }, [userId])

    return (
        <div className="p-8 bg-gray-50 min-h-screen">
            <h1 className="text-2xl font-bold mb-6 text-gray-800">Portfolio de Usuario</h1>
            
            {isLoading ? (
                <div>Cargando...</div>
            ) : (
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
                    {files.map(file => (
                        <FileCard
                            key={file.id} 
                            file={file} 
                            onDelete={undefined} 
                        />
                    ))}
                </div>
            )}
        </div>
    )
}
