
export default function DashboardLayout({children, budget, register}: {children: React.ReactNode, budget: React.ReactNode, register: React.ReactNode} ) {
  return (
    <>
    {children}
    {budget}
    {register}    
    </>
  )
}
