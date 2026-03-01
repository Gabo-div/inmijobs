import { createFileRoute } from '@tanstack/react-router';
import React, { useState } from 'react';
import type { Friend, Photo, Post } from '@/features/profile';
import { ProfileHeader } from '@/features/profile';
import { ProfileTabTodo } from '@/features/profile/components/ProfileTabTodo';
import { ProfileTabInfo } from '@/features/profile/components/ProfileTabInfo';
import { authClient } from '@/lib/auth';
import { useAppUser } from '@/lib/userContext';
import { Button } from '@/components/ui/button';


export const Route = createFileRoute('/(private)/profile')({
  component: ProfilePage,
});

const TABS = ['Todo', 'Información', 'Amigos', 'Fotos', 'Reels', 'Más'] as const;
type TabType = typeof TABS[number];

function ProfilePage() {
  const [activeTab, setActiveTab] = useState<TabType>('Todo');
  const { user, setUser } = useAppUser();
  const [pickerOpen, setPickerOpen] = useState(false);
  const [editingField, setEditingField] = useState<null | 'avatar' | 'banner'>(null);

  // show loading until session/context user available
  const { data: session } = authClient.useSession();
  if (!session || !user) {
    return <div className="p-4">Cargando perfil…</div>;
  }

  // existing user object comes from context
  // (friends, photos, posts may still rely on it below)


  const [friends] = useState<Array<Friend>>([
    { id: '1', name: 'Alexis Diaz', avatarUrl: '/assets/mascota/2.png' },
    { id: '2', name: 'Eduardo Crespo', avatarUrl: '/assets/brand/icon.png' },
  ]);

  const [photos] = useState<Array<Photo>>([
    { id: '1', url: '/assets/brand/isologo.png' },
    { id: '2', url: '/assets/mascota/2.png' },
    { id: '3', url: '/assets/brand/icon.png' },
  ]);

  const [posts] = useState<Array<Post>>([
    {
      id: '1',
      author: user,
      content: 'Mascota corporativa y concepts arts, para una red social dedicada a turistas e inmigrantes...',
      imageUrl: '/assets/mascota/1.png',
      likes: 12,
      comments: 5,
      timestamp: '2 días',
    }
  ]);

  const renderTabContent = () => {
    switch (activeTab) {
      case 'Todo':
        return <ProfileTabTodo user={user} friends={friends} photos={photos} posts={posts} />;
      case 'Información':
        return <ProfileTabInfo />;
      case 'Amigos':
        return <div className="p-4 bg-white mt-4 rounded-xl border">Vista de Amigos (En construcción)</div>;
      case 'Fotos':
        return <div className="p-4 bg-white mt-4 rounded-xl border">Vista de Fotos (En construcción)</div>;
      default:
        return <div className="p-4 bg-white mt-4 rounded-xl border">Vista de {activeTab} (En construcción)</div>;
    }
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;
    const url = URL.createObjectURL(file);
    setUser((prev) => {
      if (!prev) return prev;
      if (editingField === 'banner') {
        return { ...prev, coverUrl: url };
      }
      // default to avatar
      return { ...prev, avatarUrl: url };
    });
    setPickerOpen(false);
    setEditingField(null);
  };

  return (
    <div className="container mx-auto max-w-4xl py-3 px-2 sm:px-0 space-y-2">
      <ProfileHeader
        user={user}
        onEditProfile={(what) => {
          setEditingField(what);
          setPickerOpen(true);
        }}
      />

      {pickerOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/40">
          <div className="bg-white dark:bg-card rounded-lg p-4 w-full max-w-sm shadow-lg">
            <h3 className="text-lg font-semibold mb-2">
              {editingField === 'banner' ? 'Cambiar imagen de banner' : 'Cambiar foto de perfil'}
            </h3>
            <input
              type="file"
              accept="image/*"
              onChange={handleFileChange}
              className="mb-4"
            />
            <div className="flex justify-end gap-2">
              <Button variant="ghost" onClick={() => { setPickerOpen(false); setEditingField(null); }}>Cancelar</Button>
            </div>
          </div>
        </div>
      )}

      <div className="flex border-b dark:border-border overflow-x-auto -mx-2 px-2 mt-2">
        {TABS.map((tab) => {
          const isActive = activeTab === tab;
          return (
            <button
              key={tab}
              onClick={() => setActiveTab(tab)}
              className={`px-3 py-2 font-medium text-sm whitespace-nowrap transition-colors relative
                                ${isActive
                  ? 'text-blue-600'
                  : 'text-muted-foreground hover:text-foreground hover:bg-gray-50 dark:hover:bg-muted/20 rounded-t-md'
                }`}
            >
              {tab}
              {isActive && (
                <span className="absolute bottom-0 left-0 w-full h-0.5 bg-blue-600 rounded-t-full" />
              )}
            </button>
          );
        })}
      </div>

      <div className="min-h-[400px] transition-all">
        {renderTabContent()}
      </div>
    </div>
  );
}
