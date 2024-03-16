'use client'

import { useFormState } from 'react-dom'
import addCategoryAction from './actions/add-category-action';
import { SubmitButton } from '@/app/ui/submit-button';
import { revalidatePath } from 'next/cache';

const initialState: {message: string | null} = {
  message: null,
}

export default function AddCategoryPage() {
  const [state, formAction] = useFormState(addCategoryAction, initialState)

  console.log(state.message);

  return (
    <form className="flex flex-col items-center justify-center self-center" action={formAction}>
      <label htmlFor='name'>Name</label>
      <input type="text" name='name' required />
      <SubmitButton>Submit</SubmitButton>
      <p>{state.message}</p>
    </form>
  )
}
