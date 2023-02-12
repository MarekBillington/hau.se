import { component$, useStore, $ } from "@builder.io/qwik";
import { request } from "~/components/api/api";

export default component$(() => {
    const state = useStore({
        address: "",
        bedrooms: 1,
        bathrooms: 1,
        garage: 0,
    })

    const submit = $(async () => {
        const res = request('/house', 'POST', JSON.stringify(state))
        console.log(res)
        return (
            <>
                {/* close popup or error */}
            </>
        )
    })


    return (
        <div>
            Address: <input type="text"></input><br/>
            Bedrooms: <input type="number" min={1}></input><br/>
            Bathroom: <input type="number" min={1}></input><br/>
            Garage: <input type="number" min={0}></input><br/>
            Floorspace (m^2): <input type="number" min={0}></input><br/>
            Landarea (m^2): <input type="number" min={0}></input><br/>
            <input type="button" onClick$={submit} />
        </div>
    )
})