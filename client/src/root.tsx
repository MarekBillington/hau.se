import {
  component$,
  createContextId,
  useBrowserVisibleTask$,
  useContextProvider,
  useStore,
  useStyles$,
  useTask$,
} from "@builder.io/qwik";
import {
  QwikCityProvider,
  RouterOutlet,
  ServiceWorkerRegister,
} from "@builder.io/qwik-city";
import { Auth } from "./components/auth/auth";
import { RouterHead } from "./components/router-head/router-head";

import globalStyles from "./global.css?inline";
export const authCtx = createContextId<Auth>("auth");
export const userSession = createContextId("userSession");

export default component$(() => {
  /**
   * The root of a QwikCity site always start with the <QwikCityProvider> component,
   * immediately followed by the document's <head> and <body>.
   *
   * Dont remove the `<head>` and `<body>` elements.
   */

  const store = useStore({
    token: "",
    expiry: 0,
  } as Auth);

  useContextProvider(authCtx, store);
  useTask$(({track}) => {
    track(() => store.token)
    console.log(store.token)
    console.log(store.expiry)
  })

  useStyles$(globalStyles);

  console.log("root called");

  useBrowserVisibleTask$(() => {
    console.log("rendering root");
  });

  return (
    <QwikCityProvider>
      <head>
        <meta charSet="utf-8" />
        <link rel="manifest" href="/manifest.json" />
        <RouterHead />
      </head>
      <body lang="en">
        <RouterOutlet />
        <ServiceWorkerRegister />
      </body>
    </QwikCityProvider>
  );
});
