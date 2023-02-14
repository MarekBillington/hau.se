import { component$, Resource, useResource$ } from '@builder.io/qwik';
import { DocumentHead, Link } from '@builder.io/qwik-city';
import { request } from '~/components/api/api';
import House from './house';

export const Houses = component$(() => {

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
                                return <House {...h}/>
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

export default Houses;

export const head: DocumentHead = {
    title: 'Hause'
};
