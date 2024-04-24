// import { rackDetailsLoader, siteListLoader } from "./data"
import ApplicationLayout from "./pages/ApplicationLayout"
// import RackDetails from "./pages/RackDetails"
// import SiteList from "./pages/SiteList"

export const routes = {
  path: "application",
  element: <ApplicationLayout />,
  children: [
    { path: "", element: <div>app layer</div> },
    //   { path: "", element: <SiteList />, loader: siteListLoader },
    //   { path: "rack/:id", element: <RackDetails />, loader: rackDetailsLoader },
  ],
}
