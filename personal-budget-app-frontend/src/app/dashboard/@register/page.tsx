import Card from "@/app/ui/card"

const transactions = [
  { account: 'Checking', id: 1, date: '2021-01-01', payee: 'Target', category: 'Shopping', memo: 'Groceries', outflow: 100.00, inflow: 0.00 },
  { account: 'Checking', id: 2, date: '2021-01-01', payee: 'Target', category: 'Shopping', memo: 'Groceries', outflow: 100.00, inflow: 0.00 },
  { account: 'Checking', id: 3, date: '2021-01-01', payee: 'Target', category: 'Shopping', memo: 'Groceries', outflow: 100.00, inflow: 0.00 },
  { account: 'Checking', id: 4, date: '2021-01-01', payee: 'Target', category: 'Shopping', memo: 'Groceries', outflow: 100.00, inflow: 0.00 },
  { account: 'Checking', id: 5, date: '2021-01-01', payee: 'Target', category: 'Shopping', memo: 'Groceries', outflow: 100.00, inflow: 0.00 },
  { account: 'Checking', id: 6, date: '2021-01-01', payee: 'Target', category: 'Shopping', memo: 'Groceries', outflow: 100.00, inflow: 0.00 },
  { account: 'Checking', id: 7, date: '2021-01-01', payee: 'Target', category: 'Shopping', memo: 'Groceries', outflow: 100.00, inflow: 0.00 },
  { account: 'Checking', id: 8, date: '2021-01-01', payee: 'Target', category: 'Shopping', memo: 'Groceries', outflow: 100.00, inflow: 0.00 },
]




export default function RegisterPage() {
  return (
    <>
      <Card className='bg-amber-200'>
        <div className="mt-8 flow-root">
          <div className="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
            <div className="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
              <table className="min-w-full divide-y divide-gray-300">
                <thead>
                  <tr>
                    <th scope="col" className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0">
                      Account
                    </th>
                    <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                      Date
                    </th>
                    <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                      Payee
                    </th>
                    <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                      Category
                    </th>
                    <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                      Memo
                    </th>
                    <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                      Outflow
                    </th>
                    <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                      Inflow
                    </th>
                    <th scope="col" className="relative py-3.5 pl-3 pr-4 sm:pr-0">
                      <span className="sr-only">Edit</span>
                    </th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-200">
                  {transactions.map((transaction) => (
                    <tr key={transaction.id}>
                      <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                        {transaction.account}
                      </td>
                      <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{transaction.date}</td>
                      <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{transaction.category}</td>
                      <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{transaction.memo}</td>
                      <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{transaction.outflow}</td>
                      <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{transaction.inflow}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </Card>
    </>

  )
}
