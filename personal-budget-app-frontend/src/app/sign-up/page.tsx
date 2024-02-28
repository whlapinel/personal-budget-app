'use client'

import { SubmitButton } from '../ui/submit-button';
import signUpAction from './signup-action';
import { useFormState } from 'react-dom'
import { useRouter } from 'next/navigation';


export default function SignUpPage() {
  const [state, formAction] = useFormState(signUpAction, { message: '' })

  console.log(state);
  const router = useRouter();
  if (state.message === 'User created') {
    setTimeout(() => {
      router.push('/dashboard')
    }, 1000)
  }


  return (
    <div className='flex flex-col items-center'>
      <h1>Sign Up</h1>
      <form action={formAction} className='flex flex-col'>
        <label htmlFor='firstName'>First Name</label>
        <input type="text" name='firstName' required />
        <label htmlFor='lastName'>Last Name</label>
        <input type="text" name='lastName' required />
        <label htmlFor='username'>Email</label>
        <input type="email" name='email' required />
        <label htmlFor='username'>Username</label>
        <input type="text" name='username' required />
        <label htmlFor='password'>Password</label>
        <input type="password" name='password' required />
        <SubmitButton>Submit</SubmitButton>
        <p>{state.message}</p>
      </form>
    </div>
  )
}