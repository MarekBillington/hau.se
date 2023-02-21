import { component$ } from "@builder.io/qwik";


export default component$((props: any) => {

    return (
        <>
            <label class="styled-label">
                {props.label}
            </label>
            <input 
                class="styled-checkbox"
                type="checkbox"
                name={props.name}
                value={props.value}
                checked
                onChange$={props.change}

            />
        </>
    )
})