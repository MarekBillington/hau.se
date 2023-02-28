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
import { refreshToken } from "./components/auth/auth";
import { Auth } from "./components/interfaces/auth";
import { UserSesstion } from "./components/interfaces/user-session";
import { RouterHead } from "./components/router-head/router-head";
import globalStyles from "./global.css?inline";

export const authCtx = createContextId<Auth>("auth");
export const userSession = createContextId<UserSesstion>("userSession");


export default component$(() => {
  /**
   * ------------------------------- DO NOT DELETE ---------------------------------
   * The root of a QwikCity site always start with the <QwikCityProvider> component,
   * immediately followed by the document's <head> and <body>.
   *
   * Dont remove the `<head>` and `<body>` elements.
   */
  useStyles$(globalStyles);
  const store = useStore({
      auth: {
        token: "",
        expiry: 0,
      } as Auth,
      userSession: {} as UserSesstion
    },
    { deep: true }
  );

  useContextProvider(authCtx, store.auth);
  useContextProvider(userSession, store.userSession);
  useTask$(({track}) => {
    // track top level propery over object property
    track(() => store.auth);
    track(() => store.userSession);
    console.log('root reload');
  })

  useBrowserVisibleTask$(() => {
    if (store.auth.token == "") {
      refreshToken(store.auth);
    }
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
