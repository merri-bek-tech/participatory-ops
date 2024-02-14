import { Box, SimpleGrid, Heading } from "@chakra-ui/react";
import SiteCard from "../components/SiteCard";
import { PhysicalLayerSummary, SiteSummary } from "../types";

export default function PhysicalLayer() {
  const layer: PhysicalLayerSummary = {
    sites: [
      {
        name: "Radish House",
        id: "radish-house",
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
      },
      {
        name: "Brunswick Library",
        id: "brunswick-library",
        racks: [
          {
            name: "Lilly Pilly",
            id: "lilly-pilly",
            dimensions: {
              power: { name: "Power", status: "planned" },
              compute: { name: "Compute", status: "planned" },
              storage: { name: "Storage", status: "planned" },
              dataLink: { name: "Data link", status: "planned" },
            },
          },
          {
            name: "Murmbal",
            id: "murmbal",
            dimensions: {
              power: { name: "Power", status: "planned" },
              compute: { name: "Compute", status: "planned" },
              storage: { name: "Storage", status: "planned" },
              dataLink: { name: "Data link", status: "planned" },
            },
          },
        ],
      },
      {
        name: "CERES",
        id: "ceres",
        racks: [
          {
            name: "Kingfisher",
            id: "kingfisher",
            dimensions: {
              power: { name: "Power", status: "planned" },
              compute: { name: "Compute", status: "planned" },
              storage: { name: "Storage", status: "planned" },
              dataLink: { name: "Data link", status: "planned" },
            },
          },
        ],
      },
      {
        name: "Glenroy Community Hub",
        id: "glenroy-hub",
        racks: [
          {
            name: "Woolip",
            id: "woolip",
            dimensions: {
              power: { name: "Power", status: "planned" },
              compute: { name: "Compute", status: "planned" },
              storage: { name: "Storage", status: "planned" },
              dataLink: { name: "Data link", status: "planned" },
            },
          },
          {
            name: "Wangnarra",
            id: "wangnarra",
            dimensions: {
              power: { name: "Power", status: "planned" },
              compute: { name: "Compute", status: "planned" },
              storage: { name: "Storage", status: "planned" },
              dataLink: { name: "Data link", status: "planned" },
            },
          },
        ],
      },
    ],
  };

  return (
    <Box>
      <Heading mb={4}>Physical Layer</Heading>
      <SimpleGrid columns={[1, 2, 3]} spacing={4}>
        {layer.sites.map((site: SiteSummary) => (
          <Box>
            <SiteCard {...site} key={site.id} />
          </Box>
        ))}
      </SimpleGrid>
    </Box>
  );
}
