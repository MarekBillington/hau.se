import { component$ } from '@builder.io/qwik';
import { DocumentHead } from '@builder.io/quik-city';

export default component$(() => {

    return (
        <>
            <p>some walls</p>
        </>
    );
});

export const head: DocumentHead = {
    title: 'Houses',
  };