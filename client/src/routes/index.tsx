import { component$ } from '@builder.io/qwik';
import type { DocumentHead } from '@builder.io/qwik-city';

export default component$(() => {
  return (
    <div>
      <h1>We might want to put information here....</h1>
      <p>but who knows</p>
    </div>
  );
});

export const head: DocumentHead = {
  title: 'Hau.se',
  meta: [
    {
      name: 'description',
      content: 'Hau.se title page',
    },
  ],
};
