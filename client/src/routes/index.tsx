import { component$, useContext } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";
import { authCtx, userSession } from "~/root";

export const App = component$(() => {
  const auth = useContext(authCtx)
  const sess = useContext(userSession)
  
  console.log("main page");

  // could use Slots...?
  let content = (<>Not logged in</>)
  if (auth.token != '') {
    content = (
      <>
        Dashboard data <br />
        Hello {sess.user.firstName}
      </>
    )
  }

  return (
    <>{content}</>
  );
});

export default App;

export const head: DocumentHead = {
  title: "Hause",
  meta: [
    {
      name: "description",
      content: "Hau.se title page",
    },
  ],
};
