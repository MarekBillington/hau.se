import { component$, useStore, useStylesScoped$ } from "@builder.io/qwik";
import styles from './button.css?inline';

export default component$((props: any) => {
    useStylesScoped$(styles);

    const store = useStore(
        {...props},
        {recursive: true}
    );

    return (
        <>
            <input 
                class="styled-button" 
                type="button" 
                value={store.value}
                onClick$={store.click}
            />
        </>
    )
})
