

export default function SideNav({ children }: { children: React.ReactNode }) {

    return (
            <div className='flex flex-col gap-2 ps-0 p-2'>
                {children}
            </div>
    )
}
