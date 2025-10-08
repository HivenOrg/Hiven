"use client";

import { createContext, useContext, useEffect, useState } from "react";
import type {
  LoginUserInput,
  RegisterUserInput,
} from "@/lib/validators/auth.schema";

interface AuthContextType {
  isAuthenticated: boolean;
  login: (data: LoginUserInput) => Promise<void>;
  logout: () => Promise<void>;
  register: (data: RegisterUserInput) => Promise<void>;
  user: { id: string; name: string } | null;
}

const authContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = (): AuthContextType => {
  const context = useContext(authContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

interface AuthProviderProps {
  children: React.ReactNode;
  token?: string;
  user?: { id: string; name: string };
}

export function AuthProvider({ children }: AuthProviderProps) {
  const [authState, setAuthState] = useState<{
    isLoading: boolean;
    isAuthenticated: boolean;
    user: AuthContextType["user"] | null;
  }>({
    isLoading: true,
    isAuthenticated: false,
    user: null,
  });

  async function login(data: LoginUserInput): Promise<void> {
    console.log(`Logging in with ${data.email} and ${data.password}`);
  }

  async function logout(): Promise<void> {
    console.log("Logging out");
  }

  async function register(data: RegisterUserInput): Promise<void> {
    console.log(`Registering with ${data.email} and ${data.password}`);
  }

  useEffect(() => {
    setAuthState((prev) => ({ ...prev, isLoading: false }));
  }, [])

  return (
    <authContext.Provider
      value={{
        isAuthenticated: authState.isAuthenticated,
        user: authState.user,
        login,
        logout,
        register,
      }}
    >
      {children}
    </authContext.Provider>
  );
}
