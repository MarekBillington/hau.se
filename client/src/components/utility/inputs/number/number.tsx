import { component$, useStylesScoped$ } from "@builder.io/qwik";
import styles from "./number.css?inline";

type NumberInput = {
  label: string,
  name: string,
  value: number,
  change?: any
}

export default component$((props: NumberInput) => {
  useStylesScoped$(styles);

  return (
    <>
      <label class="styled-label">{props.label}</label>
      <input
        class="styled-number"
        type="number"
        name={props.name}
        value={props.value}
        onChange$={props.change}
      />
    </>
  );
});
