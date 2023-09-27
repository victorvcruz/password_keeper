// LoginContext.tsx
import React, { createContext, useContext, useState, ReactNode } from 'react';

interface Login {
  email: string;
  password: string;
}

interface LoginContextProps {
  user: Login | null;
  login: (loginData: Login) => void;
  logout: () => void;
}

export const LoginContext = createContext<LoginContextProps | undefined>(undefined);

export const LoginProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<Login | null>(null);

  const login = (loginData: Login) => {
    setUser(loginData);
  };

  const logout = () => {
    setUser(null);
  };

  return (
    <LoginContext.Provider value={{ user, login, logout }}>
      {children}
    </LoginContext.Provider>
  );
};

export const useLogin = () => {
  const context = useContext(LoginContext);
  if (!context) {
    throw new Error('useLogin must be used within a LoginProvider');
  }
  return context;
};
