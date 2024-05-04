import Card from "@/app/ui/card"
import PieChart from '@/app/dashboard/reports/monthly-expenses/pie-chart'
import {getNetWorths} from "@/app/lib/data/get-data"
import { cookies } from "next/headers"
import getTimeInfo from "@/app/lib/util/time-info"
import { BudgetPageData, Category, MonthlyBudget } from "@/app/lib/data/definitions"
import { BarChart } from "./bar-chart"
import TimeInfo from "@/app/lib/util/time-info"
import type { IncomeAndExpenses } from "@/app/lib/data/definitions"
import { log } from "console"

export default async function NetWorthReportPage() {

    // months: get the current month and the previous two months
    const email = cookies().get('email')?.value!;
    const today = new Date();
    const timeInfo = getTimeInfo(today.getMonth() + 1, today.getFullYear());
    const currMonthName = timeInfo.monthString;
    const netWorthReports = await getNetWorths(email)
    console.log(netWorthReports);

    // income will equal the sum of all transactions in the income category where transaction amount is greater than 0


    const months = netWorthReports.map((entry)=>{
        const date = new Date()
        log(entry.month - 1)
        date.setMonth(entry.month - 1)
        return date.toLocaleString('default', {month: "long"})
    })
    log(months)
    const netWorths = netWorthReports.map((entry)=>entry.netWorth/100)

    return (
        <Card>
            <BarChart months={months} netWorths={netWorths} />
        </Card>
    )
}