import { component$ } from "@builder.io/qwik";
import { DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  return (
    <>
      <p>A person</p>
    </>
  );
});

export const head: DocumentHead = {
  title: "Hause",
};
