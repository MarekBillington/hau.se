import { component$, Slot } from "@builder.io/qwik";
import Header from "../components/utility/header/header";

export default component$(() => {
  return (
    <>
      <main>
        <Header />
        <section>
          <Slot />
        </section>
      </main>
      <footer></footer>
    </>
  );
});
