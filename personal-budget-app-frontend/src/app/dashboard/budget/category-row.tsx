import convertToDollars from "@/app/lib/cents-to-dollars";
import { Category, Goal } from "@/app/lib/data/definitions";
import { getGoal } from "@/app/lib/data/get-data";
import { Goldman } from "next/font/google";

export default async function CategoryRow({ category, month }: { category: Category, month: number}) {
    // get goals for this category
    // this API call also needs to include the month and year
    // const goal: Goal = await getGoal(category.id);
    // console.log('CategoryRow goal:', goal);

    const goals: Goal[] = category.goals?.filter((goal: Goal) => {

        return new Date(Date.parse(goal.targetDate)).getMonth() === month;
    }) || [];

    const totalNeeded = goals.reduce((acc: number, goal: Goal) => {
        return (acc + goal.amount);
    }, 0);




    const goalName = goals.length > 0 ? goals[0].name : 'No goal assigned';





    return (
        <tr key={category.id}>
            <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                {category.name}
            </td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{convertToDollars(totalNeeded)} for {goalName}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{'placeholder assigned'}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{'placeholder available'}</td>
            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{'placeholder spent'}</td>
        </tr>
    )
}
