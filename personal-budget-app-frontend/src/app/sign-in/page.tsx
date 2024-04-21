'use client'

import { SubmitButton } from "../ui/submit-button";
import { signInAction } from "./actions";
import { useRouter } from "next/navigation";
import { useFormState } from "react-dom";
import { useEffect } from "react";
import { useSession } from "../session-context";
import type { User } from "../lib/data/definitions";

const initialState: { message: string | null, user: User | null } = {
  message: null,
  user: null
}

export default function SignInPage() {
  const [state, formAction] = useFormState(signInAction, initialState);
  const { user, setUser } = useSession();

  const router = useRouter();

  useEffect(() => {
    if (state.user) {
      console.log("SignInPage useEffect running! expires:", state.user?.expiration);
      const hasNull = Object.values(state.user).some(value => (value === null || value === undefined));
      if (!hasNull) {
        setUser(state.user as User);
        setTimeout(() => {
          router.push("/dashboard");
        }, 2000);
      }
      console.log("user: ", user);
    }
  }, [setUser, state.user, user, router]);

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
