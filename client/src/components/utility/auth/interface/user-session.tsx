import type Portfolio from "../../../interfaces/portfolio";
import type User from "../../../interfaces/user";

export default interface UserSession {
  portfolio: Portfolio,
  user: User
}
