'use server'

import { backendUrls } from "@/app/constants/backend-urls";
import { Goal } from "@/app/lib/data/definitions"
import { Periodicity } from "@/app/lib/data/definitions"
import { ca } from "date-fns/locale";
import { revalidatePath } from "next/cache";

export default async function addGoalAction(prevState: any, formData: FormData) {
    console.log('running addGoalAction');

    // validate formData

    const properties = ['categoryID', 'name', 'amount', 'targetDate', 'email', 'periodicity'];
    for (const property of properties) {
        if (!formData.has(property)) {
            return ({
                message: `Missing ${property}`
            })
        }
    }
    const categoryID = Number(formData.get('categoryID'));
    if (isNaN(categoryID)) {
        return ({
            message: 'Invalid category'
        })
    }
    const name = formData.get('name')?.toString();
    const amount = Number(formData.get('amount')) * 100; // convert to cents
    if (isNaN(amount) || amount < 0) {
        return ({
            message: 'Invalid amount'
        })
    }
    const dateFormEntry = formData.get('targetDate'); 
    const targetDate = new Date(dateFormEntry?.toString()!);
    if (isNaN(targetDate.getTime())) {
        return ({
            message: 'Invalid target date'
        })
    }
    const emailFormEntry = formData.get('email');
    const email = emailFormEntry?.toString();
    if (!email || !email.includes('@')) {
        return ({
            message: 'Invalid email'
        })
    }
    const periodicityFormEntry = formData.get('periodicity')!;
    const periodicity: Periodicity = periodicityFormEntry.toString() as Periodicity;
    const goal: Partial<Goal> = {
        categoryID: categoryID,
        name: name,
        amount: amount, // convert to cents
        targetDate: targetDate.toDateString(),
        email: email,
        periodicity: periodicity,
    }
    const response = await fetch(backendUrls.goals, {
        method: 'POST',
        headers: {
            'API_KEY': process.env.API_KEY!,
        },
        body: JSON.stringify(goal)
    });
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    const data = await response.json();
    console.log('addGoalAction data:', data);
    revalidatePath('/dashboard/budget');

    return ({
        message: 'Goal added',
        goal: goal
    })

}