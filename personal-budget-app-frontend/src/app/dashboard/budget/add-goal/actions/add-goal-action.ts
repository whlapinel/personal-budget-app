'use server'

import { Goal } from "@/app/lib/data/definitions"

export default function addGoalAction(prevState: any, formData: FormData) {

    const goal: Partial<Goal> = {
        categoryID: Number(formData.get('categoryID')),
        name: formData.get('name')?.toString(),
        amount: Number(formData.get('amount')) * 100, // convert to cents
        targetDate: new Date(formData.get('targetDate')?.toString()!),
    }

    console.log('running addGoalAction');
    return ({
        message: 'Goal added',
        goal: goal
    })

}