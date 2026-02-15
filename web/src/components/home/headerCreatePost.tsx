import type { UserModel } from "@/models/UserModel"
import { ImageIcon, Smile, Video } from "lucide-react"
import { CreatePostModal } from "../post/modals/CreatePostModal"
import { useState } from "react"

interface Props {
    userData?: UserModel
}

export const HeaderCreatePost = ({ userData }: Props) => {
    const [isModalOpen, setIsModalOpen] = useState(false)

    return (
        <section className="bg-white p-4 rounded-2xl shadow-sm border border-gray-100">
            <div className="flex gap-3 items-center">
                <img src={userData?.Image} className="w-10 h-10 rounded-full" alt="User" />
                <button
                    type="button"
                    className="text-black flex-1 bg-gray-100 hover:bg-gray-200 transition-colors rounded-full px-4 py-2 text-gray-500 cursor-pointer text-sm text-left"
                    onClick={() => setIsModalOpen(true)}
                >
                    ¿Qué estás pensando, {userData?.Name}?
                </button>
                <div className="flex gap-2 text-gray-500">
                    <Video size={20} className="text-red-500 cursor-pointer" />
                    <ImageIcon size={20} className="text-green-500 cursor-pointer" />
                    <Smile size={20} className="text-yellow-500 cursor-pointer" />
                </div>
            </div>

            <CreatePostModal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)} />
        </section>
    )
}