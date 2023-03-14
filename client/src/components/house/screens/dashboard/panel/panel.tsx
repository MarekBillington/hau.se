import { component$, useStylesScoped$ } from "@builder.io/qwik";
import { Link } from "@builder.io/qwik-city";
import styles from "./panel.css?inline";

import type House from "../../../interfaces/house";

export const Panel = component$((house: House) => {
  useStylesScoped$(styles);
    
  return (
    <div class="house-tile">
      <div>{/** add an image here eventually */}</div>
      <Link href={house.id.toString()}>{house.street_1}</Link>
    </div>
  );
});