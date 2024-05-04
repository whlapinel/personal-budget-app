'use client'

import Form from "@/app/ui/form";

import { SubmitButton } from "@/app/ui/submit-button";
import { useFormState } from "react-dom";
import addAccountAction from "./actions/add-account-action";
import { Input } from "@/app/ui/input";

const initialState: { message: string | null } = {
    message: null,
}

export default function AddAccountPage() {
    const [state, formAction] = useFormState(addAccountAction, initialState)

    return (
        <Form title="Add Account" formAction={formAction} state={state} submitBtnTitle="Add Account">
            <label htmlFor='name'>Name</label>
            <Input type="text" name='name' required />
            <label htmlFor='bankName'>Bank Name</label>
            <Input type="text" name='bankName' required />
            <label htmlFor='type'>Type</label>
            <Input type="text" name='type' required />
            <label htmlFor='startingBalance'>Starting Balance</label>
            <Input type="float" name='startingBalance' required />
        </Form>
    )

}