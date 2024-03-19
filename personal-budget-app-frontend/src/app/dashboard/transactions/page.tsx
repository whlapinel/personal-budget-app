import Card from "@/app/ui/card"
import type { Transaction, Account } from '@/app/lib/data/definitions'
import { Link } from '@/app/ui/link';
import { getTransactions } from '@/app/lib/data/get-data';
import { getAccounts } from '@/app/lib/data/get-data';
import { cookies } from 'next/headers';
import TransactionList from "./transaction-list";
import { select } from "d3";

export default async function TransactionsPage(searchParams: any) {
  const email = cookies().get('email')?.value!;
  const accounts = await getAccounts(email);
  console.log('accounts: ', accounts);
  console.log('searchParams: ', searchParams);
  console.log('searchParams.searchParams.selectedAccountID: ', searchParams.searchParams.selectedAccountID);
  const accountsList = accounts?.map((account) => {
    return (
      <Link className={" bg-blue-700 rounded p-2 text-gray-50"} key={account.id} href={`/dashboard/transactions?selectedAccountID=${account.id}`}>{account.name}</Link>
    )
  })
  const totalBalance = accounts?.reduce((acc, account) => acc + account.balance, 0);
  const selectedAccount = accounts.find((account) => account.id === Number(searchParams.searchParams.selectedAccountID));
  console.log("selectedAccount: ", selectedAccount);
  let filteredTransactions: Transaction[] = [];

  return (
    <>
      <Card className='bg-amber-200'>
        <div className={'flex flex-col gap-4'}>
          <div className="flex gap-4">
            <h2 className="text-2xl font-bold text-gray-900">Select Account</h2>
            {accountsList}
          </div>
          <div className="flex gap-4">
            <Link className=" bg-blue-700 rounded p-2 text-gray-50" href='/dashboard/transactions/add-transaction'>Add Transaction</Link>
          </div>
          {selectedAccount !== undefined ?
            <TransactionList account={selectedAccount} />
            :
            null
          }
        </div>
      </Card>
    </>
  )
}