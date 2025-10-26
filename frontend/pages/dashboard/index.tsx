import Header from '../../components/Header';
import { useAuth } from '../../contexts/AuthContext';
import Link from 'next/link';

export default function Dashboard() {
  const { user, loading } = useAuth();

  // Middleware already blocks unauthenticated users before render.
  if (loading) return null;

  return (
    <>
      <Header user={user} />
      <main style={{ padding: 24 }}>
        <h1>Dashboard</h1>
        <p>Welcome, <b>{user}</b> 🎉</p>
        <ul style={{ marginTop: 16 }}>
          <li>Try protected API calls from here, e.g. Logic/Engine hello</li>
          <li>
            <Link
              href={`/api/logic/hello?name=${encodeURIComponent(user || "Patrick")}`}
              target="_blank"
              rel="noreferrer"
            >
              /api/logic/hello
            </Link>
          </li>
          <li>
            <Link
              href={`/api/engine/hello?name=${encodeURIComponent(user || "Patrick")}`}
              target="_blank"
              rel="noreferrer"
            >
              /api/engine/hello
            </Link>
          </li>
        </ul>
      </main>
    </>
  );
}
