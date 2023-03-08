import { component$ } from "@builder.io/qwik";
import House from "~/components/house/interfaces/house";
import Button from "~/components/utility/inputs/button/button";
import Number from "~/components/utility/inputs/number/number";
import Text from "~/components/utility/inputs/text/text";

export default component$((house: House, submit: $) => {
  return (
    <div>
      <h2>{house.address[0].street_1}</h2>
      <div class="house-form">
        <Text
          label="Address:"
          value={house.address[0].street_1}
          name="address"
          change={onChange}
        />
        <br />
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
        <br />
        <Button value="Save" click={submit} />
      </div>
    </div>
  )
})