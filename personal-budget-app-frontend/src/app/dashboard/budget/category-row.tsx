import convertToDollars from "@/app/lib/cents-to-dollars";
import { Category, Goal, Transaction, Assignment } from "@/app/lib/data/definitions";
import { getAccounts, getAssignments } from "@/app/lib/data/get-data";
import { Link } from "@/app/ui/link";

export default async function CategoryRow({ category, month, year, transactions }: { category: Category, month: number, year: number, transactions: Transaction[] }) {
    // get goals for this category
    // this API call also needs to include the month and year
    // const goal: Goal = await getGoal(category.id);
    // console.log('CategoryRow goal:', goal);

    const assignments: Assignment[] = await getAssignments(category.id);
    console.log("assignments: ", assignments)
    
    
    const goals = category.goals || [];
    console.log("goals: ", goals)
    const goalsThisMonth: Goal[] = goals.filter((goal: Goal) => {
        return goal.targetDate?.getMonth() === month && goal.targetDate?.getFullYear() === year || goal.periodicity === 'monthly';
    });
    console.log("goalsThisMonth:", goalsThisMonth)

    const totalNeeded = goalsThisMonth.reduce((acc: number, goal: Goal) => {
        return (acc + goal.amount);
    }, 0);

    const assignedAmount = assignments?.find((assignment) => {
        return assignment.month === month && assignment.year === year;
    })?.amount || 0;

    const neededRemaining: number = totalNeeded - assignedAmount;
    const goalString = goalsThisMonth.length > 0 ? `${convertToDollars(neededRemaining)} for `: 'No goal assigned';
    
    const totalSpent: number = transactions.reduce((acc, transaction) => {
        return (acc + transaction.amount);
    }, 0);
    
    
    const overspent: boolean = totalSpent > assignedAmount;
    
    const goalsLink = goalsThisMonth.length > 0 ? <Link href={`/dashboard/budget/view-goals/${category.id}?month=${month}&year=${year}`}>Goals</Link> : null;
    return (
        <tr key={category.id}>
            <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                {category.name}
            </td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{goalString}{goalsLink}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                <Link href={`/dashboard/budget/edit-assignment/${category.id}?month=${month}&year=${year}`}>{convertToDollars(assignedAmount)}</Link></td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{'placeholder available'}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{convertToDollars(-1 * totalSpent)}</td>
            <td><Link href={`/dashboard/budget/add-goal/${category.id}`}>Add Goal</Link>
            </td>
        </tr>
    )
}
