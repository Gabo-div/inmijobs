import React from 'react'
import { 
  Calendar,
  Download,
  File as FileIcon, 
  FileImage, 
  FileText, 
  FileVideo, 
  Headphones, 
  Plus, 
  Trash2
} from 'lucide-react'
import { portfolioService } from '@/services/portfolioService'

export interface FileData {
  id: string
  original_name: string
  file_type: string
  file_size: number
  download_url: string
  created_at: number
}

interface FileCardProps {
  file: FileData
  onDelete?: (id: string) => void
}

export const FileCard: React.FC<FileCardProps> = ({ file, onDelete }) => {
    const timestamp = file.created_at < 10000000000 ? file.created_at * 1000 : file.created_at;
    const formattedDate = new Intl.DateTimeFormat('es-ES', {
        day: '2-digit',
        month: 'short',
        year: 'numeric'
    }).format(new Date(timestamp));
  
    const getFileIcon = (type: string) => {
    const t = type.toLowerCase();
      if (t.includes('image')) return <FileImage className="text-blue-500" />
      if (t.includes('pdf')) return <FileText className="text-red-500" />
      if (t.includes('video')) return <FileVideo className="text-purple-500" />
      if (t.includes('audio')) return <Headphones className="text-pink-500" />
      return <FileIcon className="text-gray-500" />
    }

    const formatSize = (bytes: number) => {
      if (bytes === 0) return '0 Bytes';
      const k = 1024;
      const sizes = ['Bytes', 'KB', 'MB', 'GB'];
      const i = Math.floor(Math.log(bytes) / Math.log(k));
      return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
    }

    const handleDownload = async () => {
      try {
          await portfolioService.downloadFile(file.id, file.original_name);
      } catch (error) {
          console.error("Error descargando:", error);
          alert("No se pudo descargar el archivo");
      }
    };
  

  return (
    <div className="bg-white p-4 rounded-2xl border border-gray-200 hover:shadow-md transition-all group">
      <div className="flex items-start justify-between mb-4">
        <div className="p-3 bg-gray-50 rounded-xl group-hover:bg-white transition-colors shadow-sm">
          {getFileIcon(file.file_type)}
        </div>
        
        {onDelete && (
            <button 
                onClick={(e) => {
                  e.stopPropagation()
                  onDelete(file.id)
                }}
                className="text-gray-400 hover:text-red-500 hover:bg-red-50 p-1.5 rounded-lg cursor-pointer transition-colors"
            >
                <Trash2 size={18} />
            </button>
        )}
      </div>
      
      <div className="space-y-1">
        <div className='flex items-center justify-between'>
          <h3 className="font-semibold text-gray-700 truncate text-[15px]" title={file.original_name}>
            {file.original_name}
          </h3>

          <button 
            onClick={handleDownload}
            className="p-2 hover:bg-blue-50 text-blue-600 rounded-lg cursor-pointer transition-colors"
            title="Descargar"
          >
            <Download size={18} />
          </button>
        </div>
        <div className="flex justify-between items-center text-[12px] text-gray-400 font-medium">
          <span>{formatSize(file.file_size)}</span>
          <span className="flex items-center gap-1">
            <span className="size-1 bg-gray-300 rounded-full" />
            <Calendar size={12} /> {formattedDate}
          </span>
        </div>
      </div>
    </div>
  )
}