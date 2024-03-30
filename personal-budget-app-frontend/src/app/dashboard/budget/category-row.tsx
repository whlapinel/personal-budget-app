import convertToDollars from "@/app/lib/cents-to-dollars";
import { Category, Goal, Transaction, Assignment, CategoryData} from "@/app/lib/data/definitions";
import { getAccounts, getAssignments, getGoals, getCategories, getTransactions } from "@/app/lib/data/get-data";
import { Link } from "@/app/ui/link";
import { IdentificationIcon } from "@heroicons/react/24/outline";

export default async function CategoryRow({data, month, year}: {data: CategoryData, month: number, year: number}) {

    console.log('CategoryRow data:', data)
    return (
        <tr>
            <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                {data.categoryName || 'Category not found'}
            </td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{convertToDollars(data.goalsSum)}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                <Link href={`/dashboard/budget/edit-assignment/${data.categoryID}?month=${month}&year=${year}`}>{convertToDollars(data.assigned)}</Link></td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{convertToDollars(data.available)}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{convertToDollars(data.spent)}</td>
            <td><Link href={`/dashboard/budget/add-goal/${data.categoryID}`}>Add Goal</Link>
            </td>
        </tr>
    )
}
