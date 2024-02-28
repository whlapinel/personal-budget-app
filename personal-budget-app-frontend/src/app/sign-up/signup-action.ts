'use server'
import { backendUrls } from '@/app/constants/backend-urls'
import { redirect } from 'next/navigation';
import { postUser } from '@/app/lib/data/post-data';
import { User } from '@/app/lib/data/definitions';

export default async function signUpAction(prevState: any, formData: FormData) {
    console.log('formData', formData)
    const user: any = {
        username: formData.get('username'),
        password: formData.get('password'),
        firstName: formData.get('firstName'),
        lastName: formData.get('lastName'),
        email: formData.get('email'),
    }

    console.log('stringified user', JSON.stringify(user));

    // send user to backend
    try {
        const data = await postUser(user);
        console.log('data', data);
        if (data.error) {
            return {message: 'User not created'}
        }
    } catch (err) {
        console.error(err);
        return {message: 'User not created'}
    }
    return (
        {message: 'User created'}
    )
}
