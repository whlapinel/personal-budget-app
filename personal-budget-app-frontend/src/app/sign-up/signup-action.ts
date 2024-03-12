'use server'

import { postUser } from '@/app/lib/data/post-data';
import { cookies } from "next/headers";
import { encrypt } from '@/app/lib/data/auth'
import { sessionLifeSpan } from '@/app/constants/session-life-span';

export default async function signUpAction(prevState: any, formData: FormData) {

    console.log('formData', formData)
    const user: any = {
        email: formData.get('email'),
        password: formData.get('password'),
        firstName: formData.get('firstName'),
        lastName: formData.get('lastName'),
    }

    // Create the session
    const expires = Date.now() + sessionLifeSpan;
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
        {
            message: 'User created',
            user: {
                ...user,
                expires: expires
            }
        }
    )
}
