import { component$, useStylesScoped$ } from "@builder.io/qwik";
import styles from './button.css?inline';

export default component$((props: any) => {
    useStylesScoped$(styles);

    return (
        <>
            <input 
                class="styled-button" 
                type="button" 
                value={props.value}
                onClick$={props.click}
            />
        </>
    )
})
