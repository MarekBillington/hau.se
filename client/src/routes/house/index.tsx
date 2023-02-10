import { component$, Resource, useResource$ } from '@builder.io/qwik';
import House from './house';

export const Houses = component$(() => {

    const houseList = useResource$(async () => {
        try {
            const res = await fetch('http://hause-server-1:8001/api/house');
            console.log(res)
            return await res.json();
        } catch (err) {
            console.log(err)
        }
    })

    return (
        <div>
            <div>
                <a href="/house/new">Add New House</a>
            </div>
            <div>
                <Resource 
                    value={houseList}
                    onResolved={(houses: Array<House>) => {
                        console.log(houses)    
                        
                        const hs = houses.map(h => {
                            return <House {...h}/>
                        });
                        return (
                            <>
                                {hs}
                            </>
                        )
                    }}
                />
            </div>
        </div>
    );
});

export default Houses;
