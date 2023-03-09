import {
  $,
  component$,
  Resource,
  useContext,
  useResource$,
  useStore,
} from "@builder.io/qwik";
import { DocumentHead, Link, useLocation } from "@builder.io/qwik-city";
import { request } from "~/components/utility/api/api";
import { authCtx } from "~/root";
import Button from "~/components/utility/inputs/button/button";
import HouseForm from "~/components/house/screens/main/house/house-form";

import type House from "../../../components/house/interfaces/house";

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
    const url = "house/" + store.house.id + "/active";
    await request(url, auth, "PATCH");
    store.isActive = !store.isActive;
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
            console.log(house)
            if (house.id === 0) {
              store.isNewHouse = true;
            }
            return (
              <>
                <HouseForm {...house} />
                <Button value="Save" click={submit} />
              </>
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
