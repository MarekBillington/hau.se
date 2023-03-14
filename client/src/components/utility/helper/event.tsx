import { QwikChangeEvent } from "@builder.io/qwik";
import { setProperty } from "./types";

export type QCE = QwikChangeEvent<HTMLInputElement>;

// To handle setting new values to typescript objects on input events
export const onChange = (event: QwikChangeEvent<HTMLInputElement>, object: any) => {
  type keyType = keyof typeof object;
  
  const k: keyType = event.target.name;

  const val = isNaN(event.target.valueAsNumber)
    ? event.target.value
    : event.target.valueAsNumber;
  
  setProperty(object, k, val);

  console.log(object)
  console.log(k)
  console.log(val)
};

export const onChangeNum = (event: QwikChangeEvent<HTMLInputElement>, object: any) => {
  type keyType = keyof typeof object;
  
  const k: keyType = event.target.name;

  const val = parseInt(event.target.value);
  
  setProperty(object, k, val);

  console.log(object)
  console.log(k)
  console.log(val)
};