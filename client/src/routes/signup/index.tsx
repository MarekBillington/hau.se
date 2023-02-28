import { $, component$, QwikChangeEvent, useContext, useStore } from "@builder.io/qwik";
import { DocumentHead, useNavigate } from "@builder.io/qwik-city";
import { signup } from "~/components/auth/auth";
import { setProperty } from "~/components/common/types";
import Button from "~/components/inputs/button/button";
import Password from "~/components/inputs/password/password";
import Text from "~/components/inputs/text/text";
import { authCtx } from "~/root";

export default component$(() => {
  const auth = useContext(authCtx)
  const store = useStore({
    email: "",
    firstName: "",
    lastName: "",
    password: "",
    repassword: "",
  });
  const nav = useNavigate()

  const onChange = $((event: QwikChangeEvent<HTMLInputElement>) => {
    type keyType = keyof typeof store;
    // @ts-ignore complaining about it being asignable, but it works lmao
    const k: keyType = event.target.name;

    const val = event.target.value;

    setProperty(store, k, val);
  });

  const click = $(async () => {
    const res = await signup(store)
    auth.token = res.token;
    auth.expiry = res.expiry;
    nav('/signup/new/')
  });

  return (
    <>
      <Text label="Email" name="email" value={store.email} change={onChange} />
      <br />
      <Password
        label="Password"
        name="password"
        value={store.password}
        change={onChange}
      />
      <br />
      <Password
        label="Re-type Password"
        name="repassword"
        value={store.password}
        change={onChange}
      />
      <br />
      <Text
        label="First Name"
        name="firstName"
        value={store.firstName}
        change={onChange}
      />
      <br />
      <Text
        label="Last Name"
        name="lastName"
        value={store.lastName}
        change={onChange}
      />
      <br />
      <Button value="Signup" click={click} />
    </>
  );
});

export const head: DocumentHead = {
  title: "Hause",
};
