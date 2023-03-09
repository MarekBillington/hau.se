//import { $ } from "@builder.io/qwik";

// See: https://www.typescriptlang.org/docs/handbook/release-notes/typescript-2-1.html#keyof-and-lookup-types
export function getProperty<T, K extends keyof T>(obj: T, key: K) {
  return obj[key]; // Inferred type is T[K]
}

export function setProperty<T, K extends keyof T>(obj: T, key: K, value: T[K]) {
  obj[key] = value;
}

export const setProperty2 = <T, K extends keyof T>(obj: T, key: K, value: T[K]) => {
  type t = typeof obj
  (obj as t)[key] = value
}
