'use server'


import { revalidatePath } from "next/cache";
// import all definitions
import { User, Account, Transaction, Category, Goal, Assignment, MonthlyBudget, BudgetPageData } from './definitions';
import { backendUrls } from "@/app/constants/backend-urls";


const API_KEY: string = process.env.API_KEY!;

export async function getAccounts(email: string) {
    let accounts: Account[] = [];
    try {
        const response = await fetch(`${backendUrls.accounts}/${email}`, {
            method: 'GET',
            headers: {
                'API_KEY': process.env.API_KEY!,
            },        
        });
        accounts = await response.json();
    } catch (err) {
        console.log(err);
    }
    return accounts;
}


export async function getCategories(email: string): Promise<Category[]> {
    const data = await fetch(`${backendUrls.categories}/${email}`, {
        headers: {
            'API_KEY': API_KEY
        }
    });
    const categories = await data.json();
    return categories
}

export async function getBudgetPageData(email: string, month: number, year: number): Promise<BudgetPageData> {
    const response = await fetch(`${backendUrls.budgetPageData}/${email}/${month}/${year}`, {
        headers: {
            'API_KEY': API_KEY
        }
    });
    const budgetPageData = await response.json();
    console.log("bugetPageData: ", budgetPageData);
    
    return budgetPageData
}

export async function getMonthlyBudgets(email: string, month: number, year: number): Promise<MonthlyBudget[]> {
    const data = await fetch(`${backendUrls.monthlyBudgets}/${email}/${month}/${year}`, {
        headers: {
            'API_KEY': API_KEY
        }
    });
    const monthlyBudgets = await data.json();
    return monthlyBudgets
}

export async function getUser(email: string): Promise<User> {
    const data = await fetch(`${backendUrls.users}/${email}`);
    const users = await data.json();
    console.log(users);
    return users
}

export async function getGoals(email: string, categoryID: number, month: number, year: number): Promise<Goal[]> {
    
    const data = await fetch(`${backendUrls.goals}/${email}/${categoryID}/${month}/${year}`, {
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

export async function getAssignments(categoryID: number): Promise<Assignment[]> {
    const data = await fetch(`${backendUrls.assignments}/${categoryID}`, {
        headers: {
            'API_KEY': API_KEY
        }
    });
    const assignments = await data.json();
    return assignments
}

export async function getAssignmentsByEmail(email: string): Promise<Assignment[]> {
    const data = await fetch(`${backendUrls.assignments}/email/${email}`, {
        headers: {
            'API_KEY': API_KEY
        }
    });
    const assignments = await data.json();
    return assignments
}