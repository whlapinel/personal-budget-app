'use server'

import { backendUrls } from '@/app/constants/backend-urls';
import { cookies } from 'next/headers'
import type { User } from '@/app/lib/data/definitions';
import { SignJWT } from 'jose'
import dotenv from 'dotenv';
import bcrypt from 'bcrypt'
import { useSession } from '../session-context';

dotenv.config()

const secret = process.env.SECRET_KEY
const key = new TextEncoder().encode(secret);

export async function signInAction(prevState: any, formData: FormData): Promise<{message: string, user: User}> {

    console.log('formData', formData)
    const user: any = {
        email: formData.get('email'),
        password: formData.get('password')
    }

    // retrieve encrypted password from backend
    const encryptedPassword = await getEncryptedPassword(user.email);

    // compare passwords
    try {
        const match = await bcrypt.compare(user.password, encryptedPassword);
        console.log('match:', match);
    } catch (err)
    {
        console.error(err);
        return { message: 'Error signing in', user: { email: '', id: '', password: '', firstName: '', lastName: '', expiration: 0 }}
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

    const expiration: Date = new Date(Date.now() + 60 * 1000);

    async function encrypt(payload: any) {
        return await new SignJWT(payload)
            .setProtectedHeader({ alg: "HS256" })
            .setIssuedAt()
            .setExpirationTime(expiration)
            .sign(key);
    }

    const session = await encrypt({ user, expires: expiration });

    // Save the session in a cookie
    cookies().set("session", session, { expires: expiration, httpOnly: true });
    return (
        {
            message: "Signed in successfully",
            user: {...user, expiration: expiration.getTime()}
        }
    )
}
