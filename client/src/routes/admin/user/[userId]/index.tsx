import { component$ } from "@builder.io/qwik";
import { DocumentHead, useLocation } from "@builder.io/qwik-city";
import UserScreen from "~/components/admin/screens/user/main/user-screen";


export default component$(() => {
  const loc = useLocation()

  console.log(loc.params.userId)
  return (
    <>
      <UserScreen />
    </>
  )
})

export const head: DocumentHead = {
  title: "Hause",
};
