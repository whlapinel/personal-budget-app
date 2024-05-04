'use client'

import { useFormState } from 'react-dom'
import addCategoryAction from './actions/add-category-action';
import { SubmitButton } from '@/app/ui/submit-button';
import { revalidatePath } from 'next/cache';
import Form from '@/app/ui/form';
import { Input } from '@/app/ui/input';

const initialState: { message: string | null } = {
  message: null,
}

export default function AddCategoryPage() {
  const [state, formAction] = useFormState(addCategoryAction, initialState)

  console.log(state.message);

  return (
    <Form title="Add Category" formAction={formAction} state={state} submitBtnTitle='Add Category'>
      <label htmlFor='name'>Name</label>
      <Input type="text" name='name' required />
    </Form>
  )
}
