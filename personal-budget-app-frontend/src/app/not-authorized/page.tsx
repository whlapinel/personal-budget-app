import { Link } from "../ui/link";


export default function NotAuthorizedPage() {
    return (
        <div>
            <h1>Not Authorized</h1>
            <p>
                You are not authorized to view this page. Please sign in or create an account.
            </p>
            <Link href="/sign-in">Sign In</Link>
        </div>
    );

}