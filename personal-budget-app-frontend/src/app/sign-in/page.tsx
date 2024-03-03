'use client'

import { SubmitButton } from "../ui/submit-button";
import {signInAction} from "./actions";
import { useFormState } from "react-dom";

const initialState = {
  message: '',
  token: ''
}

export default function SignInPage() {
  const [state, formAction] = useFormState(signInAction, initialState)


  return (
    <div className="flex flex-col items-center">
      <h1>Sign In</h1>
      <form action={formAction}>
        <div>
          <label htmlFor="username">Email</label>
          <input type="text" name="username" />
        </div>
        <div>
          <label htmlFor="password">Password</label>
          <input type="password" name="password" />
        </div>
        <SubmitButton>Sign In</SubmitButton>
      </form>
    </div>

  )
}
