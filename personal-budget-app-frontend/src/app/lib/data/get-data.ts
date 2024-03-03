import { revalidatePath } from "next/cache";
// import all definitions
import { User, Account, Transaction, BudgetCategory } from './definitions';
import { backendUrls } from "@/app/constants/backend-urls";

export async function getUser(userID: string): Promise<User> {
    const data = await fetch(`http://127.0.0.1:8080/users/${userID}`, { cache: 'no-store' });
    const users = await data.json();
    console.log(users);
    return users
}



export async function getAccounts(userID: string): Promise<Account[]> {
    const data = await fetch(`${backendUrls.accounts}/${userID}`, {
        cache: 'no-store',
        headers: {
            'Authorization': ``
        }
    });
    const accounts = await data.json();
    console.log(accounts);
    return accounts
}

export async function getTransactions(userID: string): Promise<Transaction[]> {
    try {
        const data = await fetch(`${backendUrls.transactions}/${userID}`, {
            cache: 'no-store',
            headers: {
                'Authorization': ``
            }
        });
        const transactions = await data.json();
        console.log(transactions);
        return transactions
    } catch (err) {
        console.log(err);
        return []
    }
}

export async function getCategories(): Promise<BudgetCategory[]> {
    try {
        const data = await fetch(`${backendUrls.categories}`, {
            cache: 'no-store',
            headers: {
                'Authorization': ``
            }
        });
        const categories = await data.json();
        console.log(categories);
        return categories
    } catch (err) {
        console.log(err);
        return []
    }
}