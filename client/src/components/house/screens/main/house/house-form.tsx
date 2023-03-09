import { $, QwikChangeEvent, component$ } from "@builder.io/qwik";
import { setProperty } from "~/components/utility/helper/types";
import Number from "~/components/utility/inputs/number/number";
import AddressForm from "./address/address-form";

import type House from "~/components/house/interfaces/house";

export default component$((house: House) => {
  console.log(house)

  const onChange = $((event: QwikChangeEvent<HTMLInputElement>) => {
    type keyType = keyof typeof house;
    // @ts-ignore complaining about it being asignable, but it works lmao
    const k: keyType = event.target.name;

    const val = isNaN(event.target.valueAsNumber)
      ? event.target.value
      : event.target.valueAsNumber;

    setProperty(house, k, val);
  });

  return (
    <div>
      <h2>{house.address.street_1}</h2>
      <div class="house-form">
        <div>
          <AddressForm 
            {...house.address}
          />
        </div>
        <div>
          <Number
            label="Bedrooms:"
            value={house.bedrooms}
            name="bedrooms"
            change={onChange}
          />
          <br />
          <Number
            label="Bathroom:"
            value={house.bathrooms}
            name="bathrooms"
            change={onChange}
          />
          <br />
          <Number
            label="Garage:"
            value={house.garage}
            name="garage"
            change={onChange}
          />
          <br />
          <Number
            label="Floorspace (m&sup2;):"
            value={house.floorspace}
            name="floorspace"
            change={onChange}
          />
          <br />
          <Number
            label="Landarea (m&sup2;):"
            value={house.landarea}
            name="landarea"
            change={onChange}
          />
        </div>
      </div>
    </div>
  )
})