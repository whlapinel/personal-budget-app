'use server'

import { backendUrls } from '@/app/constants/backend-urls';
import { cookies } from 'next/headers'
import type { User } from '@/app/lib/data/definitions';
import { SignJWT } from 'jose'
import bcrypt from 'bcrypt'
import { useSession } from '../session-context';
import { sessionLifeSpan } from '@/app/constants/session-life-span';

const secret = process.env.SECRET_KEY
const key = new TextEncoder().encode(secret);

export async function signInAction(prevState: any, formData: FormData): Promise<{message: string, user: Partial<User> | null}> {

    console.log('formData', formData)
    const user: Partial<User> = {
        email: formData.get('email')?.toString(),
        password: formData.get('password')?.toString()
    }

    // retrieve encrypted password from backend
    const encryptedPassword = await getEncryptedPassword(user.email);
    let match: boolean;

    // compare passwords
    if (encryptedPassword instanceof Error) {
        return {message: 'Error signing in', user: null}
    }
    try {
        match = await bcrypt.compare(user.password, encryptedPassword);
        console.log('match:', match);
        if (!match) {
            return {message: 'username and/or password incorrect', user: null}
        }
    } catch (err) {
        console.error(err);
        return {message: 'Error signing in', user: null}
    }

    async function getEncryptedPassword(email: string): Promise<string | Error> {
        let encryptedPassword: string;
        try {
            const response = await fetch(`${backendUrls.users}/${user!.email}`, {
                headers: {
                    'API_KEY': process.env.API_KEY!
                },
                cache: 'no-store',
            });
            const data = await response.json();
            if (data.status === 404) {
                return new Error('User not found');
            }
            console.log('data', data);
            encryptedPassword = data.password;
        } catch (err) {
            console.error(err);
            return '';
        }
        return encryptedPassword;
    }

    const expires: number = Date.now() + sessionLifeSpan;

    async function encrypt(payload: any) {
        return await new SignJWT(payload)
            .setProtectedHeader({ alg: "HS256" })
            .setIssuedAt()
            .setExpirationTime(expires)
            .sign(key);
    }

    const session = await encrypt({ user, expires: expires });

    // Save the session in a cookie
    cookies().set("session", session, { expires: expires, httpOnly: true });
    return (
        {
            message: "Signed in successfully",
            user: {...user, expiration: expires}
        }
    )
}
