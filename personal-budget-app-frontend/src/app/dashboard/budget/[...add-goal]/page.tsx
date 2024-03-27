import AddGoalForm from "./add-goal-form";
import { getCategories, getCategoryByID } from "@/app/lib/data/get-data";
import { Category } from "@/app/lib/data/definitions";
import { cookies } from "next/headers";

export default async function AddGoalPage({params}:{params: any}) {

    
    const categoryID = params['add-goal'][1];
    console.log(params['add-goal'][1]);
    const email = cookies().get('email')?.value!;
    const categories: Category[] = await getCategories(email);
    const category = categories.find((category) => {
        return category.id === Number(categoryID);
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