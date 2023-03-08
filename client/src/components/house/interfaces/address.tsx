
export default interface Address {
  id: number;
  active: boolean;
  created_at: string;
  updated_at: string;
  street_1: string;
  street_2: string;
  postcode: string;
  Town: string;
  country_id: number;
  house_id: number;
  default: boolean;
}
