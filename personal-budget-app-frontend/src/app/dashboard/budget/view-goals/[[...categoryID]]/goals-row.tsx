import convertToDollars from "@/app/lib/cents-to-dollars";
import { Goal } from "@/app/lib/data/definitions";


export default async function GoalsRow({ goal }: { goal: Goal }) {

    const parsedDate = new Date(goal.targetDate);

    return (
        <tr>
            <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                {goal.name}
            </td>
            <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                {convertToDollars(goal.amount)}
            </td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                {parsedDate.toLocaleDateString()}
            </td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                {goal.periodicity}
            </td>
        </tr>
    )
}
