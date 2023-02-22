import { $, component$, QwikChangeEvent, useStore } from "@builder.io/qwik";
import { setProperty } from "~/components/common/types";
import Button from "~/components/inputs/button/button";
import Password from "~/components/inputs/password/password";
import Text from "~/components/inputs/text/text";


export default component$(() => {
    const store = useStore({
        email: "",
        password: "",
    })

    const login = $(async () => {


    })

    const onChange = $((event: QwikChangeEvent<HTMLInputElement>) => {
        type keyType = keyof (typeof store);
        // @ts-ignore complaining about it being asignable, but it works lmao
        const k: keyType = event.target.name;
        
        const val = event.target.value;
        
        setProperty(store, k, val);
    })

    return (
        <>
            <Text 
                label="Email"
                name="email"
                value={store.email}
                change={onChange}
            />
            <Password
                label="Password"
                name="password"
                value={store.password}
                change={onChange}
            />
            <Button 
                value="Login"
                click={login}
            />
        </>
    )
})