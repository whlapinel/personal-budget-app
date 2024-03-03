'use client'

import {signOut} from './sign-out'

export default function SignOutPage() {
    return (
      <button onClick={async () => signOut()}>Sign Out</button>
    )
  }