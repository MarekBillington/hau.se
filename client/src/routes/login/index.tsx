import {
  $,
  component$,
  QwikChangeEvent,
  useContext,
  useStore,
} from "@builder.io/qwik";
import { DocumentHead, useNavigate } from "@builder.io/qwik-city";
import { getUserInfo, login } from "~/components/auth/auth";
import { setProperty } from "~/components/common/types";
import Button from "~/components/inputs/button/button";
import Password from "~/components/inputs/password/password";
import Text from "~/components/inputs/text/text";
import { authCtx, userSession } from "~/root";

export default component$(() => {
  const auth = useContext(authCtx);
  const sess = useContext(userSession);
  const store = useStore({
    email: "",
    password: "",
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
    const res = await login(store.email, store.password);
    auth.token = res.token;
    auth.expiry = res.expiry;
    getUserInfo(auth, sess)
    nav('/')
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
      <Button value="Login" click={click} />
    </>
  );
});

export const head: DocumentHead = {
  title: "Hause",
};
