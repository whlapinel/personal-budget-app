'use client'

import { useFormState } from 'react-dom'
import addTransactionAction from './actions/add-transaction-action';
import { SubmitButton } from '@/app/ui/submit-button';
import { useSession } from '@/app/session-context';
import { Account } from '@/app/lib/data/definitions';

const initialState: { message: string | null } = {
    message: null,
}
export default function AddTransactionForm({ accounts }: { accounts: Account[] }) {
    const [state, formAction] = useFormState(addTransactionAction, initialState)
    const { user } = useSession();
    if (!user) return null;
    const email = user.email;

    return (
        <form className="flex flex-col items-center justify-center self-center" action={formAction}>
            <label htmlFor="accountID">Account</label>
            <select name="accountID">
                {accounts?.map((account: Account) => (
                    <option value={account.id}>{account.name}</option>
                ))}
            </select>
            <label htmlFor="date">Date</label>
            <input type="date" name="date" />
            <label htmlFor="payee">Payee</label>
            <input type="text" name="payee" />
            <label htmlFor="memo">Memo</label>
            <input type="text" name="memo" />
            <label htmlFor="amount">Amount</label>
            <input type="float" name="amount" />
            <SubmitButton>Add</SubmitButton>
            <p>{state.message}</p>
        </form>
    )


}