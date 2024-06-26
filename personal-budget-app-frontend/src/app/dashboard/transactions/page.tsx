import Card from "@/app/ui/card"
import type { Transaction, Account } from '@/app/lib/data/definitions'
import { Link } from '@/app/ui/link';
import { getTransactions } from '@/app/lib/data/get-data';
import { getAccounts } from '@/app/lib/data/get-data';
import { cookies } from 'next/headers';
import TransactionList from "./transaction-list";
import { select } from "d3";
import {convertToDollars} from "@/app/lib/util/cents-to-dollars";

export default async function TransactionsPage({searchParams}: {searchParams: any}) {
  const email = cookies().get('email')?.value!;
  const accounts = await getAccounts(email);
  const transactions = await getTransactions(email);
  console.log('accounts: ', accounts);
  console.log('searchParams: ', searchParams);
  console.log('searchParams.selectedAccountID: ', searchParams.selectedAccountID);
  const accountsList = accounts?.map((account) => {
    return (
      <Link className={" bg-blue-700 rounded p-2 text-gray-50"} key={account.id} href={`/dashboard/transactions?selectedAccountID=${account.id}`}>{account.name} {convertToDollars(account.balance)}</Link>
    )
  })
  const totalBalance = accounts?.reduce((acc, account) => acc + account.balance, 0);
  const selectedAccount = accounts?.find((account) => account.id === Number(searchParams.selectedAccountID));
  console.log("selectedAccount: ", selectedAccount);
  let filteredTransactions: Transaction[] = [];

  return (
    <>
      <Card>
        <div className={'flex flex-col gap-4'}>
          <div className="flex gap-4">
            <h2 className="text-2xl font-bold text-gray-900">Select Account</h2>
            {accountsList}
          </div>
          <div className="flex gap-4">
            <Link className=" bg-blue-700 rounded p-2 text-gray-50" href='/dashboard/transactions/add-transaction'>Add Transaction</Link>
          </div>
          {selectedAccount !== undefined ?
            <TransactionList account={selectedAccount} transactions={transactions.filter((transaction)=> transaction.accountID === selectedAccount.id)}/>
            :
            "No account selected"
          }
        </div>
      </Card>
    </>
  )
}