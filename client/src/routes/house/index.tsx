import { component$, Resource, useResource$, useStore } from '@builder.io/qwik';
import Popup from '~/components/popup/popup';
import House from './house';

export const Houses = component$(() => {
    const state = useStore({
        isOpen: false
    })

    const houseList = useResource$(async () => {
        const res = await fetch('http://localhost:8000/house');
        console.log(res)
        return res.json();
    })

    function renderNewHouse() {
        state.isOpen = true;
    }

    const handlePopup = () => {
        state.isOpen = false;
    }

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
