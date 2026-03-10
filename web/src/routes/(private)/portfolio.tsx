import React, { useEffect, useState } from 'react'
import { createFileRoute } from '@tanstack/react-router'
import { 
    File as FileIcon, 
    Plus,
    Upload,
} from 'lucide-react'
import { FileCard } from '../../components/portfolio/fileCard'
import type { PortfolioFile } from "@/models/portfolio"
import { portfolioService } from "@/services/portfolioService"

export const Route = createFileRoute('/(private)/portfolio')({
    component: PortfolioPage,
})

export interface FileData {
  id: string
  original_name: string
  file_type: string
  file_size: number
  download_url: string
  created_at: number
}

function PortfolioPage() {
    const [files, setFiles] = useState<Array<PortfolioFile>>([])
    const [isLoading, setIsLoading] = useState(true)
    const fileInputRef = React.useRef<HTMLInputElement>(null)
    const [isUploading, setIsUploading] = useState(false)
    
    const fetchPortfolio = async () => {
        setIsLoading(true)
        try {
            const data = await portfolioService.getMyPortfolio()
            setFiles(data)

        } catch(error) {
            console.error("Error de conexión:", error)
        }
        setIsLoading(false)
    }

    useEffect(() => {
        fetchPortfolio()
    }, [])

    const handleFileUpload = async (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0]
        if (!file) return

        if(file.size > 20 * 1024 * 1024){
            alert("El archivo que intenta subir es demasiado grande.")
            return
        }

        setIsUploading(true)

        try {
            await portfolioService.uploadFile(file)
            fetchPortfolio()
        } catch(error) {
            console.error("Error en la petición:", error)
            alert("Error de conexión con el servidor")
        }
        setIsUploading(false)
        if (fileInputRef.current) fileInputRef.current.value = ''
    }

    const handleDelete = async (id: string) => {
        const confirmDelete = window.confirm("¿Estás seguro de que deseas eliminar este archivo?")
        if (!confirmDelete) return

        try {
            await portfolioService.deleteFile(id)
            setFiles(prevFiles => prevFiles.filter(file => file.id !== id))
            
        } catch (error) {
            console.error("Error al eliminar:", error);
            alert("No tienes permiso para eliminar este archivo o ya no existe.")
        }
    }

  return (
    <div className="bg-linear-to-br from-[#FFF3E6] to-[#F3E8FF] h-[calc(100vh-64px)] p-8 min-h-screen">
        <div className="flex justify-between items-center mb-8">
            <div>
                <h1 className="text-2xl font-bold text-gray-800">Mi Portfolio</h1>
                <p className="text-gray-500">Sube y visualiza tus archivos subidos</p>
            </div>
            
            <div className="relative">
                <button 
                    onClick={() => fileInputRef.current?.click()}
                    disabled={isUploading}
                    className="relative z-50 flex items-center gap-2 px-5 py-2.5 rounded-xl bg-linear-to-r from-[#F97316] to-[#8B5CF6] text-white hover:shadow-xl transition-all active:scale-95"
                >
                    <Upload size={18} />
                    {isUploading ? 'Subiendo...' : 'Subir Archivo'}
                </button>
            </div>
        </div>

        { isLoading ? (
            <div className="flex justify-center items-center h-64">
                <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-[#8B5CF6]"></div>
                <p className="ml-3 text-gray-500 font-medium">Cargando archivos...</p>
            </div>
        ) : files.length === 0 ? (
            <div className="flex flex-col items-center justify-center h-64 border-2 border-dashed border-gray-200 rounded-3xl bg-white/50">
                <FileIcon size={48} className="text-gray-300 mb-4" />
                <p className="text-gray-500 font-medium">No tienes archivos en tu portfolio</p>
                <button 
                    onClick={() => fileInputRef.current?.click()}
                    className="mt-4 text-[#8B5CF6] font-bold hover:underline"
                >
                    Sube tu primer archivo
                </button>
            </div>
        ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
                {files.map((file) => (
                    <FileCard 
                        key={file.id} 
                        file={file} 
                        onDelete={handleDelete} 
                    />
                ))}

                <div 
                    onClick={() => fileInputRef.current?.click()}
                    className="border-2 border-dashed border-gray-200 rounded-2xl flex flex-col items-center justify-center p-8 hover:border-black/20 hover:bg-gray-100/50 transition-all cursor-pointer text-gray-400 group"
                >
                    <div className="p-3 rounded-full bg-gray-100 group-hover:bg-white mb-2">
                        <Plus size={24} />
                    </div>
                    <span className="text-sm font-medium group-hover:text-[#8B5CF6]">Agregar nuevo</span>
                </div>
            </div>
        )}

        <input 
            type="file" 
            ref={fileInputRef} 
            onChange={handleFileUpload} 
            className="hidden" 
            accept=".pdf,.doc,.docx,.jpg,.jpeg,.png,.gif,.webp,.xlsx,.xls,.pptx,.ppt"
        />
    </div>
  )
}
