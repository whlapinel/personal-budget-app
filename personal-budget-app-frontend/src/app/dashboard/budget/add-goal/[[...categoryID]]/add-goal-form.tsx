'use client'

import { Category } from "@/app/lib/data/definitions";
import { useFormState } from "react-dom";
import addGoalAction from "../actions/add-goal-action";
import { SubmitButton } from "@/app/ui/submit-button";
import { useSession } from "@/app/session-context";
import { Input } from "@/app/ui/input";
import { FormContainer } from "./form-container";

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
        <FormContainer title={`Add Goal for ${category.name}`}>
            <form className="flex flex-col items-center gap-2" action={formAction}>
                <div className="grid grid-cols-2 gap-1 text-right">
                    <label htmlFor="name">Name</label>
                    <Input type="text" name="name" />
                    <label htmlFor="amount">Amount</label>
                    <Input type="float" name="amount" />
                    <label htmlFor="targetDate">Date</label>
                    <Input type="date" name="targetDate" />
                    <label htmlFor="periodicity">Periodicity</label>
                    <select name="periodicity">
                        <option value="monthly">Monthly</option>
                        <option value="yearly">Yearly</option>
                        <option value="onetime">One Time</option>
                        <option value="weekly">Weekly</option>
                        <option value="biweekly">Biweekly</option>
                        <option value="quarterly">Quarterly</option>
                    </select>
                </div>
                <input type="hidden" name="categoryID" value={category.id} />
                <input type="hidden" name="email" value={category.email} />
                <SubmitButton className=" w-36">Add</SubmitButton>
            </form>
            <p>{state.message}</p>
        </FormContainer>

    )
}