import { Account, Transaction } from "@/app/lib/data/definitions";
import { getTransactions } from "@/app/lib/data/get-data";

export default async function TransactionList({account}: {account: Account}) {
    const transactions = await getTransactions(account.id);
    return (
      <table className="min-w-full divide-y divide-gray-300">
        <thead>
          <tr>
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
          {transactions.map((transaction: Transaction) => (
            <tr key={transaction.id}>
              <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                {transaction.date.toDateString()}
              </td>
              <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                {transaction.payee}
              </td>
              <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                {typeof transaction.categoryID === 'string' ? transaction.categoryID : 'split'}
              </td>
              <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                {transaction.memo}
              </td>
              <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                {transaction.amount < 0 ? Math.abs(transaction.amount) : null}
              </td>
              <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                {transaction.amount > 0 ? transaction.amount : null}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
)

}