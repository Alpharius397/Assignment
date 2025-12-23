import { Navigate } from "react-router-dom";
import AuthContext from "@/context/AuthContext";
import { useContext } from "react";

export default function ProtectedRoute(
  props: React.PropsWithChildren,
): React.ReactNode {
  const { isLogged } = useContext(AuthContext);

  if (!isLogged()) {
    return (
      <Navigate
        to={"/login"}
        state={{ error: "Unauthorized Entry detected! Please re-login" }}
        replace
      />
    );
  }

  return props.children;
}
