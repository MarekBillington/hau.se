import { component$, useStylesScoped$ } from "@builder.io/qwik";
import styles from './popup.css?inline';

export default component$((props: {content?: any}) => {
    useStylesScoped$(styles)
    console.log(props.content)

    return (
        <div>
            fuck
            {props.content}
        </div>
    )
})