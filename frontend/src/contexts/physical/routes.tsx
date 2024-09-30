import { rackDetailsLoader, siteListLoader } from "./data"
import PhysicalLayout from "./pages/PhysicalLayout"
import RackDetails from "./pages/RackDetails"
import SiteList from "./pages/SiteList"
import NodeOverview from "./pages/NodeOverview"
import { Navigate, redirect } from "react-router-dom"

export const routes = {
  path: "physical",
  element: <PhysicalLayout />,
  children: [
    { path: "", element: <Navigate to="nodes/this" /> },
    { path: "nodes/this", element: <NodeOverview /> },
    { path: "sites", element: <SiteList />, loader: siteListLoader },
    { path: "rack/:id", element: <RackDetails />, loader: rackDetailsLoader },
  ],
}
