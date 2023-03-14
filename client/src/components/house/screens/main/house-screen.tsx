import { $, Resource, component$, useContext, useResource$, useStore, useTask$ } from "@builder.io/qwik";
import { Link, useLocation } from "@builder.io/qwik-city";
import { authCtx, userSession } from "~/root";
import { request } from "~/components/utility/api/api";
import { QCE, onChange, onChangeNum } from "~/components/utility/helper/event";
import Button from "~/components/utility/inputs/button/button";
import Number from "~/components/utility/inputs/number/number";
import Text from "~/components/utility/inputs/text/text";

import type House from "../../interfaces/house";
import Dropdown from "~/components/utility/inputs/dropdown/dropdown";
import { Country } from "~/components/utility/country/interface/country";

export default component$(() => {
  const location = useLocation();
  const auth = useContext(authCtx);
  const sess = useContext(userSession);
  const store = useStore(
    {
      isActive: false,
      isNewHouse: false,
      updated: 0,
      house: {} as House,
      countries: {
        default: sess.portfolio.country_id,
        data: []
      }
    },
    {
      recursive: true,
    }
  );

  useTask$(async () => {
    const c = await request("utility/country", auth, "GET")
    const fc = c.map((obj: Country) => {
      return {
        value: obj.id,
        text: obj.name
      }
    })
    store.countries.data = fc;
  })

  const resHouse = useResource$<House>(async ({ track }) => {
    track(() => store.isActive);
    const u = track(() => store.updated);
    let loc = location.params.houseId
    if (u != 0) {
      loc = store.house.id.toString()
    }
    const h = await request("house/" + loc, auth, "GET");
    store.house = h;
    return h;
  });

  const submit = $(async () => {
    const method = store.isNewHouse ? "POST" : "PATCH";
    const url = "house" + (store.isNewHouse ? "" : ("/" + store.house.id));

    store.house = await request(url, auth, method, JSON.stringify(store.house));
    // triggers the tracking to refresh the screen
    store.updated++;
  });

  const toggleActive = $(async () => {
    const url = "house/" + store.house.id + "/active";
    await request(url, auth, "PATCH");
    store.isActive = !store.isActive;
  });

  function isActive() {
    return store.house.active ? "Disable" : "Enable";
  }

  const handler = $((e: QCE) => onChange(e, store.house))
  const selectHandler = $((e: QCE) => onChangeNum(e, store.house))

  return (
    <>
      <div class="toolbar">
        <div class="toolbar-left">
          <Link href="/house">Back</Link>
        </div>
        <div class="toolbar-right">
          {/* Turn into styled component with vanilla */}
          <Button value={isActive()} click={toggleActive} />
        </div>
      </div>
      <div>
        <Resource
          value={resHouse}
          onResolved={(house: House) => {
            if (Object.keys(house).length == 0) {
              store.isNewHouse = true;
              store.house.portfolio_id = sess.portfolio.id
              // for the sake of working, to create a 
              store.house.country_id = store.countries.default
            }

            
            return (
              <>
                <div>
                  <h2>{house.street_1}</h2>
                  <div class="house-form">
                    <div>
                      <Text label="Street 1:" value={house.street_1} name="street_1" change={handler} />
                      <br />
                      <Text label="Street 2:" value={house.street_2} name="street_2" change={handler} />
                      <br />
                      <Text label="Postcode:" value={house.postcode} name="postcode" change={handler} />
                      <br />
                      <Text label="Town:" value={house.town} name="town" change={handler} />
                      <br />
                      {/* Get options from server task */}
                      <Dropdown label="Country: " name="country_id" default={store.house.country_id} options={store.countries.data} change={selectHandler} />
                    </div>
                    <div>
                      <Number label="Bedrooms:" value={house.bedrooms} name="bedrooms" change={handler} />
                      <br />
                      <Number label="Bathroom:" value={house.bathrooms} name="bathrooms" change={handler} />
                      <br />
                      <Number label="Garage:" value={house.garage} name="garage" change={handler} />
                      <br />
                      <Number label="Floorspace (m&sup2;):" value={house.floorspace} name="floorspace" change={handler} />
                      <br />
                      <Number label="Landarea (m&sup2;):" value={house.landarea} name="landarea" change={handler} />
                    </div>
                  </div>
                </div>
                <Button value="Save" click={submit} />
              </>
            );
          }}
        />
      </div>
    </>
  );
})
