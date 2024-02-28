'use server'

import { backendUrls } from '@/app/constants/backend-urls'
import { redirect } from 'next/navigation';
import { postCategory } from '@/app/lib/data/post-data';

export default async function addCategoryAction(prevState: any, formData: FormData) {
    console.log('formData', formData)

    const category: any = {
        name: formData.get('name'),
    }

    console.log('stringified category', JSON.stringify(category));

    const data = await postCategory(category);
    console.log('data', data);
    if (data.error) {
        return {message: 'Category not created'}
    }

    return (
        {message: 'Category created'}
    )
}
