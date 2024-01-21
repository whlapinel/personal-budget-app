'use client';

import { usePathname } from "next/navigation";


export default function PageHeader() {
    const path = usePathname();
    return (
        <header>
            <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
                <h1 className="text-3xl font-bold leading-tight tracking-tight text-gray-900">{path.charAt(1).toUpperCase() + path.slice(2)}</h1>
            </div>
        </header>
    )
}