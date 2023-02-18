import { $, component$, QwikChangeEvent, Resource, useResource$, useStore } from "@builder.io/qwik";
import { DocumentHead, Link, useLocation, useNavigate } from '@builder.io/qwik-city';
import { request } from "~/components/api/api";
import { setProperty } from "~/components/common/types";
import Button from "~/components/inputs/button/button";
import Number from "~/components/inputs/number/number";
import Text from "~/components/inputs/text/text";
import House from '../house'

export default component$(() => {
    const location = useLocation();
    const nav = useNavigate();

    const store = useStore(
        {
            house: {
                id: 0,
                address: '',
                bedrooms: 0,
                bathrooms: 0,
                garage: 0,
                floorspace: 0,
                landarea: 0
            } as House,
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
        const method = store.isNewHouse ? 'POST' : 'PATCH';
        const url = 'house' + (store.isNewHouse ? '' : '/' + store.house.id);
        
        const h = await request(url, method, JSON.stringify(store.house));
        
        // @todo can make smarter so it doesn't go back to house page
        let refresh = '/house'
        if (!isNaN(h.id)) {
            refresh = refresh + h.id
        }
        nav(refresh)
    });

    const disable = $(async () => {
        const url = 'house' + (store.isNewHouse ? '' : '/' + store.house.id);
        
        const h = await request(url, 'DELETE');
        nav('/house')
    })

    const onChange = $((event: QwikChangeEvent<HTMLInputElement>) => {
        type keyType = keyof (typeof store.house);
        // @ts-ignore complaining about it being asignable, but it works lmao
        const k: keyType = event.target.name;
        
        const val = isNaN(event.target.valueAsNumber) ? event.target.value : event.target.valueAsNumber;
        
        setProperty(store.house, k, val);
    })

    return (
        <>
            <div class="toolbar">
                <div class="toolbar-left">
                    <Link href="/house">Back</Link>
                </div>
                <div class="toolbar-right">
                    <Button value="Disable" click={disable} />
                </div>
            </div>
            <div>
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
                                    <Text label="Address:" value={store.house.address} name="address" change={onChange}/>
                                    <br/>
                                    <Number label="Bedrooms:" value={store.house.bedrooms} name="bedrooms" change={onChange}/>
                                    <br/>
                                    <Number label="Bathroom:" value={store.house.bathrooms} name="bathrooms" change={onChange}/>
                                    <br/>
                                    <Number label="Garage:" value={store.house.garage} name="garage" change={onChange}/>
                                    <br/>
                                    <Number label="Floorspace (m&sup2;):" value={store.house.floorspace} name="floorspace" change={onChange}/>
                                    <br/>
                                    <Number label="Landarea (m&sup2;):" value={store.house.landarea} name="landarea" change={onChange}/>
                                    <br/>
                                    <Button value="Save" click={submit} />
                                </div>
                            </div>
                        );
                    }}
                />
            </div>
        </>
    );
})


export const head: DocumentHead = {
    title: 'Hause',
};
