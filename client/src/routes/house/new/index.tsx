import { component$, useStore, $ } from "@builder.io/qwik";

export default component$(() => {
    const state = useStore({
        address: "",
        bedrooms: 1,
        bathrooms: 1,
        garage: 0,
    })

    const submit = $(async () => {
        let res = await fetch('http://localhost:8000/house', {
            method: 'POST',
            body: JSON.stringify(state)
        })
        console.log(res)

        return (
            <>
                {/* close popup or error */}
            </>
        )
    })


    return (
        <div>
            Address: <input type="text"></input>
            Bedrooms: <input type="number" min={1}></input>
            Bathroom: <input type="number" min={1}></input>
            Garage: <input type="number" min={0}></input>
            Floorspace (m^2): <input type="number" min={0}></input>
            Landarea (m^2): <input type="number" min={0}></input>
            <input type="button" onClick$={submit} />
        </div>
    )
})