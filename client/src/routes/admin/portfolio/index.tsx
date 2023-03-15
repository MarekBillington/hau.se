import { component$ } from "@builder.io/qwik";
import { DocumentHead } from "@builder.io/qwik-city";
import PortfolioScreen from "~/components/admin/screens/portfolio/main/portfolio-screen";

export default component$(() => {
  return (
    <>
      <PortfolioScreen />
    </>
  )
})

export const head: DocumentHead = {
  title: "Hause",
};