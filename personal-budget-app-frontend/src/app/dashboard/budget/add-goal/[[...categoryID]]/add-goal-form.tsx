'use client'

import { Category } from "@/app/lib/data/definitions";
import { useFormState } from "react-dom";
import addGoalAction from "../actions/add-goal-action";
import { useSession } from "@/app/session-context";
import { Input } from "@/app/ui/input";
import Form from "@/app/ui/form";
import type { FormHiddenInfo } from "@/app/ui/form";
import { Select } from "@/app/ui/select";

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

    const hiddenInfo: FormHiddenInfo[] = [
        {
            name: 'categoryID',
            value: category.id
        },
        {
            name: 'email',
            value: category.email
        }
    ]

    return (
        <Form title={`Add Goal for ${category.name}`} formAction={formAction} state={state} hiddenInfo={hiddenInfo}>
            <label htmlFor="name">Name</label>
            <Input type="text" name="name" />
            <label htmlFor="amount">Amount</label>
            <Input type="float" name="amount" />
            <label htmlFor="targetDate">Date</label>
            <Input type="date" name="targetDate" />
            <label htmlFor="periodicity">Periodicity</label>
            <Select name="periodicity">
                <option value="monthly">Monthly</option>
                <option value="onetime">One Time</option>
            </Select>
        </Form>

    )
}