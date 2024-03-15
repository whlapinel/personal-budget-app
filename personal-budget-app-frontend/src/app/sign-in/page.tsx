'use client'

import { userAgent } from "next/server";
import { SubmitButton } from "../ui/submit-button";
import {signInAction} from "./actions";
import { useFormState } from "react-dom";
import { useEffect } from "react";
import { useSession } from "../session-context";

const initialState = {
  message: '',
  user: null
}

export default function SignInPage() {
  const [state, formAction] = useFormState(signInAction, initialState);

  const { user, setUser } = useSession();

  useEffect(() => {
    console.log("SignInPage useEffect running! expires:", state.user?.expiration);
    setUser(state.user!);
    console.log("user: ", user);
  }, [state])

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
        <p>{"token expires in " + Math.round((user?.expiration! - Date.now())/1000) + " seconds."}</p>
      </form>
    </div>

  )
}
