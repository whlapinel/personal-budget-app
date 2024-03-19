'use client'

import {useSession} from '@/app/session-context'

export default function SignOutPage() {
  const {signOut} = useSession();
    return (
      <button onClick={async () => signOut()}>Sign Out</button>
    )
  }