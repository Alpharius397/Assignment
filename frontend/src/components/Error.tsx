import { AlertCircleIcon } from "lucide-react";
import { Alert, AlertTitle, AlertDescription } from "@/components/ui/alert";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";
import { AccessToken, RefreshToken } from "@/axios";

export default function Error() {
  const navigate = useNavigate();

  const refreshPage = () => {
    navigate(0);
  };

  const logout = () => {
    localStorage.removeItem(AccessToken);
    localStorage.removeItem(RefreshToken);
    navigate("/login", {
      replace: true,
      state: { success: "Logout Successful!" },
    });
  };

  return (
    <Alert variant="destructive" className="flex flex-col gap-4">
      <div className="flex items-center gap-2">
        <AlertCircleIcon />
        <div>
          <AlertTitle>Unable to process your request.</AlertTitle>
          <AlertDescription>
            Please verify your network connection and re-login.
          </AlertDescription>
        </div>
      </div>

      <AlertDialog>
        <AlertDialogTrigger asChild>
          <Button variant="destructive">Resolve</Button>
        </AlertDialogTrigger>

        <AlertDialogContent>
          <div className="">Confirm action?</div>
          <div className="flex justify-center gap-2">
            <AlertDialogCancel asChild>
              <Button variant="outline" onClick={refreshPage}>
                Refresh Page
              </Button>
            </AlertDialogCancel>

            <AlertDialogAction asChild>
              <Button variant="destructive" onClick={logout}>
                Logout
              </Button>
            </AlertDialogAction>
          </div>
        </AlertDialogContent>
      </AlertDialog>
    </Alert>
  );
}
