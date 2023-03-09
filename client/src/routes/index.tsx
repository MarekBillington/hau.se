import { component$, useContext } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";
import { authCtx, userSession } from "~/root";

export const App = component$(() => {
  const auth = useContext(authCtx)
  const sess = useContext(userSession)
  
  console.log("main page");

  // could use Slots...?
  let content = (<div>Not logged in</div>)
  if (auth.token != '') {
    content = (
      <div>
        Dashboard data <br />
        Hello {sess.user.first_name} {sess.user.last_name}
      </div>
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
