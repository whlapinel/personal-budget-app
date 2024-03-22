import { Goal } from "@/app/lib/data/definitions";

export default function GoalsPage({goals}: {goals: Goal[]}) {
    // this page will display a list of goals for a given category

    const goalList = goals.map((goal: Goal) => {
        return (
            <tr key={goal.id}>
                <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                    {goal.name}
                </td>
                <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{convertToDollars(goal.amount)} { }</td>
            </tr>
        )
    });



    return (
        <>
        <p>Goals Page</p>
        </>
    )
}