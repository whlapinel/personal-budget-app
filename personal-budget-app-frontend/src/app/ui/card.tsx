import { twMerge } from "tailwind-merge"

export default function Card({ children, className }: { children: React.ReactNode, className?: string }) {
    return (
        <>
            <div className={twMerge("overflow-hidden bg-white shadow sm:rounded-lg", className)}>
                <div className="px-4 py-5 sm:p-6">
                    <div className="flow-root">
                        <div className="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                            <div className="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
                                {children}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}
