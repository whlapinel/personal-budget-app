import { Category } from "@/app/lib/data/definitions";
import { useFormState } from "react-dom";
import addGoalAction from "./actions/add-goal-action";


const initialState = {
    message: null,
    goal: null
}
export default function AddGoalForm(category: Category) {
    const [state, formAction] = useFormState(addGoalAction, initialState)
    
    return (
        <input>Hello</input>

    )
}