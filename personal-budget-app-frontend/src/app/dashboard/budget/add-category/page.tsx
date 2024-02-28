'use client'

import { useFormState } from 'react-dom'
import addCategoryAction from './actions/add-category-action';


export default function AddCategoryPage() {
  const [state, formAction] = useFormState(addCategoryAction, { message: '' })

  const userID = '1';

  return (
    <form className="flex flex-col items-center justify-center self-center" action={formAction}>
      <label htmlFor='name'>Name</label>
      <input type="text" name='name' required />
      <input type="hidden" name='userid' value={userID} required />
      <button type="submit">Submit</button>
    </form>
  )
}
