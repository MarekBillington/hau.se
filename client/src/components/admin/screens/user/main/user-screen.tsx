import { $, Resource, component$, useContext, useResource$, useStore } from "@builder.io/qwik";
import { Link, useLocation } from "@builder.io/qwik-city";

import type User from "~/components/interfaces/user";
import { request } from "~/components/utility/api/api";
import { QCE, onChange } from "~/components/utility/helper/event";
import Button from "~/components/utility/inputs/button/button";
import Password from "~/components/utility/inputs/password/password";
import Text from "~/components/utility/inputs/text/text";
import { authCtx } from "~/root";


export default component$(() => {
  const loc = useLocation();
  const auth = useContext(authCtx);
  const store = useStore({
    isNewUser: false,
    isActive: true,
    updated: 0,
    user: {} as User
  }, {deep: true});

  const res = useResource$<User>(async ({track}) => {
    track(() => store.updated)
    track(() => store.isActive)
    const u = await request("user/" + loc.params.userId, auth, "GET")
    store.user = u
    return u
  })

  const submit = $(async () => {
    const method = store.isNewUser ? "POST" : "PATCH";
    const url = "user" + (store.isNewUser ? "" : ("/" + store.user.id));

    store.user = await request(url, auth, method, JSON.stringify(store.user));
    // triggers the tracking to refresh the screen
    store.updated++;
  });

  const toggleActive = $(async () => {
    return
    // const url = "user/" + store.user.id + "/active";
    // await request(url, auth, "PATCH");
    // store.isActive = !store.isActive;
  });

  function isActive() {
    return store.user.active ? "Disable" : "Enable";
  }

  const handler = $((e: QCE) => onChange(e, store.user))

  return (
    <>
    <div class="toolbar">
        <div class="toolbar-left">
          <Link href="/admin">Back</Link>
        </div>
        <div class="toolbar-right">
          <Button value={isActive()} click={toggleActive} />
        </div>
      </div>
      <div>
        <Resource 
          value={res}
          onResolved={(user: User) => {

            return (
              <>
                <Text label="First Name:" name="first_name" value={user.first_name} change={handler} />
                <br />
                <Text label="Last Name:" name="last_name" value={user.first_name} change={handler} />
                <br />
                <Password label="Password:" name="password" value={user.password} change={handler} />
                <br />
                <Button value="Save" click={submit} />
              </>
            )
          }}
        />
      </div>
    </>
  )
})