'use client'

import { revalidatePath } from 'next/cache'
import { useFormStatus } from 'react-dom'

export function SubmitButton({ children, handleSubmit }: { children: React.ReactNode, handleSubmit?: () => void}) {
    const { pending } = useFormStatus()

    if (pending) {
        console.log('pending code triggered')
        if (handleSubmit)
        handleSubmit();
    }

    return (
        <button type="submit" aria-disabled={pending}>
            {children}
        </button>
    )
}