import { Navigate, createBrowserRouter } from "react-router-dom"
import { routes as physicalRoutes } from "./contexts/physical"
import { routes as applicationRoutes } from "./contexts/application"
import { Layout, Inbox } from "./contexts/shared"
import { VirtualLayer } from "./contexts/virtual"
import { WithScheme } from "./contexts/schemes"
import BaseLayout from "./contexts/shared/pages/BaseLayout"

export const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <BaseLayout>
        <WithScheme>
          <Layout />
        </WithScheme>
      </BaseLayout>
    ),
    children: [{ path: "", element: <Navigate to="/physical" replace /> }, applicationRoutes, { path: "virtual", element: <VirtualLayer /> }, physicalRoutes, { path: "inbox", element: <Inbox /> }],
  },
])
