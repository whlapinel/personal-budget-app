'use client';

import Card from "@/app/ui/card"
import type { Transaction } from '@/app/lib/definitions'
import type { Account } from '@/app/lib/definitions'
import { useSearchParams } from "next/navigation";
import {Link} from '@/app/ui/link';
import {transactions, accounts} from '@/app/lib/data/dummydata';


const totalBalance = accounts.reduce((acc, account) => acc + account.balance, 0);

export default function RegisterPage() {

  const searchParams = useSearchParams();
  const filter = searchParams.get('filter');
  console.log(filter);
  console.log(transactions);
  
  let filteredTransactions: Transaction[] = [];


  if (filter === null || filter === 'all') {
    filteredTransactions = transactions;
  } else {
    filteredTransactions = transactions.filter((transaction) => transaction.account === filter);
  }

  return (
    <>
      <Card className='bg-amber-200'>
        <div className={'flex gap-4'}>
          <h1 className="text-2xl font-bold text-gray-900">Filter by Account</h1>
          {accounts.map((account) => (
            <div key={account.id} className='flex gap-2'>
              <Link className=" bg-blue-700 rounded p-2 text-gray-50" href={`/dashboard?view=register&filter=${account.name}`} color='blue'>{account.name} {account.balance}</Link>
            </div>
          ))}
          <Link className={filter === 'all'?" bg-blue-700 rounded p-2 text-gray-50":"bg-slate-600 text-gray-50 rounded p-2"} href={`/dashboard?view=register&filter=all`}>All ({totalBalance})</Link>
        </div>
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
            {filteredTransactions.map((transaction) => (
              <tr key={transaction.id}>
                <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                  {transaction.account}
                </td>
                <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                  {transaction.date.toDateString()}
                </td>
                <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                  {transaction.payee}
                </td>
                <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                  {typeof transaction.category === 'string' ? transaction.category : 'split'}
                </td>
                <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                  {transaction.memo}
                </td>
                <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                  {transaction.amount < 0 ? transaction.amount : null}
                </td>
                <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                  {transaction.amount > 0 ? transaction.amount : null}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </Card>
    </>

  )
}
