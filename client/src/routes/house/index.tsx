import { component$, Resource, useResource$ } from '@builder.io/qwik';
import House from './house';

export const Houses = component$(() => {

    const houseList = useResource$(async () => {
        const res = await fetch('http://localhost:8001/house');
        console.log(res)
        return res.json();
    })

    return (
        <div>
            <div>
                <a href="/house/new">Add House</a>
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
