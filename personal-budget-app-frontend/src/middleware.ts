import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
import { getSession, updateSession } from "@/app/lib/data/auth";

export async function middleware(request: NextRequest) {
  console.log('running middleware');
  const session = await getSession();
  if (!session) {
    return NextResponse.redirect(new URL('/sign-in', request.url));
  }
  if (session.expires < new Date()) {
    return NextResponse.redirect(new URL('/sign-in', request.url));
  }
  return await updateSession(request);
}

export const config = {
  matcher: ['/dashboard/:path*'],
}