import { QwikChangeEvent } from "@builder.io/qwik";
import { setProperty } from "./types";

export type QCE = QwikChangeEvent<HTMLInputElement>;

export const onChange = (event: QwikChangeEvent<HTMLInputElement>, object: any) => {
  type keyType = keyof typeof object;
  
  const k: keyType = event.target.name;

  const val = isNaN(event.target.valueAsNumber)
    ? event.target.value
    : event.target.valueAsNumber;
  
  console.log(object)
  console.log(val)
  setProperty(object, k, val);
};
