'use client'

import { useFormState } from 'react-dom'
import addTransactionAction from './actions/add-transaction-action';
import { SubmitButton } from '@/app/ui/submit-button';

const initialState: {message: string | null} = {
  message: null,
}
export default function AddTransactionPage() {
  const [state, formAction] = useFormState(addTransactionAction, initialState)

  return (
      <form className="flex flex-col items-center justify-center self-center" action={formAction}>
          <label htmlFor="date">Date</label>
          <input type="date" id="date" name="date" />
          <label htmlFor="description">Description</label>
          <input type="text" id="description" name="description" />
          <label htmlFor="amount">Amount</label>
          <input type="float" id="amount" name="amount" />
          <SubmitButton>Add</SubmitButton>
          <p>{state.message}</p>
      </form>
  )
}
