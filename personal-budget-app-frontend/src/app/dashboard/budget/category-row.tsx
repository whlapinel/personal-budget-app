import convertToDollars from "@/app/lib/cents-to-dollars";
import { Category, Goal, Transaction } from "@/app/lib/data/definitions";
import { Link } from "@/app/ui/link";

export default async function CategoryRow({ category, month, transactions }: { category: Category, month: number, transactions: Transaction[]}) {
    // get goals for this category
    // this API call also needs to include the month and year
    // const goal: Goal = await getGoal(category.id);
    // console.log('CategoryRow goal:', goal);

    const goals = category.goals || [];

    const goalsThisMonth: Goal[] = goals.filter((goal: Goal) => {
        return new Date(Date.parse(goal.targetDate)).getMonth() === month;
    });
    console.log("goalsThisMonth:", goalsThisMonth)

    const totalNeeded = goalsThisMonth.reduce((acc: number, goal: Goal) => {
        return (acc + goal.amount);
    }, 0);

    const goalName = goalsThisMonth.length > 0 ? goalsThisMonth[0].name : 'No goal assigned';

    const totalSpent: number = transactions.reduce((acc, transaction) => {
        return (acc + transaction.amount);
    }, 0);

    return (
        <tr key={category.id}>
            <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                {category.name}
            </td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{convertToDollars(totalNeeded)} for {goalName}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{'placeholder assigned'}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{'placeholder available'}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{convertToDollars(totalSpent)}</td>
            <td><Link href={`/dashboard/budget/add-goal/${category.id}`}>Add Goal</Link>
            </td>
        </tr>
    )
}
