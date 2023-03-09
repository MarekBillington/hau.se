import type Address from "./address";

export default interface House {
  id: number;
  active: boolean;
  created_at: string;
  updated_at: string;
  portfolio_id: number;
  property_id: number;
  address: Address;
  bedrooms: number;
  bathrooms: number;
  garage: number;
  floorspace: number;
  landarea: number;
}
