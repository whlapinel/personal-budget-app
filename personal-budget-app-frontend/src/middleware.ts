'use server'

import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
import { getToken, refreshToken } from "@/app/lib/data/auth";

// FIXME: middleware is calling refreshToken and error returned is "Cookies can only be modified in 
// a Server Action or Route Handler."

export async function middleware(request: NextRequest) {
  console.log('running middleware');
  const session = await getToken();
  if (!session || session.expires < new Date()) {
    return NextResponse.redirect(new URL('/not-authorized', request.url));
  }
  return NextResponse.next();
}

export const config = {
  matcher: ['/dashboard/:path*'],
}