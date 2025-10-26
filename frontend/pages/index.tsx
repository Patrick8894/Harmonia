import { useAuth } from '../contexts/AuthContext';
import Header from '../components/Header';
import { logout } from '../lib/api';
import Link from "next/link";

export default function Home() {
  const { user, loading, refresh } = useAuth();

  const onLogout = async () => {
    await logout();
    await refresh();
  };

  return (
    <>
      <Header user={user} />
      <main style={{ padding: 24 }}>
        <h1>Harmonia Home</h1>
        {loading ? (
          <p>Loading...</p>
        ) : user ? (
          <>
            <p>Signed in as <b>{user}</b></p>
            <button onClick={onLogout}>Logout</button>
          </>
        ) : (
          <>
            <p>You are not signed in.</p>
            <Link href="/login">
              <button>Login</button>
            </Link>
          </>
        )}
        <hr style={{ margin: '24px 0' }} />
        <ul>
          <li>
            <Link href="/api/hello" target="_blank" rel="noreferrer">
              /api/hello
            </Link>
          </li>
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
