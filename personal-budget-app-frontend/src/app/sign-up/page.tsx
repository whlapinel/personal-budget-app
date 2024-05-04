'use client'

import { SubmitButton } from '../ui/submit-button';
import signUpAction from './signup-action';
import { useFormState } from 'react-dom'
import { useRouter } from 'next/navigation';
import Form from '../ui/form'
import { Input } from '../ui/input';


export default function SignUpPage() {
  const [state, formAction] = useFormState(signUpAction, { message: null })

  console.log(state);

  return (
    <div className='flex flex-col items-center justify-center min-h-[80vh]'>
      <Form title="Sign Up" formAction={formAction} state={state} submitBtnTitle='Sign Up!'>
          <label htmlFor='firstName'>First Name</label>
          <Input type="text" name='firstName' required />
          <label htmlFor='lastName'>Last Name</label>
          <Input type="text" name='lastName' required />
          <label htmlFor='email'>Email</label>
          <Input type="email" name='email' required />
          <label htmlFor='password'>Password</label>
          <Input type="password" name='password' required />
      </Form>
    </div>
  )
}