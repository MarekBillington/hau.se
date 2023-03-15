import { request } from "../../api/api"
import Auth from "../../auth/interface/auth"
import { Country } from "../interface/country"


export const getCountries = async (auth: Auth) => {
  const c = await request("utility/country", auth, "GET")
  const fc = c.map((obj: Country) => {
    return {
      value: obj.id,
      text: obj.name
    }
  })
  return fc
};

