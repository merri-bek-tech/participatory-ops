import { Box, Heading, Text } from "@chakra-ui/react";
import { LayerNav, Link } from "../../shared";

export default function VirtualLayer() {
  return (
    <Box>
      <LayerNav activeLayer="virtual" />
    </Box>
  );
}
