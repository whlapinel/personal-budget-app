
export default function DashboardLayout({
    children, budget, register
}: {
    children: React.ReactNode,
    budget: any,
    register: any
}) {
    return (
        <div>
            {children}
            {budget}
            {register}
        </div>
    )
}
