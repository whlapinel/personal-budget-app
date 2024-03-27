import { getCategories } from "@/app/lib/data/get-data";
import { cookies } from "next/headers";
import { Goal } from "@/app/lib/data/definitions";
import Card from "@/app/ui/card";
import { Link } from "@/app/ui/link";
import GoalsRow from "./goals-row";


export default async function ViewGoalsPage({ params, searchParams }: { params: any, searchParams: any }) {
  const categoryID = Number(params.categoryID[0]);
  const month = Number(searchParams.month);
  const year = Number(searchParams.year);
  const viewedMonth = new Date(year, month);
  const email = cookies().get('email')?.value!;

  const categories = await getCategories(email);
  const category = categories.find((category) => {
    return category.id === categoryID;
  })!;
  console.log('ViewGoalsPage category:', category);
  const goals = category.goals || [];
  console.log('ViewGoalsPage goals:', goals);



  return (
    <>
      <Card className='bg-amber-200'>
        <h1 className="text-2xl font-bold text-gray-900">Goals for {viewedMonth.toLocaleString('default', { month: 'long' })}{' '}{year}</h1>
        <table className="min-w-full divide-y divide-gray-300">
          <thead>
            <tr>
              <th scope="col" className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0">
                Name
              </th>
              <th scope="col" className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0">
                Amount
              </th>
              <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                Date
              </th>
              <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                Periodicity
              </th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200">
            {goals?.map((goal) => {

              return (
                <GoalsRow key={goal.id} goal={goal} />
              )
            }
            )}
          </tbody>
        </table>
      </Card>
    </>
  )
}