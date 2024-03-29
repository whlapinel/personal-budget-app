
import Card from "@/app/ui/card"
import { Link } from "@/app/ui/link";
import { getTransactions, getCategories, getAccounts, getAssignmentsByEmail, getMonthlyBudget } from "@/app/lib/data/get-data";
import { Category, CategoryData, Goal, MonthlyBudget, Transaction } from "@/app/lib/data/definitions";
import { cookies } from "next/headers";
import CategoryRow from "./category-row";
import convertToDollars from "@/app/lib/cents-to-dollars";
import { backendUrls } from "@/app/constants/backend-urls";

export default async function BudgetPage({ searchParams }: { searchParams: any }) {
  const monthParam = Number(searchParams.month);
  const yearParam = Number(searchParams.year);
  const today = new Date();
  const viewedMonth = monthParam;
  const viewedYear = yearParam || today.getFullYear();
  const viewedDate = new Date(viewedYear, viewedMonth);
  console.log("BudgetPage() month: ", viewedMonth)
  console.log("BudgetPage() year: ", viewedYear)
  console.log("BudgetPage() viewedDate: ", viewedDate)
  const monthString = viewedDate.toLocaleString('default', { month: 'long' });
  const email = cookies().get('email')?.value!;
  console.log("BudgetPage() email: ", email);
  const monthlyBudget: MonthlyBudget = await getMonthlyBudget(email, viewedMonth, viewedYear);
  console.log("monthlyBudget: ", monthlyBudget);
  console.log("monthlyBudget.categoryArray: ", monthlyBudget.categoryArray);

  return (
    <>
      <Card className='bg-amber-200'>
        <h1 className="text-2xl font-bold text-gray-900">{monthString} Budget</h1>
        <h2 className="text-xl font-bold text-gray-900">Unassigned: {convertToDollars(monthlyBudget.unassigned)}</h2>
        <Link href={`/dashboard/budget?month=${viewedMonth - 1}&year=${viewedYear}`} className="bg-blue-700 rounded p-2 text-gray-50">Previous Month</Link>
        <Link href={`/dashboard/budget?month=${viewedMonth + 1}&year=${viewedYear}`} className="bg-blue-700 rounded p-2 text-gray-50">Next Month</Link>
        <Link className=" bg-blue-700 rounded p-2 text-gray-50" href='/dashboard/budget/add-category'>Add Budget Category</Link>
        <table className="min-w-full divide-y divide-gray-300">
          <thead>
            <tr>
              <th scope="col" className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0">
                Name
              </th>
              <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                Needed to meet goal
              </th>
              <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                Assigned
              </th>
              <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                Available
              </th>
              <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                Spent
              </th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200">
            {monthlyBudget.categoryArray?.map((categoryData: CategoryData) => {

              return (
                <CategoryRow key={categoryData.id} categoryData={categoryData} month={viewedMonth} year={viewedYear} />
              )
            }
            )}
          </tbody>
        </table>
      </Card>
    </>
  )
}


