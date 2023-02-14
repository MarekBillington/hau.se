import {component$} from '@builder.io/qwik'
import { DocumentHead } from '@builder.io/qwik-city';

export default component$(() => {
    return (
        <div>
            Look at me, king of the castle
        </div>
    )
})


export const head: DocumentHead = {
    title: 'Hause'
};
