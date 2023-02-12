import { component$, useStore } from '@builder.io/qwik';
import { Link } from '@builder.io/qwik-city';

interface House {
    id: number,
    address: string,
    bedrooms: string,
    bathrooms: string,
    garage: string,
    floorspace: string,
    landarea: string,
}

export const House = component$((house: House) => {
    const store = useStore(
        {house},
        {recursive: true}
    )
    return (
        <div>
            <Link href={store.house.id.toString()}>
                {store.house.address}
            </Link>
        </div>
    );
})

export default House;