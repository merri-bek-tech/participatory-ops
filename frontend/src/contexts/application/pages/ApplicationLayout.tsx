import { Box } from "@chakra-ui/react"
import { LayerNav } from "../../shared"
import { Outlet } from "react-router-dom"

export default function ApplicationLayout() {
  return (
    <Box>
      <LayerNav activeLayer="application" />
      <Outlet />
    </Box>
  )
}
