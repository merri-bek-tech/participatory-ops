// import { rackDetailsLoader, siteListLoader } from "./data"
import ApplicationLayout from "./pages/ApplicationLayout"
import AppList from "./pages/AppList"
// import RackDetails from "./pages/RackDetails"
// import SiteList from "./pages/SiteList"

export const routes = {
  path: "application",
  element: <ApplicationLayout />,
  children: [
    { path: "", element: <AppList /> },
    //   { path: "rack/:id", element: <RackDetails />, loader: rackDetailsLoader },
  ],
}
