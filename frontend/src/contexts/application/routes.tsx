// import { rackDetailsLoader, siteListLoader } from "./data"
import ApplicationLayout from "./pages/ApplicationLayout"
import ApplicationList from "./pages/ApplicationList"
// import RackDetails from "./pages/RackDetails"
// import SiteList from "./pages/SiteList"

export const routes = {
  path: "application",
  element: <ApplicationLayout />,
  children: [
    { path: "", element: <ApplicationList /> },
    //   { path: "rack/:id", element: <RackDetails />, loader: rackDetailsLoader },
  ],
}
