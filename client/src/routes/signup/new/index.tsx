import { component$, Resource, useContext, useResource$, useStore } from "@builder.io/qwik";
import { request } from "~/components/api/api";
import User from "~/components/interfaces/user";
import { authCtx } from "~/root";


export default component$(() => {
  const auth = useContext(authCtx)
  //const sesn = useContext(userSession)
  const store = useStore({
      user: {} as User,
      portfolio: {
        name: ""
      }
    }, { deep: true })

  const userRes = useResource$(async () => {
    const user = await request("user/setup", auth, "GET")
    store.user = user
    console.log(user)
    return user
  })

  // const onChange = $((event: QwikChangeEvent<HTMLInputElement>) => {
  //   type keyType = keyof typeof store;
  //   // @ts-ignore complaining about it being asignable, but it works lmao
  //   const k: keyType = event.target.name;

  //   const val = isNaN(event.target.valueAsNumber)
  //     ? event.target.value
  //     : event.target.valueAsNumber;

  //   setProperty(store, k, val);
  // });

  return (
    <>
    <Resource
      value={userRes}
      onPending={() => <div>Loading...</div>}
      onResolved={(user: User) => {
        console.log(user)
        return (
          <>
            <h2>Hi {user.firstName}</h2>
          </>
        )
      }}
    />
    </>
  )
})

