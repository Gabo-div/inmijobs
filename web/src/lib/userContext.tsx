import React, { createContext, useContext, useState, useEffect } from "react";
import { authClient } from "./auth";
import type { User } from "@/features/profile/types";

interface UserContextType {
  user: User | null;
  setUser: React.Dispatch<React.SetStateAction<User | null>>;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { data: session } = authClient.useSession();
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    if (session) {
      setUser({
        id: session.user.id,
        name: session.user.name ?? "Usuario",
        role: "",
        tagline: undefined,
        location: "",
        flags: [],
        avatarUrl: session.user.image ?? "/assets/mascota/1.png",
        coverUrl: "",
        friendsCount: 0,
        email: session.user.email ?? "",
      });
    }
  }, [session]);

  return <UserContext.Provider value={{ user, setUser }}>{children}</UserContext.Provider>;
};

export const useAppUser = () => {
  const ctx = useContext(UserContext);
  if (!ctx) {
    throw new Error("useAppUser must be used within UserProvider");
  }
  return ctx;
};
