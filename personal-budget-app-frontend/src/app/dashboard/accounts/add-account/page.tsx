'use client'

import { SubmitButton } from "@/app/ui/submit-button";
import { useFormState } from "react-dom";
import addAccountAction from "./actions/add-account-action";

const initialState: {message: string | null} = {
    message: null,
}

export default function AddAccountPage() {
    const [state, formAction] = useFormState(addAccountAction, initialState)

    return (
        <form className="flex flex-col items-center justify-center self-center" action={formAction}>
            <label htmlFor='name'>Name</label>
            <input type="text" name='name' required />
            <label htmlFor='bankName'>Bank Name</label>
            <input type="text" name='bankName' required />
            <label htmlFor='type'>Type</label>
            <input type="text" name='type' required />
            <label htmlFor='balance'>Balance</label>
            <input type="float" name='balance' required />
            <SubmitButton>Submit</SubmitButton>
            <p>{state.message}</p>
        </form>
    )

}