'use client';

import SideNav from '@/app/ui/side-nav';
import { useSearchParams } from 'next/navigation';
import Link from 'next/link';

export default function DashboardLayout({
    children
}: {
    children: React.ReactNode
}) {

    const searchParams = useSearchParams();

    return (
        <>
            <div className='grid grid-cols-[1fr_8fr]'>
                <SideNav>
                    <Link className=" bg-blue-700 rounded p-2 text-gray-50" href='/dashboard/budget'>Budget</Link>
                    <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={'/dashboard/register'}>Register</Link>
                    <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={'/dashboard/reports'}>Reports</Link>
                </SideNav>
                {children}
            </div>
        </>
    )
}
