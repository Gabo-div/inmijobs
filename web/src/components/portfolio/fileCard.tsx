import React from 'react'
import { 
  File as FileIcon, 
  FileImage, 
  FileText, 
  FileVideo, 
  Headphones, 
  Plus 
} from 'lucide-react'

export interface FileData {
  id: number
  name: string
  type: string
  size: string
  date: string
}

interface FileCardProps {
  file: FileData
  onDelete?: (name: string) => void
}

export const FileCard: React.FC<FileCardProps> = ({ file, onDelete }) => {
  
    const getFileIcon = (type: string) => {
        if (type.includes('png') || type.includes('jpg') || type.includes('jpeg')) return <FileImage className="text-blue-500" />
        if (type.includes('pdf')) return <FileText className="text-red-500" />
        if (type.includes('mp4')) return <FileVideo className="text-purple-500" />
        if (type.includes('mp3') || type.includes('wav')) return <Headphones className="text-pink-500" />
        return <FileIcon className="text-gray-500" />
    }
  

  return (
    <div className="bg-white p-4 rounded-2xl border border-gray-200 hover:shadow-md transition-all group cursor-pointer active:scale-[0.98]">
      <div className="flex items-start justify-between mb-4">
        <div className="p-3 bg-gray-50 rounded-xl group-hover:bg-white transition-colors shadow-sm">
          {getFileIcon(file.type)}
        </div>
        
        <button 
          onClick={(e) => {
            e.stopPropagation()
            onDelete?.(file.name)
          }}
          className="text-gray-400 hover:text-red-500 hover:bg-red-50 p-1.5 rounded-lg transition-colors"
        >
          <Plus size={18} className="rotate-45" />
        </button>
      </div>
      
      <div className="space-y-1">
        <h3 className="font-semibold text-gray-700 truncate text-[15px]" title={file.name}>
          {file.name}
        </h3>
        
        <div className="flex justify-between items-center text-[12px] text-gray-400 font-medium">
          <span>{file.size}</span>
          <span className="flex items-center gap-1">
            <span className="size-1 bg-gray-300 rounded-full" />
            {file.date}
          </span>
        </div>
      </div>
    </div>
  )
}