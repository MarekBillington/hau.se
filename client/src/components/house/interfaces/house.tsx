

export default interface House {
  id: number;
  active: boolean;
  created_at: string;
  updated_at: string;
  portfolio_id: number;
  street_1: string;
  street_2: string;
  postcode: string;
  town: string;
  country_id: number;
  property_id: number;
  bedrooms: number;
  bathrooms: number;
  garage: number;
  floorspace: number;
  landarea: number;
}
