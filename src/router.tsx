import { createBrowserRouter } from "react-router-dom";
import { PhysicalLayer } from "./contexts/physical";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <PhysicalLayer />,
  },
]);
