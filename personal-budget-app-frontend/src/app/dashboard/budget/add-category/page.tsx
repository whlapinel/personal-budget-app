'use client'

import { useFormState } from 'react-dom'
import addCategoryAction from './actions/add-category-action';
import { useEffect } from 'react';
import { useState } from 'react';
import { useSession } from '@/app/session-context';

const initialState: {message: string | null} = {
  message: null,
}

export default function AddCategoryPage() {
  const [state, formAction] = useFormState(addCategoryAction, initialState)
  const {user} = useSession();

  console.log(state.message);


  return (
    <form className="flex flex-col items-center justify-center self-center" action={formAction}>
      <label htmlFor='name'>Name</label>
      <input type="text" name='name' required />
      <button type="submit">Submit</button>
    </form>
  )
}
