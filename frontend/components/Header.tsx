import Link from 'next/link';

export default function Header({ user }: { user?: string }) {
  return (
    <header style={{ padding: '12px 16px', borderBottom: '1px solid #eee' }}>
      <nav style={{ display: 'flex', gap: 16 }}>
        <Link href="/">Home</Link>
        <Link href="/dashboard">Dashboard</Link>
        <Link href="/swagger/index.html" target="_blank" rel="noreferrer">Swagger</Link>
        {!user && <Link href="/login">Login</Link>}
      </nav>
    </header>
  );
}