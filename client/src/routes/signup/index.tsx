import { $, component$, useContext, useStore } from "@builder.io/qwik";
import { DocumentHead, useNavigate } from "@builder.io/qwik-city";
import { getUserInfo, signup } from "~/components/utility/auth/auth";
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
    first_name: "",
    last_name: "",
    password: "",
    repassword: "",
  });
  const nav = useNavigate();

  const helper = $((e: QCE) => onChange(e, store));

  const click = $(async () => {

    // perform some validation
    if (store.password != store.repassword) {
      alert("The passwords do not match")
      return
    }

    const res = await signup(store);
    if (!res.error) {
      auth.token = res.token;
      auth.expiry = res.expiry;
      getUserInfo(auth, sess);
      nav("/signup/new/");
    }
  });

  return (
    <>
      <Text label="Email" name="email" value={store.email} change={helper} />
      <br />
      <Password label="Password" name="password" value={store.password} change={helper} />
      <br />
      <Password label="Re-type Password" name="repassword" value={store.repassword} change={helper} />
      <br />
      <Text label="First Name" name="first_name" value={store.first_name} change={helper} />
      <br />
      <Text label="Last Name" name="last_name" value={store.last_name} change={helper} />
      <br />
      <Button value="Signup" click={click} />
    </>
  );
});

export const head: DocumentHead = {
  title: "Hause",
};
