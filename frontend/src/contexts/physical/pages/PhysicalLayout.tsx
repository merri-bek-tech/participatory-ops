import { Box } from "@chakra-ui/react";
import { LayerNav } from "../../shared";
import { Outlet } from "react-router-dom";

export default function PhysicalLayout() {
  return (
    <Box>
      <LayerNav activeLayer="physical" />
      <Outlet />
    </Box>
  );
}
