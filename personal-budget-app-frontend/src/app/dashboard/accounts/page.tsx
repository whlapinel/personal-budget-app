import { backendUrls } from "@/app/constants/backend-urls";
import Link from "next/link";
import { cookies } from "next/headers";
import { Account } from "@/app/lib/data/definitions";

export default async function AccountsPage() {

    const email = cookies().get('email')?.value!;
    const accounts = await getAccounts(email);

    async function getAccounts(email: string) {
        let accounts: Account[] = [];
        try {
            const response = await fetch(`${backendUrls.accounts}/${email}`, {
                method: 'GET',
                headers: {
                    'API_KEY': process.env.API_KEY!,
                },        
            });
            accounts = await response.json();
        } catch (err) {
            console.log(err);
        }
        return accounts;
    }


    return (
        <div>
            <h1>Accounts</h1>
            {accounts?.map((account: Account) => (<p key={account.id}>{account.name}</p>))}

            <Link color="blue" href='/dashboard/accounts/add-account'>Add Account</Link>
        </div>
    )

}