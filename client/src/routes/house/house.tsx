import { component$, useStore, useStylesScoped$ } from '@builder.io/qwik';
import { Link } from '@builder.io/qwik-city';
import styles from './house.css?inline';

interface House {
    id: number,
    address: string,
    bedrooms: number,
    bathrooms: number,
    garage: number,
    floorspace: number,
    landarea: number,
}

export const House = component$((house: House) => {
    useStylesScoped$(styles);

    const store = useStore(
        {house},
        {recursive: true}
    );

    return (
        <div class="house-tile">
            <div></div>
            <Link href={store.house.id.toString()}>
                {store.house.address}
            </Link>
        </div>
    );
})

export default House;
