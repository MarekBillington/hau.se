import { $, QwikChangeEvent, component$ } from "@builder.io/qwik";
import { setProperty } from "~/components/utility/helper/types";
import Text from "~/components/utility/inputs/text/text";

import type Address from "~/components/house/interfaces/address";

export default component$((address: Address) => {

  const onChange = $((event: QwikChangeEvent<HTMLInputElement>) => {
    type keyType = keyof typeof address;
    // @ts-ignore complaining about it being asignable, but it works lmao
    const k: keyType = event.target.name;

    const val = isNaN(event.target.valueAsNumber)
      ? event.target.value
      : event.target.valueAsNumber;

    setProperty(address, k, val);
  });

  return (
    <>
      <Text
          label="Street 1:"
          value={address.street_1}
          name="street_1"
          change={onChange}
        />
        <br />
        <Text
          label="Street 2:"
          value={address.street_2}
          name="street_2"
          change={onChange}
        />
        <br />
        <Text
          label="Postcode:"
          value={address.postcode}
          name="postcode"
          change={onChange}
        />
        <br />
        <Text
          label="Town/City:"
          value={address.town}
          name="town"
          change={onChange}
        />
    </>
  )
})