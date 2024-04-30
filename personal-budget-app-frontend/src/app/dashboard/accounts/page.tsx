import { backendUrls } from "@/app/constants/backend-urls";
import Link from "next/link";
import { cookies } from "next/headers";
import { Account } from "@/app/lib/data/definitions";
import Card from "@/app/ui/card";
import { getAccounts } from "@/app/lib/data/get-data";
import { convertToDollars } from "@/app/lib/util/cents-to-dollars";

export default async function AccountsPage() {

    const email = cookies().get('email')?.value!;
    const accounts = await getAccounts(email);

    const total = accounts?.reduce((acc, account) => {
        return acc + account.balance;
    }, 0)



    return (
        <Card>
        <h1 className="text-2xl font-bold text-gray-900">Accounts</h1>
        <h2 className="text-xl font-bold text-gray-900">Total: {convertToDollars(total)}</h2>
        <div className="flex gap-2">
        <Link className=" bg-blue-700 rounded p-2 text-gray-50" href='/dashboard/accounts/add-account'>Add Account</Link>
        </div>
            <table className="min-w-full divide-y divide-gray-300">
          <thead>
            <tr>
              <th scope="col" className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0">
                Name
              </th>
              <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                Type
              </th>
              <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                Balance
              </th>
            </tr>
          </thead>
            <tbody className="divide-y divide-gray-200">
                {accounts?.map((account: Account) => {
                return (
                    <tr key={account.id}>
                    <td className="whitespace-nowrap">
                        <div className="text-sm text-gray-900">{account.name}</div>
                    </td>
                    <td className="px-3 py-3.5 whitespace-nowrap">
                        <div className="text-sm text-gray-900">{account.type}</div>
                    </td>
                    <td className="px-3 py-3.5 whitespace-nowrap">
                        <div className="text-sm text-gray-900">{convertToDollars(account.balance)}</div>
                    </td>
                    </tr>
                )
                })}
            </tbody>

          </table>

        </Card>
    )

}