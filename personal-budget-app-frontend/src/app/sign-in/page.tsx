'use client'

import { SubmitButton } from "../ui/submit-button";
import { signInAction } from "./actions";
import { useRouter } from "next/navigation";
import { useFormState } from "react-dom";
import { useEffect } from "react";
import { useSession } from "../session-context";
import type { User } from "../lib/data/definitions";
import Form from "@/app/ui/form"
import { Input } from "../ui/input";

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
    <div className="flex flex-col justify-center items-center min-h-[80vh]">
      <Form title="Sign In" formAction={formAction} state={state} submitBtnTitle="Sign In">
          <label htmlFor="email">Email</label>
          <Input type="email" name="email" />
          <label htmlFor="password">Password</label>
          <Input type="password" name="password" />
      </Form>
    </div>
  )
}
