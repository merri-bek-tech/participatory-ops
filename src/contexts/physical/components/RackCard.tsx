import { Card, CardBody, CardHeader, Heading, Stack, StackDivider } from "@chakra-ui/react";
import RackPhysicalDimension from "./RackPhysicalDimension";

export default function RackCard() {
  return (
    <Card>
      <CardHeader>
        <Heading size="md">Murnong</Heading>
      </CardHeader>

      <Stack spacing="1px">
        <RackPhysicalDimension name="Power" />
        <RackPhysicalDimension name="Compute" />
        <RackPhysicalDimension name="Storage" />
        <RackPhysicalDimension name="Data link" />
      </Stack>
    </Card>
  );
}
