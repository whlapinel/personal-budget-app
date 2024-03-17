'use server'

import { backendUrls } from "@/app/constants/backend-urls";
import type { Account } from "@/app/lib/data/definitions";
import { cookies } from "next/headers";

export default async function addAccountAction(prevState: any, formData: any){
    const email = cookies().get('email')?.value;
    const account: Partial<Account> = {
        name: formData.get('name')?.toString(),
        type: formData.get('type')?.toString(),
        balance: Number(formData.get('balance')),
        email: email,
        bankName: formData.get('bankName')?.toString(),
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