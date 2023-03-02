import {
  $,
  component$,
  Resource,
  useContext,
  useResource$,
  useStore,
  useStylesScoped$,
} from "@builder.io/qwik";
import { DocumentHead, Link } from "@builder.io/qwik-city";
import { request } from "~/components/api/api";
import Checkbox from "~/components/inputs/checkbox/checkbox";
import { authCtx } from "~/root";
import styles from "./house.css?inline";

import type House from "../../components/interfaces/house";

export default component$(() => {
  const auth = useContext(authCtx);
  const store = useStore({
    isActive: true,
    houses: [],
  });

  const houseList = useResource$<Array<House>>(async ({ track }) => {
    track(() => store.isActive);
    const url = "house?active=" + store.isActive;

    const houses = await request(url, auth, "GET");
    store.houses = houses;
    return houses;
  });

  const checkActive = $(() => {
    store.isActive = !store.isActive;
  });

  return (
    <div>
      <div class="toolbar">
        <div class="toolbar-left">
          <Link href="/house/0/">Add New House</Link>
        </div>
        <div>
          <Checkbox
            label="Is Active"
            name="isActive"
            value={store.isActive}
            change={checkActive}
          />
        </div>
      </div>
      <div>
        <Resource
          value={houseList}
          onPending={() => <div>Loading...</div>}
          onResolved={(houses: Array<House>) => {
            let hs = {};
            if (Array.isArray(houses)) {
              hs = houses.map((h) => {
                return <Panel {...h} />;
              });
            }
            return <>{hs}</>;
          }}
        />
      </div>
    </div>
  );
});

export const Panel = component$((house: House) => {
  useStylesScoped$(styles);

  return (
    <div class="house-tile">
      <div>{/** add an image here eventually */}</div>
      <Link href={house.id.toString()}>{house.address}</Link>
    </div>
  );
});

export const head: DocumentHead = {
  title: "Hause",
};
