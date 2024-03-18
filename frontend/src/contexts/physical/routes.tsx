import { rackDetailsLoader, siteListLoader } from "./data"
import PhysicalLayout from "./pages/PhysicalLayout"
import RackDetails from "./pages/RackDetails"
import SiteList from "./pages/SiteList"

export const routes = {
  path: "physical",
  element: <PhysicalLayout />,
  children: [
    { path: "", element: <SiteList />, loader: siteListLoader },
    { path: "rack/:id", element: <RackDetails />, loader: rackDetailsLoader },
  ],
}
