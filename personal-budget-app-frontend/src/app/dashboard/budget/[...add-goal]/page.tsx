import AddGoalForm from "./add-goal-form";
import { getCategoryByID } from "@/app/lib/data/get-data";
import { Category } from "@/app/lib/data/definitions";

export default async function AddGoalPage({params}:{params: any}) {

    const categoryID = params['add-goal'][1];
    console.log(params['add-goal'][1]);
    

    const category: Category = await getCategoryByID(categoryID);
    console.log('AddGoalPage category:', category);
    

    return (
        <AddGoalForm category={category}/>
    )

}