
import Card from "@/app/ui/card"
import { Link } from "@/app/ui/link";
import { getBudgetPageData } from "@/app/lib/data/get-data";
import { CategoryData} from "@/app/lib/data/definitions";
import { cookies } from "next/headers";
import CategoryRow from "./category-row";
import {convertToDollars} from "@/app/lib/util/cents-to-dollars";
import { TimeInfo } from "@/app/lib/data/definitions";
import getTimeInfo from "@/app/lib/util/time-info";

export default async function BudgetPage({ searchParams }: { searchParams: any }) {
  const monthParam = Number(searchParams.month); // NOT ZERO INDEXED
  const yearParam = Number(searchParams.year);

  const timeInfo: TimeInfo = getTimeInfo(monthParam, yearParam); 
  const email = cookies().get('email')?.value!;
  console.log("BudgetPage() email: ", email);
  
  const budgetPageData = await getBudgetPageData(email, timeInfo.viewedMonth, timeInfo.viewedYear);

  const categoryRows = budgetPageData?.categoryRows;


  return (
    <>
      <Card>
        <h1 className="text-2xl font-bold text-gray-900">{timeInfo.monthString}{" "}{timeInfo.viewedYear} Budget</h1>
        <h2 className="text-xl font-bold text-gray-900">Unassigned: {convertToDollars(budgetPageData.totalUnassigned)}</h2>
        <div className="flex gap-2">
        <Link href={`/dashboard/budget?month=${timeInfo.prevMonth}&year=${timeInfo.prevYear}`} className="bg-blue-700 rounded p-2 text-gray-50">Previous Month</Link>
        <Link href={`/dashboard/budget?month=${timeInfo.nextMonth}&year=${timeInfo.nextYear}`} className="bg-blue-700 rounded p-2 text-gray-50">Next Month</Link>
        <Link className=" bg-blue-700 rounded p-2 text-gray-50" href='/dashboard/budget/add-category'>Add Budget Category</Link>
        </div>
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
                <CategoryRow key={categoryData.categoryID} data={categoryData} month={timeInfo.viewedMonth} year={timeInfo.viewedYear} />
              )
            }
            ):<p>No categories found</p>}
          </tbody>
        </table>
      </Card>
    </>
  )
}


