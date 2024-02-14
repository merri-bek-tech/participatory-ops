import { Box, Card, CardBody, CardHeader, HStack, Heading, Stack, Image, useColorModeValue } from "@chakra-ui/react";
import RackCard from "./RackCard";
import { SiteSummary } from "../types";
import Site from "/site.png";

export default function SiteCard({ name, racks }: SiteSummary) {
  return (
    <Box>
      <Card bgColor={useColorModeValue("gray.100", "gray.900")} maxWidth="300px">
        <CardHeader>
          <Heading size="md">
            <HStack justify="space-between">
              <span>{name}</span>
              <Image src={Site} alt="rack" height="30px" />
            </HStack>
          </Heading>
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
