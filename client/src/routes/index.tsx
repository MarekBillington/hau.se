import { component$ } from '@builder.io/qwik';
import type { DocumentHead } from '@builder.io/qwik-city';

export const App = component$(() => {
  
  return (
    <>
      Eventually we want upfront metrics over there somewhere...
    </>
  );
});

export default App;

export const head: DocumentHead = {
  title: 'Hau.se',
  meta: [
    {
      name: 'description',
      content: 'Hau.se title page',
    },
  ],
};
