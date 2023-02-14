import { component$, useStylesScoped$ } from '@builder.io/qwik';
import { Link } from '@builder.io/qwik-city';
import styles from './header.css?inline';

export default component$(() => {
  useStylesScoped$(styles);
   
  return (
    <header>
      <div class="logo">
        <Link href='/'>
          <h1>&#127968; Hause</h1>
        </Link>
      </div>
      <ul>
        <li>
          <Link href='/house'>
            Houses
          </Link>
        </li>
        <li>
          <Link href='/tenant'>
            Tenants
          </Link>
        </li>
        <li>
          <Link href='/report'>
            Reports
          </Link>
        </li>
        <li>
          <Link href='/admin'>
            Admin
          </Link>
        </li>
      </ul>
    </header>
  );
});
