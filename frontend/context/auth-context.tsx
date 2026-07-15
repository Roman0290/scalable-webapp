"use client";

import { createContext, useContext, useMemo, useState } from "react";
import type { ReactNode } from "react";
import type { AuthSession } from "@/types/auth";

type AuthContextValue = {
  session: AuthSession | null;
  setSession: (session: AuthSession | null) => void;
};

const AuthContext = createContext<AuthContextValue | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [session, setSession] = useState<AuthSession | null>(null);

  const value = useMemo(
    () => ({ session, setSession }),
    [session],
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export function useAuthContext() {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error("useAuthContext must be used inside AuthProvider");
  }

  return context;
}