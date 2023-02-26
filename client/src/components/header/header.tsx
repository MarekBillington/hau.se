import { component$, useContext, useStylesScoped$ } from "@builder.io/qwik";
import { Link } from "@builder.io/qwik-city";
import { authCtx } from "~/root";
import styles from "./header.css?inline";

export default component$(() => {
  const auth = useContext(authCtx);

  useStylesScoped$(styles);

  let links;
  if (auth.token != "") {
    links = (
      <ul>
        <li>
          <Link href="/house">Houses</Link>
        </li>
        <li>
          <Link href="/tenant">Tenants</Link>
        </li>
        <li>
          <Link href="/report">Reports</Link>
        </li>
        <li>
          <Link href="/admin">Admin</Link>
        </li>
      </ul>
    );
  } else {
    links = (
      <ul>
        <li>
          <Link href="/login">Login</Link>
        </li>
        <li>
          <Link href="/signup">Signup</Link>
        </li>
      </ul>
    );
  }

  return (
    <header>
      <div class="logo">
        <Link href="/">
          <h1>&#127968; Hause</h1>
        </Link>
      </div>
      {links}
    </header>
  );
});
