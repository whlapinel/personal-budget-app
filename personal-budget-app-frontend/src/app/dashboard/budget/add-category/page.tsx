'use client'

import { useFormState } from 'react-dom'
import addCategoryAction from './actions/add-category-action';
import { useEffect } from 'react';
import { useState } from 'react';

const initialState = {
  message: '',
}

export default function AddCategoryPage() {
  const [state, formAction] = useFormState(addCategoryAction, initialState)
  const [token, setToken] = useState('');

  console.log(state.message);


  useEffect(() => {
    setToken(localStorage.getItem('token') || '');
    console.log('token', token);    
  }, [])

  return (
    <form className="flex flex-col items-center justify-center self-center" action={formAction}>
      <label htmlFor='name'>Name</label>
      <input type="text" name='name' required />
      <input type="hidden" name='token' value={token} required />
      <button type="submit">Submit</button>
    </form>
  )
}
