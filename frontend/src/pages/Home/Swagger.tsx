import { URL } from "@/axios";
import SwaggerUI from "swagger-ui-react";
import "swagger-ui-react/swagger-ui.css";
import Sidebar from "@/pages/Home/SideBar";
import { SidebarTrigger } from "@/components/ui/sidebar";

export default function Swagger() {
  return (
    <div className="flex w-full m-0 relative">
      <Sidebar />
      <div className="m-2">
        <SidebarTrigger />
      </div>
      <div className="w-full">
        <SwaggerUI url={URL.Swagger} />
      </div>
    </div>
  );
}
