'use server'

import { SignJWT, jwtVerify } from "jose";
import { cookies } from "next/headers";

const secret = process.env.SECRET_KEY
const key = new TextEncoder().encode(secret);

const tokenLifeSpan = 60000; // 1 minute 
function getExpiration() {
  return Date.now() + tokenLifeSpan;
}

export async function encrypt(payload: any) {
  const expires = getExpiration();
  return await new SignJWT(payload)
    .setProtectedHeader({ alg: "HS256" })
    .setIssuedAt()
    .setExpirationTime(expires)
    .sign(key);
}

export async function decrypt(input: string): Promise<any> {
  const { payload } = await jwtVerify(input, key, {
    algorithms: ["HS256"],
  });
  console.log(Date.now())
  console.log(payload);
  return payload;
}

export async function getToken() {
  const token = cookies().get("session")?.value;
  if (!token) return null;
  return await decrypt(token);
}

export async function refreshToken() {
  const token = cookies().get("session")?.value
  if (!token) return;
  // Refresh the session so it doesn't expire
  const parsed = await decrypt(token);
  parsed.expires = getExpiration();
  const newToken = await encrypt(parsed);
}