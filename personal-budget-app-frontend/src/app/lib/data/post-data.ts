import { backendUrls } from "@/app/constants/backend-urls";
import { User } from "./definitions";

export async function postUser(user: Partial<User>) {
    const response = await fetch(backendUrls.signUp, {
        cache: 'no-store',
        method: 'POST',
        body: JSON.stringify(user)
    });
    const data = await response.json();
    return data;
}

