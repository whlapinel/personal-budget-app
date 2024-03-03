'use server'

import { backendUrls } from '@/app/constants/backend-urls';
import { cookies } from 'next/headers'
import type { User } from '@/app/lib/data/definitions';
import {SignJWT} from 'jose'
import dotenv from 'dotenv';

dotenv.config()

const secret = process.env.SECRET_KEY
const key = new TextEncoder().encode(secret);

export async function signInAction(prevState: any, formData: FormData) {
    console.log('formData', formData)
    const user: any = {
        username: formData.get('email'),
        password: formData.get('password')
    }

    async function encrypt(payload: any) {
        return await new SignJWT(payload)
            .setProtectedHeader({ alg: "HS256" })
            .setIssuedAt()
            .setExpirationTime("20 sec from now")
            .sign(key);
    }

    const expires = new Date(Date.now() + 10 * 1000);
    const session = await encrypt({ user, expires });

    // Save the session in a cookie
    cookies().set("session", session, { expires, httpOnly: true });
    return (
    {
        message: "Signed in successfully",
        session: session
    }
)

// console.log('stringified user', JSON.stringify(user));

// // send user to backend
// let message: string;
// let token: string;
// try {
//     const response = await fetch(backendUrls.signin, {
//         cache: 'no-store',
//         method: 'POST',
//         body: JSON.stringify(user)
//     });
//     const headers = response.headers;
//     console.log('headers', headers);
//     const cookies = headers.getSetCookie();
//     console.log('cookies', cookies);
//     const data = await response.json();
//     console.log('data', data);
//     token = data;
//     if (data.error) return { message: 'Error signing in' }
//     if (data.message === 'user not found' || data.message === 'password does not match') {
//         return { message: 'Username or password is incorrect' }
//     }
// } catch (err) {
//     console.error(err);
//     return { message: 'Error signing in' }
// }
// return (
//     {
//         message: "Signed in successfully",
//         token: token
//     }
// )
}
