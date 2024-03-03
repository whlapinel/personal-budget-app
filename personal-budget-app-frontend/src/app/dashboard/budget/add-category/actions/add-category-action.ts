'use server'

import { backendUrls } from '@/app/constants/backend-urls'
import { redirect } from 'next/navigation';
import { postCategory } from '@/app/lib/data/post-data';

export default async function addCategoryAction(prevState: any, formData: FormData) {
    console.log('formData', formData)
    console.log('prevState', prevState)

    const category: any = {
        name: formData.get('name'),
    }
    const token = formData.get('token');

    console.log('stringified category', JSON.stringify(category));


    const response = await fetch(backendUrls.categories, {
        headers: {
            'Authorization': `${token}`
        },
        cache: 'no-store',
        method: 'POST',
        body: JSON.stringify(category)
    });
    const data = await response.json();
    console.log('data', data);
    if (data.error) {
        return { message: 'Category not created' }
    }
    if (data.message === 'no token' || data.message === 'invalid token') {
        return { message: 'Please sign in again.' }
    }
    return (
        {
            message: 'Category created'
        }
    )
}
