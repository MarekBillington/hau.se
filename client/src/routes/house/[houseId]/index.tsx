import {  component$ } from "@builder.io/qwik";
import { DocumentHead } from "@builder.io/qwik-city";
import HouseScreen from "~/components/house/screens/main/house-screen";

export default component$(() => {
  return <HouseScreen />
});

export const head: DocumentHead = {
  title: "Hause",
};
