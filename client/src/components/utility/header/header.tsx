import { $, component$, useContext, useStylesScoped$ } from "@builder.io/qwik";
import { Link, useNavigate } from "@builder.io/qwik-city";
import { authCtx, userSession } from "~/root";
import { logout } from "../auth/auth";
import Button from "../inputs/button/button";
import styles from "./header.css?inline";

export default component$(() => {
  const auth = useContext(authCtx);
  const sess = useContext(userSession);
  const nav = useNavigate()

  useStylesScoped$(styles);

  const doLogout = $(() => { 
    logout(auth, sess)
    nav("/")
  })
  
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
          <Button
            value="Logout"
            click={doLogout}
          />
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
