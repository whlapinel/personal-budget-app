'use client';

import Link from 'next/link'
import { navigation } from '@/app/navigation'
import { usePathname } from "next/navigation";
import { useContext } from 'react';
import { SessionProvider, useSession } from '../session-context';

export default function NavBar() {

    const { user, setUser } = useSession();

    const path = usePathname();

    return (
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
            <div className="flex h-16 justify-between">
                <div className="flex">
                    <div className="flex flex-shrink-0 flex-col items-center">
                        <Link href="/">
                            <img src="/money-42.png" className=' size-16' />
                        </Link>
                    </div>
                    <div className="hidden sm:-my-px sm:ml-6 sm:flex sm:space-x-8">
                        {navigation.map((item) => {
                            { item.current = item.href === path }
                            return (
                                <Link
                                    key={item.name}
                                    href={item.href}
                                    className={
                                        item.current
                                            ? 'inline-flex items-center border-b-2 px-1 pt-1 text-sm font-medium border-indigo-500 text-gray-900'
                                            :
                                            'inline-flex items-center border-b-2 px-1 pt-1 text-sm font-medium'
                                    }
                                >
                                    {item.name}
                                </Link>
                            )
                        })}
                    </div>
                </div>
                <span className="inline-flex items-center px-1 pt-1 text-sm font-medium">User: {user ? user.email : "Please Sign In"}</span>
            </div>
        </div>
    )
}
