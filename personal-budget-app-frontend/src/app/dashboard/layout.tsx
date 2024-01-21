'use client';

import { useState } from 'react';
import { Button } from '@/app/ui/button';

export default function DashboardLayout({
    children, budget, register, reports
}: {
    children: React.ReactNode,
    budget: any,
    register: any,
    reports: any,
}) {

    const [view, setView] = useState('budget');

    return (
        <>
            {children}
            <div className='flex justify-center gap-2 my-2'>
                <Button color={view === 'budget'?'blue':'dark'} onClick={() => setView('budget')}>Budget</Button>
                <Button color={view === 'register'?'blue':'dark'} onClick={() => setView('register')}>Register</Button>
                <Button color={view === 'reports'?'blue':'dark'} onClick={() => setView('reports')}>Reports</Button>
            </div>
            <div>
                {view === 'budget' && budget}
                {view === 'register' && register}
                {view === 'reports' && reports}
                {view === null && <div>Nothing to see here</div>}
            </div>
        </>
    )
}
