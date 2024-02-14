import { Card, CardHeader, Heading, Stack } from "@chakra-ui/react";
import RackPhysicalDimension from "./RackPhysicalDimension";
import { RackSummary } from "../types";

export default function RackCard({ name, dimensions }: RackSummary) {
  return (
    <Card>
      <CardHeader>
        <Heading size="md">{name}</Heading>
      </CardHeader>

      <Stack spacing="1px">
        <RackPhysicalDimension {...dimensions.power} />
        <RackPhysicalDimension {...dimensions.compute} />
        <RackPhysicalDimension {...dimensions.storage} />
        <RackPhysicalDimension {...dimensions.dataLink} />
      </Stack>
    </Card>
  );
}
