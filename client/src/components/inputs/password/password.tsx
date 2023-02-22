import { component$ } from "@builder.io/qwik";


export default component$((props: any) => {
    return (
        <>
            <label>{props.label}</label>
            <input
                type="password"
                value={props.value}
                name={props.name}
                onChange$={props.change}
            />
        </>
    )
})