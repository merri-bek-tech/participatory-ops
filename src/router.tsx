import { Navigate, createBrowserRouter } from "react-router-dom";
import { PhysicalLayout, SiteList, RackDetails } from "./contexts/physical";
import { Layout } from "./contexts/shared";
import { VirtualLayer } from "./contexts/virtual";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      { path: "", element: <Navigate to="/physical" replace /> },
      { path: "virtual", element: <VirtualLayer /> },
      {
        path: "physical",
        element: <PhysicalLayout />,
        children: [
          { path: "", element: <SiteList /> },
          { path: "rack/:id", element: <RackDetails /> },
        ],
      },
    ],
  },
]);
