
import Card from "@/app/ui/card"
import { Link } from "@/app/ui/link";
import { Dialog } from "@/app/ui/dialog";
import { getUser, getCategories } from "@/app/lib/data/get-data";
import { BudgetCategory } from "@/app/lib/data/definitions";

export default async function BudgetPage() {

  const userID = 1;
  const today = new Date();
  const month = today.toLocaleString('default', { month: 'long' });

  let categories: BudgetCategory[] = [];

  try {
    categories = await getCategories();
    console.log(categories);
  } catch (err) {
    console.error(err);
    return (
      <p>Error getting categories"</p>
    )
  }


  return (
    <>
      <Card className='bg-amber-200'>
        <h1 className="text-2xl font-bold text-gray-900">{month} Budget</h1>
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
              <tr key={category.id}>
                <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                  {category.name}
                </td>
                <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{'placeholder'}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </Card>
    </>
  )
}


