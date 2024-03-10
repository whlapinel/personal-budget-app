'use server'

import { backendUrls } from "@/app/constants/backend-urls";
import { SignJWT, jwtVerify } from "jose";
import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

const secret = process.env.SECRET_KEY
console.log("secret: ", secret);

const key = new TextEncoder().encode(secret);
console.log("key", key);


const sessionDuration = 60 * 1000; // 1 minute

export async function encrypt(payload: any) {
  return await new SignJWT(payload)
    .setProtectedHeader({ alg: "HS256" })
    .setIssuedAt()
    .setExpirationTime(Date.now() + sessionDuration)
    .sign(key);
}

export async function decrypt(input: string): Promise<any> {
  const { payload } = await jwtVerify(input, key, {
    algorithms: ["HS256"],
  });
  return payload;
}

export async function getSession() {
  const session = cookies().get("session")?.value;
  if (!session) return null;
  return await decrypt(session);
}

export async function getEncryptedPassword(email: string): Promise<string> {
    let encryptedPassword: string;
    try {
      const response = await fetch(`${backendUrls.users}/${email}`, {
        headers: {
          'API_KEY': process.env.API_KEY!
        },
        cache: 'no-store',
      });
      const data = await response.json();
      console.log('data', data);
      encryptedPassword = data.password;
    } catch (err) {
      console.error(err);
      return '';
    }
  return encryptedPassword;
}


export async function updateSession(): Promise<boolean> {
  const session = cookies().get("session")?.value;
  if (!session) return false;

  // Refresh the session so it doesn't expire
  const parsed = await decrypt(session);
  console.log("parsed: ", parsed);
  
  parsed.expires = new Date(Date.now() + sessionDuration);
  cookies().set('session', await encrypt(parsed), {httpOnly: true, expires: parsed.expires});
  // const res = NextResponse.next();
  // res.cookies.set({
  //   name: "session",
  //   value: await encrypt(parsed),
  //   httpOnly: true,
  //   expires: parsed.expires,
  // });
  return true;
}