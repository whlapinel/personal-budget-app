import convertToDollars from "@/app/lib/cents-to-dollars";
import { Category, Goal, Transaction, Assignment, CategoryData } from "@/app/lib/data/definitions";
import { getAccounts, getAssignments } from "@/app/lib/data/get-data";
import { Link } from "@/app/ui/link";
import { IdentificationIcon } from "@heroicons/react/24/outline";

export default async function CategoryRow({ categoryData, month, year}: { categoryData: CategoryData, month: number, year: number}) {
    // get goals for this category
    // this API call also needs to include the month and year
    // const goal: Goal = await getGoal(category.id);
    // console.log('CategoryRow goal:', goal);
    const id = categoryData.id

    
    const goalsLink = categoryData.needed > 0 ? <Link href={`/dashboard/budget/view-goals/${IdentificationIcon}?month=${month}&year=${year}`}>Goals</Link> : null;
    return (
        <tr>
            <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                {categoryData.name}
            </td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{goalsLink}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                <Link href={`/dashboard/budget/edit-assignment/${id}?month=${month}&year=${year}`}>{convertToDollars(categoryData.assigned)}</Link></td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{'placeholder available'}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{convertToDollars(categoryData.spent)}</td>
            <td><Link href={`/dashboard/budget/add-goal/${id}`}>Add Goal</Link>
            </td>
        </tr>
    )
}
