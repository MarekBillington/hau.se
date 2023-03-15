import { Resource, component$, useContext, useResource$, useStore } from "@builder.io/qwik";
import { Link } from "@builder.io/qwik-city";
import User from "~/components/interfaces/user";
import { request } from "~/components/utility/api/api";
import { authCtx } from "~/root";

export default component$(() => {
  const auth = useContext(authCtx)
  const store = useStore({
    users: []
  }, {deep: true})

  const res = useResource$<Array<User>>(async () => {
    // @todo add ordering list by last name
    const users = await request("user", auth, "GET")
    store.users = users
    return users
  })
  
  return (
    <>
    <h4>User List:</h4>
    <Resource 
      value={res}
      onResolved={(users: Array<User>) => {
        console.log(users)
        let us = {};
        if (Array.isArray(users)) {
          us = users.map((u: User) => {
            return (
              <>
                <Link href={"user/" + u.id}>{u.last_name + ", " + u.first_name}</Link>
                <br />
              </>
            );
          });
        }
        return (
          <>{us}</>
        )
      }}
    />
    </>
  )
})
