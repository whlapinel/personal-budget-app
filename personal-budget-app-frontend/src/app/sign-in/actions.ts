'use server'

import { backendUrls } from '@/app/constants/backend-urls';
import { cookies } from 'next/headers'
import type { User } from '@/app/lib/data/definitions';
import { SignJWT } from 'jose'
import dotenv from 'dotenv';
import bcrypt from 'bcrypt'

dotenv.config()

const secret = process.env.SECRET_KEY
const key = new TextEncoder().encode(secret);

export async function signInAction(prevState: any, formData: FormData) {
    console.log('formData', formData)
    const user: any = {
        email: formData.get('email'),
        password: formData.get('password')
    }

    // retrieve encrypted password from backend
    const encryptedPassword = await getEncryptedPassword(user.email);

    try {
        const match = await bcrypt.compare(user.password, encryptedPassword);
        console.log('match:', match);
    } catch (err)
    {
        console.error(err);
        return { message: 'Error signing in' }
    }

    async function getEncryptedPassword(email: string): Promise<string> {
        let encryptedPassword: string;
        try {
            const response = await fetch(`${backendUrls.users}/${user.email}`, {
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

    async function encrypt(payload: any) {
        return await new SignJWT(payload)
            .setProtectedHeader({ alg: "HS256" })
            .setIssuedAt()
            .setExpirationTime("1 minute from now")
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
