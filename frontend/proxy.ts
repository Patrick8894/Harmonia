import { NextResponse, NextRequest } from 'next/server';

export async function proxy(req: NextRequest) {
  const { pathname, origin, search } = req.nextUrl;

  // Only protect these paths
  const PROTECTED = [/^\/dashboard(\/|$)/];
  const needsAuth = PROTECTED.some((re) => re.test(pathname));
  if (!needsAuth) return NextResponse.next();

  // Call the existing /api/auth/me endpoint
  const res = await fetch(`${origin}/api/auth/me`, {
    headers: { cookie: req.headers.get('cookie') || '' },
  }).catch(() => null);

  let user = '';
  if (res && res.ok) {
    const data = await res.json().catch(() => ({}));
    user = (data as any).user || '';
  }

  if (user) {
    // user is logged in
    return NextResponse.next();
  }

  // Not authenticated → redirect to login
  const loginUrl = new URL('/login', origin);
  loginUrl.searchParams.set('next', pathname + (search || ''));
  return NextResponse.redirect(loginUrl);
}

export const config = {
  matcher: ['/dashboard/:path*'],
};
