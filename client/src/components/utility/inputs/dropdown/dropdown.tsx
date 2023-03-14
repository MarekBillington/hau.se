import { component$, useStylesScoped$ } from "@builder.io/qwik";
import styles from "./dropdown.css?inline";

export default component$((props: any) => {
  useStylesScoped$(styles);

  return (
    <>
      <label class="styled-label">{props.label}</label>
      <select 
        onChange$={props.change}
        name={props.name}

      >
        {props.options.map((obj: {value: number, text: string}) => {
          const s = props.default == obj.value
          return (
            <option value={obj.value} selected={s}>{obj.text}</option>
          )
        })}
      </select>
    </>
  );
});
