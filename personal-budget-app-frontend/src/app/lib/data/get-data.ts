import { revalidatePath } from "next/cache";
// import all definitions
import { User, Account, Transaction, Category } from './definitions';
import { backendUrls } from "@/app/constants/backend-urls";

const API_KEY: string = process.env.API_KEY!;
export async function getUser(email: string): Promise<User> {
    const data = await fetch(`${backendUrls.users}/${email}`);
    const users = await data.json();
    console.log(users);
    return users
}

export async function getAccounts(email: string): Promise<Account[]> {
    const data = await fetch(`${backendUrls.accounts}/${email}`, {
        headers: {
            'API_KEY': API_KEY
        }
    });
    const accounts = await data.json();
    console.log(accounts);
    return accounts
}

export async function getTransactions(email: string): Promise<Transaction[]> {
    try {
        const data = await fetch(`${backendUrls.transactions}/${email}`, {
            headers: {
                'API_KEY': API_KEY
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

export async function getCategories(email: string): Promise<Category[]> {
    console.log('getCategories email: ', email)
    try {
        const data = await fetch(`${backendUrls.categories}/${email}`, {
            headers: {
                'API_KEY': API_KEY
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