import { getAssignments, getCategories, getMonthlyBudgets } from "@/app/lib/data/get-data";
import EditAssignmentForm from "./edit-assignment-form";
import { cookies } from "next/headers";
import { Assignment, MonthlyBudget } from "@/app/lib/data/definitions";

export default async function EditAssignmentPage({params, searchParams}: {params: any, searchParams: any}) {

    const monthParam = Number(searchParams.month);
    const yearParam = Number(searchParams.year);
  
    if (!params.categoryID) {
        console.log('AddGoalPage missing categoryID');
        return null;
    }
    const email = cookies().get('email')?.value!;

    const categoryID = Number(params.categoryID[0]);
    console.log("categoryID: ", categoryID);
    if (isNaN(categoryID)) {
        console.log('AddGoalPage invalid categoryID');
        return null;
    }
    const budgets: MonthlyBudget[] = await getMonthlyBudgets(email, monthParam, yearParam);

    const currBudget = budgets?.find((assignment) => {
        return assignment.month === monthParam && assignment.year === yearParam;
    });
    const currAssigned = currBudget ? currBudget.assigned : null;

    const categories = await getCategories(email);
    const category = categories.find((category) => {
        return category.id === categoryID;
    })!;



    return (
        <EditAssignmentForm category={category} month={monthParam} year={yearParam} currAssignmentAmount={currAssigned || 0}/>
    )
}
