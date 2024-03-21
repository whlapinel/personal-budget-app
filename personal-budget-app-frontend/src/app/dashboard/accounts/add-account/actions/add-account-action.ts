'use server'

import { backendUrls } from "@/app/constants/backend-urls";
import type { Account } from "@/app/lib/data/definitions";
import { cookies } from "next/headers";

export default async function addAccountAction(prevState: any, formData: any){
    const email = cookies().get('email')?.value;
    const account: Partial<Account> = {
        email: email,
        name: formData.get('name')?.toString(),
        type: formData.get('type')?.toString(),
        bankName: formData.get('bankName')?.toString(),
        startingBalance: Number(formData.get('startingBalance')) * 100, // convert to cents
    }
    console.log('running addAccountAction. account:', account);
    try {
        const response = await fetch(`${backendUrls.accounts}`, {
            method: 'POST',
            headers: {
                'API_KEY': process.env.API_KEY!,
            },
            body: JSON.stringify(account)
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        console.log('addAccountAction data:', data);

        return ({
            message: 'Account added'
        })
    } catch (error) {
        console.error('addAccountAction error:', error);
        return ({
            message: 'Account not added'
        })
    }
}