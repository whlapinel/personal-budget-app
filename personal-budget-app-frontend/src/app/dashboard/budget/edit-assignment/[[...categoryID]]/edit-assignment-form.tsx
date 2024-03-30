'use client'

import { useFormState } from "react-dom";
import { SubmitButton } from "@/app/ui/submit-button";
import { useSession } from "@/app/session-context";
import { editBudgetAction } from "../actions/edit-budget-action";
import { Category } from "@/app/lib/data/definitions";
import convertToDollars from "@/app/lib/cents-to-dollars";

const initialState: any = {
    message: null,
    assignment: null
}

export default function EditAssignmentForm({ category, month, year, currAssignmentAmount }: { category: Category, month: number, year: number, currAssignmentAmount: number }) {
    const [state, formAction] = useFormState(editBudgetAction, initialState)

    console.log('EditAssignmentForm category:', category);


    return (
        <div className="flex justify-center">
            <form className="flex justify-center items-center flex-col w-64" action={formAction}>
                <p>Currently Assigned: {convertToDollars(currAssignmentAmount)}</p>
                <label htmlFor="amount">New Amount</label>
                <input type="float" name="amount" />
                <input type="hidden" name="month" value={month} />
                <input type="hidden" name="year" value={year} />
                <input type="hidden" name="categoryID" value={category.id} />
                <SubmitButton>Submit</SubmitButton>
                <p>{state.message}</p>
            </form>
        </div>

    )
}