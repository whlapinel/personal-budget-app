
import Card from "@/app/ui/card"
import PieChart from '@/app/dashboard/reports/monthly-expenses/pie-chart'
import { getCategories, getMonthlyBudgets } from "@/app/lib/data/get-data"
import { cookies } from "next/headers"
import getTimeInfo from "@/app/lib/util/time-info"
import { BudgetPageData, Category, MonthlyBudget } from "@/app/lib/data/definitions"


export default async function ReportsPage({ searchParams }: { searchParams: any}) {
  const monthParam = Number(searchParams.month); 
  const yearParam = Number(searchParams.year);
  const email = cookies().get('email')?.value!;
  const timeInfo = getTimeInfo(monthParam, yearParam)
  const monthlyBudgets = await getMonthlyBudgets(email, monthParam, yearParam)
  console.log('monthlyBudgets: ', monthlyBudgets)

  const thisMonthsBudgets: MonthlyBudget[] = monthlyBudgets?.filter((budget) => budget.month === monthParam && budget.year === yearParam)
  const categories: Category[] = await getCategories(email)

  const budgetData = thisMonthsBudgets?.map((budget) => {
    return({
      name: categories?.find((category) => category.id === budget.categoryID)?.name,
      amount: budget.spent
    }
    )
  })
  console.log('budgetData: ', budgetData)

  const names = budgetData?.map((budget: any)=> budget.name)
  const amounts = budgetData?.map((budget: any)=> (budget.amount/100))
  console.log("names: ", names)
  console.log("amounts: ", amounts)


  return (
    <Card className='bg-white-200'>
        <PieChart names={names} amounts={amounts} />
    </Card>
  )
}