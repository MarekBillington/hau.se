import { component$, useContext } from "@builder.io/qwik";
import { DocumentHead, Link } from "@builder.io/qwik-city";
import UserDashboard from "~/components/admin/screens/dashboard/user/user-dashboard";
import { userSession } from "~/root";

export default component$(() => {
  const sesn = useContext(userSession)
  
  return (
    <>
      <div>
        <h2>{sesn.portfolio.name}</h2>
        <Link href="portfolio">edit portfolio</Link>
      </div>
      <UserDashboard />
    </>
  );
});

export const head: DocumentHead = {
  title: "Hause",
};
