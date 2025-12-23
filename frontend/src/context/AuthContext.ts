import { AccessToken, RefreshToken } from "@/axios";
import { createContext } from "react";

type AuthUser = {
  isLogged: () => boolean;
};

// simple context to ensure user is logged in
export function isLogged() {
  return (
    localStorage.getItem(AccessToken) !== null &&
    localStorage.getItem(RefreshToken) !== null
  );
}

const AuthContext = createContext<AuthUser>({ isLogged });

export default AuthContext;
