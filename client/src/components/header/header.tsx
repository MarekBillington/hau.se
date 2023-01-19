import { component$, useStylesScoped$ } from '@builder.io/qwik';
import { QwikLogo } from '../icons/qwik';
import styles from './header.css?inline';

export default component$(() => {
  useStylesScoped$(styles);

  return (
    <header>
      <div class="logo">
        <a href="/" >
          <h1>&#127968; Hause</h1>
        </a>
      </div>
      <ul>
        <li>
          <a href="/house">
            Houses
          </a>
        </li>
        <li>
          <a href="/tenant">
            Tenants
          </a>
        </li>
        <li>
          <a href="/report">
            Reports
          </a>
        </li>
      </ul>
    </header>
  );
});
