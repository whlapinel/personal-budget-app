'use client'

import { useSession } from '@/app/session-context'
import { Button } from '../ui/button';
import { Link } from '../ui/link';

export default function SignOutPage() {
  const { signOut } = useSession();
  return (
    <div className='flex flex-col items-center justify-center min-h-[80vh]'>
      <Button color="blue" onClick={async () => signOut()}>Sign Out</Button>
    </div>
  )
}