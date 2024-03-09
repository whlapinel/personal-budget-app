'use server'
import { postUser } from '@/app/lib/data/post-data';
import { SignJWT, jwtVerify } from "jose";
import { cookies } from "next/headers";
import { key } from '@/app/lib/data/auth'

export default async function signUpAction(prevState: any, formData: FormData) {

    console.log('formData', formData)
    const user: any = {
        email: formData.get('email'),
        password: formData.get('password'),
        firstName: formData.get('firstName'),
        lastName: formData.get('lastName'),
    }

    async function encrypt(payload: any) {
        return await new SignJWT(payload)
            .setProtectedHeader({ alg: "HS256" })
            .setIssuedAt()
            .setExpirationTime("1 minute from now")
            .sign(key);
    }

    // Create the session
    const expires = new Date(Date.now() + 10 * 1000);
    const session = await encrypt({ user, expires });

    // Save the session in a cookie
    cookies().set("session", session, { expires, httpOnly: true });

    // encrypt password before sending to backend
    const bcrypt = require('bcrypt');
    const saltRounds = 10;
    const hash = await bcrypt.hash(user.password, saltRounds);
    user.password = hash;
    try {
        const data = await postUser(user);
        if (data.error) {
            return { message: data.error }
        }
    } catch (error) {
        console.error('Error creating user', error);
        return { message: 'Error creating user' }
    }
    return (
        { message: 'User created' }
    )
}
