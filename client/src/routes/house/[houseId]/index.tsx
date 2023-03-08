import {
  $,
  component$,
  QwikChangeEvent,
  Resource,
  useContext,
  useResource$,
  useStore,
} from "@builder.io/qwik";
import { DocumentHead, Link, useLocation } from "@builder.io/qwik-city";
import { request } from "~/components/api/api";
import { setProperty } from "~/components/common/types";
import Button from "~/components/inputs/button/button";
import Number from "~/components/inputs/number/number";
import Text from "~/components/inputs/text/text";
import { authCtx } from "~/root";

import type House from "../../../components/interfaces/house";

export default component$(() => {
  const location = useLocation();
  const auth = useContext(authCtx);
  const store = useStore(
    {
      isActive: false,
      isNewHouse: false,
      updated: 0,
      house: {} as House,
    },
    {
      recursive: true,
    }
  );

  const resHouse = useResource$<House>(async ({ track }) => {
    track(() => store.isActive);
    track(() => store.updated);
    const h = await request("house/" + location.params.houseId, auth, "GET");
    store.house = h;
    return h;
  });

  const submit = $(async () => {
    const method = store.isNewHouse ? "POST" : "PATCH";
    const url = "house" + (store.isNewHouse ? "" : "/" + store.house.id);

    await request(url, auth, method, JSON.stringify(store.house));

    // triggers the tracking to refresh the screen
    store.updated++;
  });

  const toggleActive = $(async () => {
    const url =
      "house" +
      "/" +
      store.house.id +
      (store.house.active ? "/disable" : "/enable");

    await request(url, auth, "PATCH");

    store.isActive = !store.isActive;
  });

  const onChange = $((event: QwikChangeEvent<HTMLInputElement>) => {
    type keyType = keyof typeof store.house;
    // @ts-ignore complaining about it being asignable, but it works lmao
    const k: keyType = event.target.name;

    const val = isNaN(event.target.valueAsNumber)
      ? event.target.value
      : event.target.valueAsNumber;

    setProperty(store.house, k, val);
  });

  function isActive() {
    return store.house.active ? "Disable" : "Enable";
  }

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
            if (house.id === 0) {
              store.isNewHouse = true;
            }
            return (
              <div>
                <h2>{house.address[0].street_1}</h2>
                <div class="house-form">
                  <Text
                    label="Address:"
                    value={store.house.address[0].street_1}
                    name="address"
                    change={onChange}
                  />
                  <br />
                  <Number
                    label="Bedrooms:"
                    value={store.house.bedrooms}
                    name="bedrooms"
                    change={onChange}
                  />
                  <br />
                  <Number
                    label="Bathroom:"
                    value={store.house.bathrooms}
                    name="bathrooms"
                    change={onChange}
                  />
                  <br />
                  <Number
                    label="Garage:"
                    value={store.house.garage}
                    name="garage"
                    change={onChange}
                  />
                  <br />
                  <Number
                    label="Floorspace (m&sup2;):"
                    value={store.house.floorspace}
                    name="floorspace"
                    change={onChange}
                  />
                  <br />
                  <Number
                    label="Landarea (m&sup2;):"
                    value={store.house.landarea}
                    name="landarea"
                    change={onChange}
                  />
                  <br />
                  <Button value="Save" click={submit} />
                </div>
              </div>
            );
          }}
        />
      </div>
    </>
  );
});

export const head: DocumentHead = {
  title: "Hause",
};
