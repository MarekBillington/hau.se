import { component$, useStore } from '@builder.io/qwik';

interface House {
    id: number,
    address: string,
}

export const House = component$((house: House) => {
    const store = useStore(
        {house},
        {recursive: true}
    )
    return (
        <div>
            <h6>{store.house.address}</h6>
        </div>
    );
})

export default House;