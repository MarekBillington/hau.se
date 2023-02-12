import { $, component$, Resource, useResource$ } from "@builder.io/qwik";
import { Link, useLocation } from '@builder.io/qwik-city';
import { request } from "~/components/api/api";
import House from '../house'

export default component$(() => {
    const location = useLocation();

    const house = useResource$(async () => {
        return request('house/' + location.params.houseId, 'GET', {})
    })

    const submit  = $(async () => {
        return request('house/' + location.params.houseId, 'PATCH', {})
    })

    return (
        <>
            <Link href="/house">Back</Link>
            <br />
            <Resource 
                value={house}
                onResolved={(house: House) => {
                    console.log(house)
                    return (
                        <div>
                            <h2>{house.address}</h2>
                            <div>
                                Address: <input type="text" value={house.address}></input>
                                <br />
                                Bedrooms: <input type="number" min={1} value={house.bedrooms}></input>
                                <br />
                                Bathroom: <input type="number" min={1} value={house.bedrooms}></input>
                                <br />
                                Garage: <input type="number" min={0} value={house.garage}></input>
                                <br />
                                Floorspace (m^2): <input type="number" min={0} value={house.floorspace}></input>
                                <br />
                                Landarea (m^2): <input type="number" min={0} value={house.landarea}></input>
                                <br />
                                <input type="button" onClick$={submit} title="Submit"></input>
                            </div>
                        </div>
                    )
                }}
            />
        </>
    )
})