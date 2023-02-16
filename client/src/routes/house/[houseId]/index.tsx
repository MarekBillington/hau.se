import { $, component$, QwikChangeEvent, Resource, useResource$, useStore } from "@builder.io/qwik";
import { DocumentHead, Link, useLocation } from '@builder.io/qwik-city';
import { request } from "~/components/api/api";
import Button from "~/components/inputs/button/button";
import Number from "~/components/inputs/number/number";
import Text from "~/components/inputs/text/text";
import House from '../house'

export default component$(() => {
    const location = useLocation();

    const store = useStore(
        {
            house: {} as House,
            isNewHouse: false,
        },
        {
            recursive: true
        }
    );

    const resHouse = useResource$(async () => {
        return request('house/' + location.params.houseId, 'GET', {})
    });

    const submit = $(async () => {
        console.log(JSON.stringify(store.house));
        console.log(store.house);

        const method = store.isNewHouse ? 'POST' : 'PATCH';
        const url = 'house/' + (store.isNewHouse ? '' : store.house.id);
        return request(url, method, JSON.stringify(store.house));
    });

    const onChange = $((event: QwikChangeEvent<HTMLInputElement>) => {
        type keyType = keyof typeof store.house;
        const n: keyType = event.target.name
        store['house'][n] = event.target.value;
    })

    return (
        <>
            <Link href="/house">Back</Link>
            <br />
            <Resource 
                value={resHouse}
                onResolved={(house: House) => {
                    store.house = house;
                    if (house.id === 0) {
                        store.isNewHouse = true;
                    }
                    return (
                        <div>
                            <h2>{house.address}</h2>
                            <div class="house-form"> 
                                {/* @todo https://qwik.builder.io/docs/components/state/#using-context */}

                                Address: <Text 
                                        value={store.house.address}
                                        name="address"
                                        change={onChange}
                                    />
                                <br/>
                                Bedrooms: <Number value={store.house.bedrooms} name="bedrooms" change={onChange}/>
                                <br/>
                                Bathroom: <Number value={store.house.bathrooms}  name="bathrooms" change={onChange}/>
                                <br/>
                                Garage: <Number value={store.house.garage}  name="garage" change={onChange}/>
                                <br/>
                                Floorspace (m&sup2;): <Number value={store.house.floorspace}  name="floorspace" change={onChange}/>
                                <br/>
                                Landarea (m&sup2;): <Number value={store.house.landarea}  name="landarea" change={onChange}/>
                                <br/>
                                <Button value="Submit" click={submit}/>
                            </div>
                        </div>
                    );
                }}
            />
        </>
    );
})


export const head: DocumentHead = {
    title: 'Hause',
};
