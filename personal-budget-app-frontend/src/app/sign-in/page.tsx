'use client'

import { SubmitButton } from "../ui/submit-button";
import {signInAction} from "./actions";
import { useFormState } from "react-dom";

const initialState = {
  message: '',
  session: ''
}

export default function SignInPage() {
  const [state, formAction] = useFormState(signInAction, initialState)

  
  return (
    <div className="flex flex-col items-center">
      <h1>Sign In</h1>
      <form action={formAction}>
        <div>
          <label htmlFor="email">Email</label>
          <input type="email" name="email" />
        </div>
        <div>
          <label htmlFor="password">Password</label>
          <input type="password" name="password" />
        </div>
        <SubmitButton>Sign In</SubmitButton>
        <p>{state.message}</p>
      </form>
    </div>

  )
}
