'use client'

import { Category } from "@/app/lib/data/definitions";
import { useFormState } from "react-dom";
import addGoalAction from "./actions/add-goal-action";
import { Goal } from "@/app/lib/data/definitions"
import { HTMLInputTypeAttribute } from "react";
import { SubmitButton } from "@/app/ui/submit-button";
import { useSession } from "@/app/session-context";

const initialState: any = {
    message: null,
    goal: null
}
export default function AddGoalForm({ category }: { category: Category }) {
    const [state, formAction] = useFormState(addGoalAction, initialState)
    const { user } = useSession();
    if (!user) return null;
    const email = user.email;

    console.log('AddGoalForm category:', category);
    

    return (
        <form action={formAction}>
            <label htmlFor="name">Name</label>
            <input type="text" name="name" />
            <label htmlFor="amount">Amount</label>
            <input type="float" name="amount" />
            <label htmlFor="targetDate">Date</label>
            <input type="date" name="targetDate" />
            <select name="periodicity">
                <option value="monthly">Monthly</option>
                <option value="yearly">Yearly</option>
                <option value="onetime">One Time</option>
                <option value="weekly">Weekly</option>
                <option value="biweekly">Biweekly</option>
                <option value="quarterly">Quarterly</option>
            </select>
            <input type="hidden" name="categoryID" value={category.id} />
            <input type="hidden" name="email" value={category.email} />
            <SubmitButton>Add</SubmitButton>
            <p>{state.message}</p>
        </form>

    )
}