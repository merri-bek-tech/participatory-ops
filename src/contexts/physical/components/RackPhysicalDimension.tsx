import { Heading, Box, Text } from "@chakra-ui/react";
import { PhysicalDimensionSummary } from "../types";

function PhysicalDimensionBox({ planned, children }: { planned: boolean; children: React.ReactNode }) {
  if (planned) {
    return (
      <Box bgColor="gray.300" p={2} borderColor="gray.500" borderWidth="3px" borderStyle="dashed">
        {children}
      </Box>
    );
  }

  return <Box>{children}</Box>;
}

export default function RackPhysicalDimension({ name, status }: PhysicalDimensionSummary) {
  return (
    <PhysicalDimensionBox planned={status == "planned"}>
      <Heading size="xs" textTransform="uppercase">
        {name}
      </Heading>
      <Text fontSize="sm">{status}</Text>
    </PhysicalDimensionBox>
  );
}
