'use server'

import { revalidatePath } from "next/cache";
// import all definitions
import { User, Account, Transaction, Category, Goal } from './definitions';
import { backendUrls } from "@/app/constants/backend-urls";


const API_KEY: string = process.env.API_KEY!;
export async function getUser(email: string): Promise<User> {
    const data = await fetch(`${backendUrls.users}/${email}`);
    const users = await data.json();
    console.log(users);
    return users
}

export async function getGoals(email: string): Promise<Goal[]> {
    
    // FIXME I think this URL should be changed somehow... not sure
    const data = await fetch(`${backendUrls.goals}/${email}`, {
        headers: {
            'API_KEY': API_KEY
        }
    });
    const goals = await data.json();
    return goals
}

export async function getCategoryByID(categoryID: string): Promise<Category> {
    // FIXME complete this function
    const data = await fetch (`${backendUrls.categories}/id/${categoryID}`, {
        headers: {
            'API_KEY': API_KEY,
        }
    });
    const category: Category = await data.json();
    console.log(category);
    return category
}

export async function getAccounts(email: string): Promise<Account[]> {
    const data = await fetch(`${backendUrls.accounts}/${email}`, {
        headers: {
            'API_KEY': API_KEY
        }
    });
    const accounts = await data.json();
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
        // parse date string into date object
        transactions.forEach((transaction: Transaction) => {
            transaction.date = new Date(transaction.date);
        });
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
        const categories: Category[] = await data.json();
        for (const category of categories) {
            if (category.goals) {
                for (const goal of category.goals) {
                    goal.targetDate = new Date(goal.targetDate);
                }
            }
        }
        console.log(categories);
        console.log(categories[0].goals)
        return categories
    } catch (err) {
        console.log(err);
        return []
    }
}

export async function getAssignments(categoryID: number): Promise<any> {
    const data = await fetch(`${backendUrls.assignments}/${categoryID}`, {
        headers: {
            'API_KEY': API_KEY
        }
    });
    const assignments = await data.json();
    return assignments
}