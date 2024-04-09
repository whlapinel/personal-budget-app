import SideNav from '@/app/ui/side-nav';
import { Link } from '@/app/ui/link';

export default function DashboardLayout({
    children
}: {
    children: React.ReactNode
}) {

    const currMonth = new Date().getMonth(); // 0-indexed
    const currYear = new Date().getFullYear();

    const monthYearParamsString = `month=${currMonth + 1}&year=${currYear}`;


    return (
        <>
            <div className='grid grid-cols-[1fr_8fr]'>
                <SideNav>
                    <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={`/dashboard/budget?${monthYearParamsString}`}>Budget</Link>
                    <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={'/dashboard/accounts'}>Accounts</Link>
                    <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={'/dashboard/transactions'}>Transactions</Link>
                    <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={`/dashboard/reports?${monthYearParamsString}`}>Reports</Link>
                </SideNav>
                <div className='flex items-center justify-center'>
                    {children}
                </div>
            </div>
        </>
    )
}
