import { component$, Resource, useResource$, useStylesScoped$ } from '@builder.io/qwik';
import { DocumentHead, Link } from '@builder.io/qwik-city';
import { request } from '~/components/api/api';
import House from './house';
import styles from './house.css?inline';

export default component$(() => {

    const houseList = useResource$(async () => {
        return request('house', 'GET', {})
    })

    return (
        <div>
            <div>
                <Link href='/house/0/'>Add New House</Link>
            </div>
            <div>
                <Resource 
                    value={houseList}
                    onResolved={(houses: Array<House>) => {
                        let hs = {}
                        if (Array.isArray(houses)) {
                            hs = houses.map(h => {
                                return <Panel {...h}/>
                            });
                        }
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

export const Panel = component$((house: House) => {
    useStylesScoped$(styles);
    
    return (
        <div class="house-tile">
            <div></div>
            <Link href={house.id.toString()}>
                {house.address}
            </Link>
        </div>
    );
})

export const head: DocumentHead = {
    title: 'Hause'
};
