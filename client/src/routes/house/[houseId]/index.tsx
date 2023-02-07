import { $, component$, Resource, useResource$ } from "@builder.io/qwik";
import { useLocation } from '@builder.io/qwik-city';
import House from '../house'

export default component$(() => {
    const location = useLocation();

    const house = useResource$(async () => {
        const res = await fetch("http://localhost:8000/house/"+location.params.houseId)
        console.log(res)
        return res.json()
    })

    const submit  = $(async () => {
        // @todo WHY WONT YOU WORK? >> stupid CORS
        const res = await fetch("http://localhost:8000/house/"+location.params.houseId, {
            headers: {
                'Access-Control-Allow-Origin': '*'
            },
            mode:"no-cors",
            method: 'PUT'
        })
        console.log(res.json)
    })

    return (
        <>
            <a href="/house">Back</a>
            <br />
            <Resource 
                value={house}
                onResolved={(house: House) => {
                    console.log(house)
                    return (
                        <div>
                            <h2>{house.address}</h2>
                            <div>
                                Address: <input type="text"></input>
                                Bedrooms: <input type="number" min={1}></input>
                                Bathroom: <input type="number" min={1}></input>
                                Garage: <input type="number" min={0}></input>
                                Floorspace (m^2): <input type="number" min={0}></input>
                                Landarea (m^2): <input type="number" min={0}></input>
                                <input type="button" onClick$={submit} title="Submit"></input>
                            </div>
                        </div>
                    )
                }}
            />
        </>
    )
})