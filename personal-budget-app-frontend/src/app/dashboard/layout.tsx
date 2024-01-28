'use client';

import SideNav from '@/app/ui/side-nav';
import { useSearchParams } from 'next/navigation';
import Link from 'next/link';


export default function DashboardLayout({
    children, budget, register, reports,
}: {
    children: React.ReactNode,
    budget: any,
    register: any,
    reports: any,
}) {

    const searchParams = useSearchParams();
    const view = searchParams.get('view');

    return (
        <>
            <div className='grid grid-cols-[1fr_8fr]'>
                <SideNav>
                        <Link className=" bg-blue-700 rounded p-2 text-gray-50" href='/dashboard?view=budget'>Budget</Link>
                        <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={'/dashboard?view=register'}>Register</Link>
                        <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={'/dashboard?view=reports'}>Reports</Link>
                </SideNav>
                {children}
                <div>
                    {view === 'budget' && budget}
                    {view === 'register' && register}
                    {view === 'reports' && reports}
                    {view === null && <div>Nothing to see here</div>}
                </div>
            </div>
        </>
    )
}
