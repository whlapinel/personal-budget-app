
import Card from "@/app/ui/card"
import { Link } from "@/app/ui/link";
import { getCategories } from "@/app/lib/data/get-data";
import { getGoals } from "@/app/lib/data/get-data";
import { Category, Goal } from "@/app/lib/data/definitions";
import { cookies } from "next/headers";
import CategoryRow from "./category-row";

export default async function BudgetPage({searchParams}: {searchParams: any}) {
  const monthParam = Number(searchParams.month);
  const yearParam = Number(searchParams.year);
  const today = new Date();
  const viewedMonth = monthParam || today.getMonth();
  const viewedYear = yearParam || today.getFullYear();
  const viewedDate = new Date(viewedYear, viewedMonth);
  console.log("BudgetPage() month: ", viewedMonth)
  console.log("BudgetPage() year: ", viewedYear)
  console.log("BudgetPage() viewedDate: ", viewedDate)
  const monthString = viewedDate.toLocaleString('default', { month: 'long' });
  const email = cookies().get('email')?.value!;
  console.log("BudgetPage() email: ", email);
  const categories: Category[] = await getCategories(email);
  console.log(categories);


  return (
    <>
      <Card className='bg-amber-200'>
        <h1 className="text-2xl font-bold text-gray-900">{monthString} Budget</h1>
        <Link className=" bg-blue-700 rounded p-2 text-gray-50" href='/dashboard/budget/add-category'>Add Budget Item</Link>
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
              {categories?.map((category) => (
                <CategoryRow key={category.id} category={category} month={viewedMonth} />
              ))}
            </tbody>
        </table>
      </Card>
    </>
  )
}


