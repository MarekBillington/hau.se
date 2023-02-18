import { component$, useStylesScoped$ } from "@builder.io/qwik";
import { styled } from 'styled-vanilla-extract/qwik';
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

// @todo look at stlyed components

export const Button = styled.input`
    color: var(--hause-dark-blue);
    width: 64px;
    height: 25px;
    border: 1.5px solid;
    border-radius: 3px;
    background: none;
    margin-top: 10px;
    margin-bottom: 10px;
`;
