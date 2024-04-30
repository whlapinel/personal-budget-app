import { Account, Transaction } from "@/app/lib/data/definitions";
import { getTransactions } from "@/app/lib/data/get-data";
import {convertToDollars} from "@/app/lib/util/cents-to-dollars";

export default async function TransactionList({account, transactions}: {account: Account, transactions: Transaction[]}) {
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
                {transaction.categoryName ? transaction.categoryName : ''}
              </td>
              <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                {transaction.memo}
              </td>
              <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                {transaction.amount < 0 ? `${convertToDollars(Math.abs(transaction.amount))}` : null}
              </td>
              <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                {transaction.amount > 0 ? `${convertToDollars(transaction.amount)}` : null}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
)

}