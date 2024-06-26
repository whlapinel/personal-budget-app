import AddTransactionForm from '@/app/dashboard/transactions/add-transaction/add-transaction-form';
import { getAccounts, getCategories } from '@/app/lib/data/get-data';
import { cookies } from 'next/headers';

export default async function AddTransactionPage() {
  const email = cookies().get('email')?.value!;
  const accounts = await getAccounts(email);
  const categories = await getCategories(email);
  console.log("AddTransactionPage getAccounts(): ", accounts)

  return (
    <AddTransactionForm accounts={accounts} categories={categories}/>
  )
}
