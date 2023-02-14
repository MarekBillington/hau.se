import { component$, useStore, useStylesScoped$ } from "@builder.io/qwik";
import styles from './text.css?inline';


export default component$((props: any) => {
    useStylesScoped$(styles);

    const store = useStore(
        {...props},
        {recursive: true}
    );

    return (
        <>
            <input
                class="styled-text"
                type="text"
                value={store.value} 
                onClick$={store.click}
            />
        </>
    )
})