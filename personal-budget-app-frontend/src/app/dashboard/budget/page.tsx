
import Card from "@/app/ui/card"
import { Link } from "@/app/ui/link";
import { getBudgetPageData } from "@/app/lib/data/get-data";
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
  const nextMonth = viewedMonth === 11 ? 0 : viewedMonth + 1;
  const prevMonth = viewedMonth === 0 ? 11 : viewedMonth - 1;
  const nextYear = viewedMonth === 11 ? viewedYear + 1 : viewedYear;
  const prevYear = viewedMonth === 0 ? viewedYear - 1 : viewedYear;
  console.log("BudgetPage() month: ", viewedMonth)
  console.log("BudgetPage() year: ", viewedYear)
  console.log("BudgetPage() viewedDate: ", viewedDate)
  const monthString = viewedDate.toLocaleString('default', { month: 'long' });
  const email = cookies().get('email')?.value!;
  console.log("BudgetPage() email: ", email);

  const budgetPageData = await getBudgetPageData(email, viewedMonth, viewedYear);

  const categoryRows = budgetPageData?.categoryRows;


  return (
    <>
      <Card className='bg-amber-200'>
        <h1 className="text-2xl font-bold text-gray-900">{monthString}{" "}{viewedYear} Budget</h1>
        <h2 className="text-xl font-bold text-gray-900">Unassigned: {convertToDollars(budgetPageData.totalUnassigned)}</h2>
        <Link href={`/dashboard/budget?month=${prevMonth}&year=${prevYear}`} className="bg-blue-700 rounded p-2 text-gray-50">Previous Month</Link>
        <Link href={`/dashboard/budget?month=${nextMonth}&year=${nextYear}`} className="bg-blue-700 rounded p-2 text-gray-50">Next Month</Link>
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

            {categoryRows ? categoryRows?.map((categoryData: CategoryData) => {
              return (
                <CategoryRow key={categoryData.categoryID} data={categoryData} month={viewedMonth} year={viewedYear} />
              )
            }
            ):<p>No categories found</p>}
          </tbody>
        </table>
      </Card>
    </>
  )
}


