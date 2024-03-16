'use server'

import { postUser } from '@/app/lib/data/post-data';

export default async function signUpAction(prevState: any, formData: FormData) {

    console.log('formData', formData)
    const user: any = {
        email: formData.get('email'),
        password: formData.get('password'),
        firstName: formData.get('firstName'),
        lastName: formData.get('lastName'),
    }

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
            message: 'User created'
        }
    )
}
