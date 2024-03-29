import SideNav from '@/app/ui/side-nav';
import {Link} from '@/app/ui/link';

export default function DashboardLayout({
    children
}: {
    children: React.ReactNode
}) {

    const currMonth = new Date().getMonth();
    const currYear = new Date().getFullYear();


    return (
        <>
            <div className='grid grid-cols-[1fr_8fr]'>
                <SideNav>
                    <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={`/dashboard/budget?month=${currMonth}&year=${currYear}`}>Budget</Link>
                    <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={'/dashboard/accounts'}>Accounts</Link>
                    <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={'/dashboard/transactions'}>Transactions</Link>
                    <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={'/dashboard/reports'}>Reports</Link>
                </SideNav>
                {children}
            </div>
        </>
    )
}
