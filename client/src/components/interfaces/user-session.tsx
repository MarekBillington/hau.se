import Portfolio from "./portfolio";
import type User from "./user";

export default interface UserSession {
  portfolio: Portfolio,
  user: User
}
