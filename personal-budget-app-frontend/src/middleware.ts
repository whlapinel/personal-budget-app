'use server'

import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
import { getToken, refreshToken } from "@/app/lib/data/auth";
import { sessionLifeSpan } from './app/constants/session-life-span';

export async function middleware(request: NextRequest) {
  console.log('running middleware');
  const session = await getToken();
  console.log('session', session);
  if (!session || session.expires < new Date()) {
    return NextResponse.redirect(new URL('/not-authorized', request.url));
  }
  const response = NextResponse.next();
  response.cookies.set('email', session.user.email, { httpOnly: true, expires: Date.now() + sessionLifeSpan })
  return response;
}

export const config = {
  matcher: ['/dashboard/:path*'],
}