'use server'

import { backendUrls } from "@/app/constants/backend-urls";
import { Transaction } from "@/app/lib/data/definitions"

export default async function addTransactionAction(prevState: any, formData: any){

    // validate formData
    const transaction: Partial<Transaction> = {
        account: Number(formData.get('account')),
        date: new Date(formData.get('date')?.toString()!),
        payee: formData.get('payee')?.toString(),
        amount: Number(formData.get('amount')),
        memo: formData.get('memo')?.toString(),
        category: Number(formData.get('category')),
    }

    // send formData to server
    const response = await fetch(backendUrls.transactions, {
        method: 'POST',
        headers: {
            'API_KEY': process.env.API_KEY!,
        },
        body: JSON.stringify(transaction)
    });
    const data = await response.json();
    console.log('running addTransactionAction. data:', data);


    return ({
        message: 'Transaction added'
    })
}