'use client'

import { useFormState } from 'react-dom'
import addTransactionAction from './actions/add-transaction-action';
import { SubmitButton } from '@/app/ui/submit-button';
import { useSession } from '@/app/session-context';
import { Account } from '@/app/lib/data/definitions';
import { Category } from '@/app/lib/data/definitions'
import { Radio, RadioGroup } from '@/app/ui/radio';
import Form from '@/app/ui/form';
import { Input } from '@/app/ui/input';
import type { FormHiddenInfo } from '@/app/ui/form';

const initialState: { message: string | null } = {
    message: null,
}
export default function AddTransactionForm({ accounts, categories }: { accounts: Account[], categories: Category[] }) {
    const [state, formAction] = useFormState(addTransactionAction, initialState)
    const { user } = useSession();
    if (!user) return null;
    const email = user.email;

    const hiddenInfo: FormHiddenInfo[] = [
        {
            name: 'email',
            value: email
        }
    ]

    return (
        <Form title="Add Transaction" formAction={formAction} state={state} hiddenInfo={hiddenInfo}>
            <label htmlFor="accountID">Account</label>
            <select name="accountID">
                {accounts?.map((account: Account) => (
                    <option key={account.id} value={account.id}>{account.name}</option>
                ))}
            </select>
            <label htmlFor="date">Date</label>
            <Input type="date" name="date" />
            <label htmlFor="payee">Payee</label>
            <Input type="text" name="payee" />
            <label htmlFor="type">Type</label>
            <select name='type'>
                <option value="debit">Debit</option>
                <option value="credit">Credit</option>
            </select>
            <label htmlFor="memo">Memo</label>
            <Input type="text" name="memo" />
            <label htmlFor="amount">Amount</label>
            <Input type="float" name="amount" />
            <label htmlFor="categoryID">Category</label>
            <select name="categoryID">
                {categories?.map((category: Category) => (
                    <option key={category.id} value={category.id}>{category.name}</option>
                ))}
            </select>
            <input type="hidden" name="email" value={email} />
        </Form>
    )


}