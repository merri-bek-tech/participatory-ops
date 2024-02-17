import { Box, Grid, GridItem, Heading, Text } from "@chakra-ui/react";
import SiteCard from "../components/SiteCard";
import { PhysicalLayerSummary, SiteSummary } from "../types";
import { LayerNav, Link } from "../../shared";

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
              power: { name: "Power", status: "active" },
              compute: { name: "Compute", status: "active" },
              storage: { name: "Storage", status: "active" },
              dataLink: { name: "Data link", status: "active" },
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
              power: { name: "Power", status: "active" },
              compute: { name: "Compute", status: "error" },
              storage: { name: "Storage", status: "warning" },
              dataLink: { name: "Data link", status: "active" },
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
      <LayerNav activeLayer="physical" />
      <Grid templateColumns={["repeat(1, 1fr)", "repeat(2m 1fr)", "repeat(3, 1fr)"]} gap={4}>
        {layer.sites.map((site: SiteSummary) => (
          <GridItem width="100%" key={site.id} colSpan={Math.min(site.racks.length, 3)}>
            <SiteCard {...site} key={site.id} />
          </GridItem>
        ))}
      </Grid>
    </Box>
  );
}
