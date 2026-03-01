import React, { useState } from 'react'
import { createFileRoute } from '@tanstack/react-router'
import { 
  ChevronDown,
  File as FileIcon, 
  FileImage, 
  FileText, 
  FileVideo, 
  Headphones,
  Plus,
  Upload,
} from 'lucide-react'
import { FileCard } from '../../components/portfolio/fileCard'

export const Route = createFileRoute('/(private)/portfolio')({
  component: PortfolioPage,
})

const MOCK_FILES = [
  { id: 1, name: 'Curriculum.pdf', type: 'application/pdf', size: '1.2 MB', date: '2024-02-15' },
  { id: 2, name: 'Proyecto Final.zip', type: 'application/zip', size: '45 MB', date: '2024-02-10' },
  { id: 3, name: 'Diseño.png', type: 'image/png', size: '2.5 MB', date: '2024-02-18' },
  { id: 4, name: 'Video Introductorio.mp4', type: 'video/mp4', size: '120 MB', date: '2024-01-05' },
]

function PortfolioPage() {
  const [files] = useState(MOCK_FILES)
  const [isUploadMenuOpen, setIsUploadMenuOpen] = useState(false)

  const uploadOptions = [
    { label: 'Imagen', icon: FileImage, color: 'text-blue-500' },
    { label: 'Video', icon: FileVideo, color: 'text-purple-500' },
    { label: 'Audio', icon: Headphones, color: 'text-pink-500' },
    { label: 'Documento', icon: FileText, color: 'text-red-500' },
    { label: 'Otro', icon: FileIcon, color: 'text-gray-500' },
  ]

  const handleDelete = (name: string) => {
    alert(`Borrando el archivo: ${name}`)
  }

  return (
    <div className="bg-linear-to-br from-[#FFF3E6] to-[#F3E8FF] h-[calc(100vh-64px)] p-8 min-h-screen">
        {isUploadMenuOpen && (
            <div 
                className="fixed inset-0 z-40 bg-transparent" 
                onClick={() => setIsUploadMenuOpen(false)}
            />
        )}

        <div className="flex justify-between items-center mb-8">
            <div>
                <h1 className="text-2xl font-bold text-gray-800">Mi Portfolio</h1>
                <p className="text-gray-500">Sube y visualiza tus archivos subidos</p>
            </div>
            
            <div className="relative">
                <button 
                    onClick={() => setIsUploadMenuOpen(!isUploadMenuOpen)}
                    className="relative z-50 flex items-center gap-2 px-5 py-2.5 rounded-xl bg-linear-to-r from-[#F97316] to-[#8B5CF6] text-white hover:shadow-xl transition-all active:scale-95"
                >
                    <Upload size={18} />
                    <span className="font-medium">Subir Archivo</span>
                    <ChevronDown size={16} className={`transition-transform ${isUploadMenuOpen ? 'rotate-180' : ''}`} />
                </button>

                {isUploadMenuOpen && (
                    <div className="absolute right-0 mt-2 w-56 bg-white rounded-2xl shadow-2xl border border-gray-100 p-2 z-50 overflow-hidden animate-in fade-in slide-in-from-top-2 duration-200">
                        {uploadOptions.map((option) => (
                            <button
                            key={option.label}
                            className="w-full flex items-center gap-3 px-3 py-2.5 text-sm text-gray-700 hover:bg-gray-50 rounded-xl transition-colors group"
                            onClick={() => setIsUploadMenuOpen(false)}
                            >
                            <div className="p-2 rounded-lg bg-gray-50 group-hover:bg-white shadow-xs">
                                <option.icon size={18} className={option.color} />
                            </div>
                            <span className="font-medium">{option.label}</span>
                            </button>
                        ))}
                    </div>
                )}
            </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
            {files.map((file) => (
                <FileCard 
                    key={file.id} 
                    file={file} 
                    onDelete={handleDelete} 
                />
            ))}

            <div 
                onClick={() => setIsUploadMenuOpen(true)}
                className="border-2 border-dashed border-gray-200 rounded-2xl flex flex-col items-center justify-center p-8 hover:border-black/20 hover:bg-gray-100/50 transition-all cursor-pointer text-gray-400 group"
            >
                <div className="p-3 rounded-full bg-gray-100 group-hover:bg-white mb-2">
                    <Plus size={24} />
                </div>
                <span className="text-sm font-medium group-hover:text-[#8B5CF6]">Agregar nuevo</span>
            </div>
        </div>
    </div>
  )
}
