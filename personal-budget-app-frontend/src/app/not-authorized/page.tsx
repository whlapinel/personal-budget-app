import { Link } from "../ui/link";


export default function NotAuthorizedPage() {
    return (
        <div className="flex flex-col justify-center items-center min-h-[80vh] gap-4">
            <h1 className="text-4xl">Not Authorized</h1>
            <div className="flex gap-3">
                <Link className=" bg-blue-700 rounded p-2 text-gray-50 w-24 text-center" href="/sign-in">Sign In</Link>
                <Link className=" bg-blue-700 rounded p-2 text-gray-50 w-24 text-center" href="/sign-up">Sign Up</Link>
            </div>
        </div>
    );

}