import { backendUrls } from "@/app/constants/backend-urls";
import { User } from "./definitions";

export async function postUser(user: Partial<User>) {
    const API_KEY: string = process.env.API_KEY!;
    const response = await fetch(backendUrls.users, {
        cache: 'no-store',
        headers: {
            'API_KEY': API_KEY,
        },
        method: 'POST',
        body: JSON.stringify(user)
    });
    const data = await response.json();
    console.log('running postUser. data:', data);
    return data;
}

