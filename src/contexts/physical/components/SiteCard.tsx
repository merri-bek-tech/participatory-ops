import { Card, CardBody, CardHeader, Heading, Stack, StackDivider, useColorModeValue } from "@chakra-ui/react";
import RackCard from "./RackCard";

export default function SiteCard() {
  return (
    <Card bgColor={useColorModeValue("gray.100", "gray.900")}>
      <CardHeader>
        <Heading size="md">Radish House</Heading>
      </CardHeader>

      <CardBody pt={0}>
        <Stack divider={<StackDivider />} spacing="4">
          <RackCard />
        </Stack>
      </CardBody>
    </Card>
  );
}
