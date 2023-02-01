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
            <a href={store.house.id.toString()}>
                {store.house.address}
            </a>
        </div>
    );
})

export default House;