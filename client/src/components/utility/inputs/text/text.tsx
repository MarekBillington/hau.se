import { component$, useStylesScoped$ } from "@builder.io/qwik";
import styles from "./text.css?inline";

export default component$((props: any) => {
  useStylesScoped$(styles);

  return (
    <>
      <label class="styled-label">{props.label}</label>
      <input
        class="styled-text"
        type="text"
        name={props.name}
        value={props.value}
        onChange$={props.change}
      />
    </>
  );
});
