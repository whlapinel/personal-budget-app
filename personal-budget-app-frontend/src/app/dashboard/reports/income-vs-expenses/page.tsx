import Card from "@/app/ui/card"
import PieChart from '@/app/dashboard/reports/monthly-expenses/pie-chart'
import { getCategories, getMonthlyBudgets } from "@/app/lib/data/get-data"
import { cookies } from "next/headers"
import getTimeInfo from "@/app/lib/util/time-info"
import { BudgetPageData, Category, MonthlyBudget } from "@/app/lib/data/definitions"
import { BarChart } from "./bar-chart"

export default async function IncomeVsExpensesPage(){

    return(

    <Card>
        <BarChart/>
    </Card>
    )
}