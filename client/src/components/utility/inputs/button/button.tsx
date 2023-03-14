import { component$, useStylesScoped$ } from "@builder.io/qwik";
import styles from "./button.css?inline";

export default component$((props: any) => {
  useStylesScoped$(styles);

  return (
    <>
      <button
        class="styled-button"
        type="button"
        onClick$={props.click}
      > 
        {props.value}
      </button>
    </>
  );
});
