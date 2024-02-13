import { createBrowserRouter } from "react-router-dom";
import { PhysicalLayer } from "./contexts/physical";
import { Layout } from "./contexts/shared";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [{ path: "", element: <PhysicalLayer /> }],
  },
]);
