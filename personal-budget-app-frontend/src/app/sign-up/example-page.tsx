import { redirect } from "next/navigation";
import { getToken, signin, signOut } from "@/app/lib/data/auth";

export default async function Page() {
  const session = await getToken();
  return (
    <section>
      <form
        action={async (formData) => {
          "use server";
          await signin(formData);
        //   redirect("/");
        }}
      >
        <input type="email" name="email" placeholder="Email" />
        <br />
        <button type="submit">Login</button>
      </form>
      <form
        action={async () => {
          "use server";
          await signOut();
          redirect("/");
        }}
      >
        <button type="submit">Logout</button>
      </form>
      <pre>{JSON.stringify(session, null, 2)}</pre>
    </section>
  );
}