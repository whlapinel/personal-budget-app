'use client'

import { useFormState } from "react-dom";
import { SubmitButton } from "@/app/ui/submit-button";
import { useSession } from "@/app/session-context";
import { editBudgetAction } from "../actions/edit-budget-action";
import { Category } from "@/app/lib/data/definitions";
import {convertToDollars} from "@/app/lib/util/cents-to-dollars";
import Form from "@/app/ui/form";
import { Input } from "@/app/ui/input";

const initialState: any = {
    message: null,
    assignment: null
}

export default function EditAssignmentForm({ category, month, year, currAssignmentAmount }: { category: Category, month: number, year: number, currAssignmentAmount: number }) {
    const [state, formAction] = useFormState(editBudgetAction, initialState)

    console.log('EditAssignmentForm category:', category);

    const hiddenInfo = [
        {
            name: 'categoryID',
            value: category.id
        },
        {
            name: 'month',
            value: month
        },
        {
            name: 'year',
            value: year
        }
    ]

    return (
        <div className="flex justify-center">
                <p>Currently Assigned: {convertToDollars(currAssignmentAmount)}</p>
            <Form formAction={formAction} submitBtnTitle="Update Assignment" title={`Edit Assignment to ${category.name} for ${month}/${year}`} state={state} hiddenInfo={hiddenInfo}>
                <label htmlFor="amount">New Amount</label>
                <Input type="float" name="amount" />
            </Form>
        </div>

    )
}