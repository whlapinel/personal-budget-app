'use server'

import { cookies } from "next/headers";

export async function signOut() {
    // Destroy the session
    console.log("running signOut()");
    
    cookies().set("session", "", { expires: new Date(0) });
  }
  