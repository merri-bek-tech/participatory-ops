import { Card, CardHeader, Heading, Stack, Image, Text, HStack } from "@chakra-ui/react";
import RackPhysicalDimension from "./RackPhysicalDimension";
import { RackSummary } from "../types";
import Rack from "/rack.png";

export default function RackCard({ name, dimensions }: RackSummary) {
  return (
    <Card>
      <CardHeader>
        <Heading size="md">
          <HStack justify="space-between">
            <span>{name}</span>
            <Image src={Rack} alt="rack" height="30px" />
          </HStack>
        </Heading>
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
