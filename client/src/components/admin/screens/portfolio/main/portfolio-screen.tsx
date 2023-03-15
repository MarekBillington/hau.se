import { $, component$, useContext, useStore, useTask$ } from "@builder.io/qwik";
import { Link } from "@builder.io/qwik-city";
import { request } from "~/components/utility/api/api";
import { getCountries } from "~/components/utility/country/task/task";
import { QCE, onChange, onChangeNum } from "~/components/utility/helper/event";
import Button from "~/components/utility/inputs/button/button";
import Dropdown from "~/components/utility/inputs/dropdown/dropdown";
import Text from "~/components/utility/inputs/text/text";
import { authCtx, userSession } from "~/root";

export default component$(() => {
  const auth = useContext(authCtx)
  const sesn = useContext(userSession)
  const store = useStore({
    portfolio: sesn.portfolio,
    countries: {
      default: sesn.portfolio.country_id,
      data: []
    }
  }, {deep: true})

  useTask$(async () => {
    store.countries.data = await getCountries(auth);
  })

  const submit = $(async () => {
    const body = {
      name: store.portfolio.name,
      country_id: store.portfolio.country_id
    }
    const p = await request("portfolio/" + store.portfolio.id, auth, "PATCH", JSON.stringify(body));
    if (!p.error) {
      sesn.portfolio = p;
    }
  })

  const handler = $((e: QCE) => onChange(e, store.portfolio))
  const selectHandler = $((e: QCE) => onChangeNum(e, store.portfolio))

  return (
    <>
      <div class="toolbar">
        <div class="toolbar-left">
          <Link href="/admin">Back</Link>
        </div>
        <div class="toolbar-right">
        </div>
      </div>
      <div>
        <Text label="Name:" value={store.portfolio.name} name="name" change={handler} />
        <br />
        <Dropdown label="Country:" name="country_id" options={store.countries.data} change={selectHandler} default={store.countries.default} />
        <br />
        <Button value="Submit" click={submit} />
      </div>
    </>
  )
})