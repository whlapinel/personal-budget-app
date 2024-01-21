
import Link from 'next/link'
import { navigation } from '@/app/navigation'
import { UserButton } from "@clerk/nextjs";

export default function NavBar() {
    return (
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
            <div className="flex h-16 justify-between">
                <div className="flex">
                    <div className="flex flex-shrink-0 items-center">
                        <span>logo placeholder</span>
                    </div>
                    <div className="hidden sm:-my-px sm:ml-6 sm:flex sm:space-x-8">
                        {navigation.map((item) => (
                            <Link
                                key={item.name}
                                href={item.href}
                                className={
                                    'inline-flex items-center border-b-2 px-1 pt-1 text-sm font-medium'
                                }
                            >
                                {item.name}
                            </Link>
                        ))}
                    </div>
                </div>
                <div className="hidden sm:ml-6 sm:flex sm:items-center">
                    <UserButton />
                </div>
            </div>
        </div>
    )
}
