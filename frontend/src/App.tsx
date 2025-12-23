import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import Login from "@/pages/Login";
import { Toaster } from "sonner";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Register from "@/pages/Register";
import Data from "@/pages/Home/Data";
import AuthContext, { isLogged } from "@/context/AuthContext";
import { SidebarProvider } from "@/components/ui/sidebar";
import ProtectedRoute from "@/components/ProtectedRoute";
import Swagger from "@/pages/Home/Swagger";

const queryClient = new QueryClient();

export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Toaster richColors position="top-center" />

      <AuthContext value={{ isLogged }}>
        <SidebarProvider>
          <BrowserRouter>
            <Routes>
              <Route index path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
              <Route
                path="/swagger"
                element={
                  <ProtectedRoute>
                    <Swagger />
                  </ProtectedRoute>
                }
              />
              <Route
                path="/"
                element={
                  <ProtectedRoute>
                    <Data />
                  </ProtectedRoute>
                }
              />
            </Routes>
          </BrowserRouter>
        </SidebarProvider>
      </AuthContext>
    </QueryClientProvider>
  );
}
