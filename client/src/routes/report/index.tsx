import { component$ } from "@builder.io/qwik";
import { DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  return (
    <>
      <p>Some data</p>
    </>
  );
});

export const head: DocumentHead = {
  title: "Hause",
};
