import { component$, useStore, useStylesScoped$ } from "@builder.io/qwik";
import styles from "./number.css?inline";

export default component$((props: any) => {
    useStylesScoped$(styles);

    const store = useStore(
        {...props},
        {recursive: true}
    )

    return (
        <>
            <input 
                class="styled-number"
                type="number"
                value={store.value}
                onClick$={store.click}
            />
        </>
    )
})