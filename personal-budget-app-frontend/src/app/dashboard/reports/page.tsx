import getTimeInfo from "@/app/lib/util/time-info";
import Card from "@/app/ui/card";
import { Link } from "@/app/ui/link";
import { cookies } from "next/headers";

export default async function ViewReports() {

    const currMonth = new Date().getMonth() + 1;
    const currYear = new Date().getFullYear();
    const email = cookies().get('email')?.value!;
    return (
        <Card>
            <div className="flex flex-col gap-2 items-center">
            <h1>View Reports</h1>
            <div className="grid grid-cols-3 gap-2">
                <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={`/dashboard/reports/monthly-expenses?month=${currMonth}&year=${currYear}`}>Monthly Expenses (Pie Chart)</Link>
                <Link className=" bg-blue-700 rounded p-2 text-gray-50" href="/dashboard/reports/income-vs-expenses">Income & Expenses vs. Time (Bar Chart)</Link>
                <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={`/dashboard/reports/net-worth-report`}>Net Worth vs. Time (Bar Chart)</Link>
            </div>
            </div>
        </Card>
    )
}