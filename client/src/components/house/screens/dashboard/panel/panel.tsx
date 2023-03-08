import { component$, useStylesScoped$ } from "@builder.io/qwik";
import House from "../../../interfaces/house";
import styles from "./panel.css?inline";
import { Link } from "@builder.io/qwik-city";

export const Panel = component$((house: House) => {
  useStylesScoped$(styles);

  const address = house.address[0]
    
  return (
    <div class="house-tile">
      <div>{/** add an image here eventually */}</div>
      <Link href={house.id.toString()}>{address.street_1}</Link>
    </div>
  );
});