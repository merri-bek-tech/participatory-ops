import { Box, Card, CardBody, CardHeader, Heading, Stack, StackDivider, useColorModeValue } from "@chakra-ui/react";
import RackCard from "./RackCard";
import { SiteSummary } from "../types";

export default function SiteCard({ name, racks }: SiteSummary) {
  return (
    <Box>
      <Card bgColor={useColorModeValue("gray.100", "gray.900")} maxWidth="300px">
        <CardHeader>
          <Heading size="md">{name}</Heading>
        </CardHeader>

        <CardBody pt={0}>
          <Stack spacing="4">
            {racks.map((rack) => (
              <RackCard {...rack} key={rack.id} />
            ))}
          </Stack>
        </CardBody>
      </Card>
    </Box>
  );
}
