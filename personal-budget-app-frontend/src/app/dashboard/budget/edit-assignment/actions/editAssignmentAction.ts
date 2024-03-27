'use server'

import { backendUrls } from "@/app/constants/backend-urls";
import { Assignment } from "@/app/lib/data/definitions";
import { cookies } from "next/headers"
import { revalidatePath } from "next/cache";

export async function editAssignmentAction(prevState: any, formData: FormData) {
    const email = cookies().get('email')?.value!;
    console.log('running editAssignmentAction');
    console.log('email:', email);
    console.log('formData:', formData);
    const month = Number(formData.get('month'));
    const year = Number(formData.get('year'));
    const categoryID = Number(formData.get('categoryID'));
    const amount = Number(formData.get('amount')) * 100; // convert to cents
    const assignment: Assignment = {
        email: email,
        categoryID: categoryID,
        month: month,
        year: year,
        amount: amount
    }
    console.log('Assignment:', assignment);
    const response = await fetch(`${backendUrls.assignments}`, {
        method: 'POST',
        headers: {
            'API_KEY': process.env.API_KEY!,
        },
        body: JSON.stringify(assignment)
    });
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    const data = await response.json();
    console.log('addGoalAction data:', data);
    revalidatePath('/dashboard/budget');
    return ({
        message: 'Assignment updated',
        assignment: assignment
    })
}