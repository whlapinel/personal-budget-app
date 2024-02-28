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

export async function getCategories(userID: number): Promise<BudgetCategory[]> {
    const data = await fetch(`${backendUrls.categories}/${userID}`, { cache: 'no-store' });
    const categories = await data.json();
    console.log(categories);
    return categories
}