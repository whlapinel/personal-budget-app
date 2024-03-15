'use server'

import { backendUrls } from '@/app/constants/backend-urls'
import { redirect } from 'next/navigation';
import {cookies} from 'next/headers';
import type {Category} from '@/app/lib/data/definitions'

export default async function addCategoryAction(prevState: any, formData: FormData) {
    const API_KEY: string = process.env.API_KEY!;
    console.log('formData', formData)
    console.log('prevState', prevState)

    const category: Partial<Category> = {
        name: formData.get('name')!.toString(),
        email: cookies().get('email')!.value
    }

    console.log('stringified category', JSON.stringify(category));


    const response = await fetch(backendUrls.categories, {
        headers: {
            'API_KEY': API_KEY,
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
    return (
        {
            message: 'Category created'
        }
    )
}
