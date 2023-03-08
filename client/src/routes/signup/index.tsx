import { $, component$, QwikChangeEvent, useContext, useStore } from "@builder.io/qwik";
import { DocumentHead, useNavigate } from "@builder.io/qwik-city";
import { getUserInfo, signup } from "~/components/utility/auth/auth";
import { setProperty } from "~/components/utility/helper/types";
import Button from "~/components/utility/inputs/button/button";
import Password from "~/components/utility/inputs/password/password";
import Text from "~/components/utility/inputs/text/text";
import { authCtx, userSession } from "~/root";

export default component$(() => {
  const auth = useContext(authCtx)
  const sess = useContext(userSession)
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
    getUserInfo(auth, sess)
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
