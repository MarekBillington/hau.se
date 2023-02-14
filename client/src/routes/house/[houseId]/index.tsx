import { $, component$, Resource, useResource$, useStore } from "@builder.io/qwik";
import { DocumentHead, Link, useLocation } from '@builder.io/qwik-city';
import { request } from "~/components/api/api";
import Button from "~/components/inputs/button/button";
import Number from "~/components/inputs/number/number";
import Text from "~/components/inputs/text/text";
import House from '../house'

export default component$(() => {
    const location = useLocation();

    const store = useStore(
        {house: {
            id: 0,
            address: '',
            bedrooms: '',
            bathrooms: '',
            garage: '',
            floorspace: '',
            landarea: ''
        }},
        {recursive: true}
    )

    const resHouse = useResource$(async () => {
        return request('house/' + location.params.houseId, 'GET', {})
    })

    const submit  = $(async () => {
        console.log(JSON.stringify(store.house));
        console.log(store.house);

        const method = store.house.id == 0 ? 'POST' : 'PATCH';
        return request('house/' + location.params.houseId, method, JSON.stringify(store.house));
    })

    return (
        <>
            <Link href="/house">Back</Link>
            <br />
            <Resource 
                value={resHouse}
                onResolved={(house: House) => {
                    store.house = house;
                    console.log(house)
                    return (
                        <div>
                            <h2>{house.address}</h2>
                            <div class="house-form">
                                {/* @todo https://qwik.builder.io/docs/components/state/#using-context */}

                                Address: <Text value={store.house.address} />
                                <br/>
                                Bedrooms: <Number value={store.house.bedrooms} />
                                <br/>
                                Bathroom: <Number value={store.house.bathrooms} />
                                <br/>
                                Garage: <Number value={store.house.garage} />
                                <br/>
                                Floorspace (m&sup2;): <Number value={store.house.floorspace} />
                                <br/>
                                Landarea (m&sup2;): <Number value={store.house.landarea} />
                                <br/>
                                <Button value="Submit" click={submit}/>
                            </div>
                        </div>
                    )
                }}
            />
        </>
    )
})


export const head: DocumentHead = {
    title: 'Hause',
};
