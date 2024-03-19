'use server'

import { backendUrls } from "@/app/constants/backend-urls";
import { Transaction } from "@/app/lib/data/definitions";

export default async function addTransactionAction(prevState: any, formData: any){


    // validate formData
    const transaction: Partial<Transaction> = {
        accountID: Number(formData.get('accountID')),
        date: new Date(formData.get('date')?.toString()!),
        payee: formData.get('payee')?.toString(),
        amount: Number(formData.get('amount')),
        memo: formData.get('memo')?.toString(),
        categoryID: formData.get('categoryID') ? Number(formData.get('categoryID')) : null,
    }
    console.log('running addTransactionAction. transaction:', transaction);

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
    if (data.error) {
        return ({message: data.error})
    }
    return ({
        message: 'Transaction added'
    })
}