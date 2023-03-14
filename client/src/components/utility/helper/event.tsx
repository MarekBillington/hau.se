import { QwikChangeEvent } from "@builder.io/qwik";
import { setProperty } from "./types";

export type QCE = QwikChangeEvent<HTMLInputElement>;

// To handle setting new values to typescript objects on input events
export const onChange = (event: QCE, object: any) => {
  type keyType = keyof typeof object;
  
  const k: keyType = event.target.name;

  const val = isNaN(event.target.valueAsNumber)
    ? event.target.value
    : event.target.valueAsNumber;
  
  setProperty(object, k, val);
};

// To handle objects where we want to force the value to a number
// examples are "select"
export const onChangeNum = (event: QCE, object: any) => {
  type keyType = keyof typeof object;
  
  const k: keyType = event.target.name;

  const val = parseInt(event.target.value);
  
  setProperty(object, k, val);
};