import { createContext, useContext, useEffect, useState, ReactNode } from 'react';
import { getMe } from '../lib/api';

type AuthCtx = {
  user: string;
  loading: boolean;
  refresh: () => Promise<void>;
};

const Ctx = createContext<AuthCtx>({
  user: '',
  loading: true,
  refresh: async () => {},
});

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState('');
  const [loading, setLoading] = useState(true);

  const refresh = async () => {
    try {
      const { user } = await getMe();
      setUser(user || '');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    void refresh();
  }, []);

  return <Ctx.Provider value={{ user, loading, refresh }}>{children}</Ctx.Provider>;
}

export const useAuth = () => useContext(Ctx);
