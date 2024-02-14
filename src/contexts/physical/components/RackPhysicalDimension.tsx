import { Heading, Text, HStack } from "@chakra-ui/react";
import { PhysicalDimensionSummary } from "../types";
import StatusBox from "./StatusBox";

export default function RackPhysicalDimension({ name, status }: PhysicalDimensionSummary) {
  return (
    <StatusBox status={status}>
      <HStack justify="space-between">
        <Heading size="xs" textTransform="uppercase">
          {name}
        </Heading>
        <Text fontSize="sm">{status}</Text>
      </HStack>
    </StatusBox>
  );
}
