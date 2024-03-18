import { Navigate, createBrowserRouter } from "react-router-dom"
import { routes as physicalRoutes } from "./contexts/physical"
import { Layout, Inbox } from "./contexts/shared"
import { VirtualLayer } from "./contexts/virtual"

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [{ path: "", element: <Navigate to="/physical" replace /> }, { path: "virtual", element: <VirtualLayer /> }, physicalRoutes, { path: "inbox", element: <Inbox /> }],
  },
])
