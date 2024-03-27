import AddGoalForm from "./add-goal-form";
import { getCategories, getCategoryByID } from "@/app/lib/data/get-data";
import { Category } from "@/app/lib/data/definitions";
import { cookies } from "next/headers";

export default async function AddGoalPage({params}:{params: any}) {

    if (!params.categoryID) {
        console.log('AddGoalPage missing categoryID');
        return null;
    }

    const categoryID = Number(params.categoryID[0]);
    console.log("categoryID: ", categoryID);
    if (isNaN(categoryID)) {
        console.log('AddGoalPage invalid categoryID');
        return null;
    }

    const email = cookies().get('email')?.value!;
    const categories: Category[] = await getCategories(email);
    const category = categories.find((category) => {
        return category.id === categoryID;
    })!;
    if (!category) {
        console.log('AddGoalPage category not found');
        return null;
    }
    console.log('AddGoalPage category:', category);
    
    return (
        <AddGoalForm category={category}/>
    )

}