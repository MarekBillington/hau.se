import {
  $,
  component$,
  useContext,
  useStore,
} from "@builder.io/qwik";
import { DocumentHead, useNavigate } from "@builder.io/qwik-city";
import { getUserInfo, login } from "~/components/utility/auth/auth";
import { onChange, QCE } from "~/components/utility/helper/event";
import Button from "~/components/utility/inputs/button/button";
import Password from "~/components/utility/inputs/password/password";
import Text from "~/components/utility/inputs/text/text";
import { authCtx, userSession } from "~/root";

export default component$(() => {
  const auth = useContext(authCtx);
  const sess = useContext(userSession);
  const store = useStore({
    email: "",
    password: "",
  });

  const nav = useNavigate()

  const helper = $((e: QCE) => onChange(e, store))

  const click = $(async () => {
    const res = await login(store.email, store.password);
    auth.token = res.token;
    auth.expiry = res.expiry;
    getUserInfo(auth, sess)
    nav("/")
  });

  return (
    <>
      <Text label="Email" name="email" value={store.email} change={helper} />
      <br />
      <Password label="Password" name="password" value={store.password} change={helper} />
      <br />
      <Button value="Login" click={click} />
    </>
  );
});

export const head: DocumentHead = {
  title: "Hause",
};
