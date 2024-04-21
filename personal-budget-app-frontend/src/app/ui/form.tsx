import { SubmitButton } from "./submit-button"

export type FormHiddenInfo = {
    name: string,
    value: string | number
}


export default function Form({ title, children, formAction, state, hiddenInfo }: { title: string, children: React.ReactNode, formAction: (payload: FormData) => void, state: any, hiddenInfo?: FormHiddenInfo[] }) {
    return (
        <div className="bg-white-200 p-4 rounded-lg shadow-md flex flex-col justify-center items-center gap-2">
            <h2 className="text-2xl font-semibold">{title}</h2>
            <form action={formAction} className="flex flex-col">
                <div className="grid grid-cols-2 gap-1 text-right">

                    {children}
                </div>
                {hiddenInfo?.map(info => (
                    <input key={info.name} type="hidden" name={info.name} value={info.value} />
                ))}
                <SubmitButton className=" w-36">Add</SubmitButton>
            </form>
            <p>{state.message}</p>
        </div>
    )
}