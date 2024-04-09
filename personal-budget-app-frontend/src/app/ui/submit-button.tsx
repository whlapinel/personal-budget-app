'use client'

import { revalidatePath } from 'next/cache'
import { useFormStatus } from 'react-dom'
import { Button } from './button'

export function SubmitButton({ children, handleSubmit, className}: { children: React.ReactNode, handleSubmit?: () => void, className?: string}) {
    const { pending } = useFormStatus()

    if (pending) {
        console.log('pending code triggered')
        if (handleSubmit)
        handleSubmit();
    }

    return (
        <Button type="submit" aria-disabled={pending} className={className}>
            {children}
        </Button>
    )
}