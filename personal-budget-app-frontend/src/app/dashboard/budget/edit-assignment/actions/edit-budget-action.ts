'use server'

import { backendUrls } from "@/app/constants/backend-urls";
import { Assignment, MonthlyBudget } from "@/app/lib/data/definitions";
import { cookies } from "next/headers"
import { revalidatePath } from "next/cache";

export async function editBudgetAction(prevState: any, formData: FormData) {
    const email = cookies().get('email')?.value!;
    console.log('running editAssignmentAction');
    console.log('email:', email);
    console.log('formData:', formData);
    const month = Number(formData.get('month'));
    const year = Number(formData.get('year'));
    const categoryID = Number(formData.get('categoryID'));
    const assigned = Number(formData.get('amount')) * 100; // convert to cents
    const assignment: MonthlyBudget = {
        email: email,
        categoryID: categoryID,
        month: month,
        year: year,
        assigned: assigned
    }
    console.log('Assignment:', assignment);

    try {
        const response = await fetch(`${backendUrls.monthlyBudgets}`, {
            method: 'POST',
            headers: {
                'API_KEY': process.env.API_KEY!,
            },
            body: JSON.stringify(assignment)
        });
        const data = await response.json();
        console.log('addGoalAction data:', data);
        if (data.error) {
            return ({
                message: data.error,
                assignment: null
            })
        }
        revalidatePath('/dashboard/budget');
        return ({
            message: 'Assignment updated',
            assignment: assignment
        })
    } catch (err) {
        console.error('Error in editBudgetAction:', err);
        return ({
            message: err,
            assignment: null
        })


    }
}