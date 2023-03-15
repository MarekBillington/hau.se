import { $, component$, Resource, useContext, useResource$, useStore, useTask$ } from "@builder.io/qwik";
import { request } from "~/components/utility/api/api";
import { authCtx, userSession } from "~/root";

import type User from "~/components/interfaces/user";
import type Portfolio from "~/components/interfaces/portfolio";
import Text from "~/components/utility/inputs/text/text";
import Button from "~/components/utility/inputs/button/button";
import { onChange, onChangeNum, QCE } from "~/components/utility/helper/event";
import { useNavigate } from "@builder.io/qwik-city";
import Dropdown from "~/components/utility/inputs/dropdown/dropdown";
import { Country } from "~/components/utility/country/interface/country";

export default component$(() => {
  const auth = useContext(authCtx);
  const sesn = useContext(userSession);
  const store = useStore({
      user: {} as User,
      portfolio: {} as Portfolio,
      countries: []
    }, { deep: true }
  );
  const nav = useNavigate()

  useTask$(async () => {
    const c = await request("utility/country", auth, "GET")
    const fc = c.map((obj: Country) => {
      return {
        value: obj.id,
        text: obj.name
      }
    })
    store.countries = fc;
  })

  const userRes = useResource$(async () => {
    const setup = await request("user/setup", auth, "GET");
    store.user = setup.user;
    store.portfolio = setup.portfolio;
    return setup;
  })

  const submit = $(async () => {
    const body = {
      name: store.portfolio.name,
      user_id: store.user.id,
      country_id: store.portfolio.country_id
    }
    const p = await request("portfolio/create", auth, "POST", JSON.stringify(body));
    if (!p.error) {
      sesn.portfolio = p;
      nav("/");
    }
  })

  const handler = $((e: QCE) => onChange(e, store.portfolio))
  const selectHandler = $((e: QCE) => onChangeNum(e, store.portfolio))

  return (
    <>
    <Resource
      value={userRes}
      onPending={() => <div>Loading...</div>}
      onResolved={(setup) => {
        return (
          <>
            <h2>Hello {setup.user.first_name}</h2>
            <p>We just need a bit more info to get started!</p>
            <Text label="Portfolio Name: " value={store.portfolio.name} name="name" change={handler}/>
            <br />
            <Dropdown label="Country: " name="country_id" options={store.countries} change={selectHandler} />
            <br />
            <Button value="Submit" click={submit}/>
          </>
        )
      }}
    />
    </>
  )
})

