
export function FormContainer({ children, title,  }: { children: React.ReactNode, title: string}) {
    return (
        <div className="bg-white-200 p-4 rounded-lg shadow-md flex flex-col justify-center items-center gap-2">
            <h2 className="text-2xl font-semibold">{title}</h2>
            {children}
        </div>
    )
}