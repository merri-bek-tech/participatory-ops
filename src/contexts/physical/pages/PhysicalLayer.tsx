import { Box, Grid, Heading } from "@chakra-ui/react";
import SiteCard from "../components/SiteCard";
import { SiteSummary } from "../types";

export default function PhysicalLayer() {
  const site: SiteSummary = {
    name: "Radish House",
    racks: [
      {
        name: "Murnong",
        id: "murnong",
        dimensions: {
          power: { name: "Power", status: "planned" },
          compute: { name: "Compute", status: "planned" },
          storage: { name: "Storage", status: "planned" },
          dataLink: { name: "Data link", status: "planned" },
        },
      },
    ],
  };

  return (
    <Box>
      <Heading mb={4}>Physical Layer</Heading>
      <Grid templateColumns="repeat(2, 1fr)" gap={4}>
        <SiteCard {...site} />
      </Grid>
    </Box>
  );
}
