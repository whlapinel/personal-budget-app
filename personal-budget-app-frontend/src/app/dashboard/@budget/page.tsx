
import Card from "@/app/ui/card"


const categories = [
  { name: 'Food', id: 1, amount: 0 },
  { name: 'Rent', id: 2, amount: 0 },
  { name: 'Utilities', id: 3, amount: 0 },
  { name: 'Entertainment', id: 4, amount: 0 },
  { name: 'Transportation', id: 5, amount: 0 },
  { name: 'Other', id: 6, amount: 0 },
]

export default function BudgetPage() {

  const today = new Date();
  const month = today.toLocaleString('default', { month: 'long' });



  return (
    <>
      <Card className='bg-amber-200'>
        <h1 className="text-2xl font-bold text-gray-900">{month} Budget</h1>
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
            {categories.map((category) => (
              <tr key={category.id}>
                <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                  {category.name}
                </td>
                <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{category.amount}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </Card>
    </>
  )
}


