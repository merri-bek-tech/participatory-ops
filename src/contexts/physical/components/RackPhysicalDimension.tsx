import { Heading, Box, Text, BoxProps, useColorModeValue, HStack } from "@chakra-ui/react";
import { PhysicalDimensionStatus, PhysicalDimensionSummary } from "../types";

type StatusColor = "green" | "gray";

function PhysicalDimensionBox({ dashed, children, colour }: { dashed: boolean; colour: StatusColor; children: React.ReactNode }) {
  let boxProps: BoxProps = {
    p: 2,
  };

  if (dashed) {
    boxProps.borderColor = useColorModeValue("gray.500", "gray.600");
    boxProps.borderWidth = "3px";
    boxProps.borderStyle = "dashed";
  }

  let bgHue = useColorModeValue("300", "600");
  if (colour === "gray") bgHue = useColorModeValue("200", "500");

  return (
    <Box {...boxProps} bgColor={`${colour}.${bgHue}`}>
      {children}
    </Box>
  );
}

function colourForStatus(status: PhysicalDimensionStatus): StatusColor {
  const statusColorMap: Record<PhysicalDimensionStatus, StatusColor> = {
    active: "green",
    planned: "gray",
  };
  return statusColorMap[status];
}

export default function RackPhysicalDimension({ name, status }: PhysicalDimensionSummary) {
  return (
    <PhysicalDimensionBox dashed={status == "planned"} colour={colourForStatus(status)}>
      <HStack justify="space-between">
        <Heading size="xs" textTransform="uppercase">
          {name}
        </Heading>
        <Text fontSize="sm">{status}</Text>
      </HStack>
    </PhysicalDimensionBox>
  );
}
