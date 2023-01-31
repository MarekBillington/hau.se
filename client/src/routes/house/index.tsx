import { component$, Resource, useResource$ } from '@builder.io/qwik';
import { DocumentHead } from '@builder.io/quik-city';
import House from './house';

export const Houses = component$(() => {

    // const houseList = useResource$(async () => {
    //     const res = await fetch('http://localhost:8000/house');
    //     return res.json();
    // })

    return (
        <div>
            {/* <Resource 
                value={houseList}
                onResolved={(houses: Array<House>) => {
                    const hs = houses.map(h => {
                        return <House {...h}/>
                    });
                    return (
                        <>
                            {hs}
                        </>
                    )
                }}
            /> */}
            Hello
        </div>
    );
});

export default Houses;
