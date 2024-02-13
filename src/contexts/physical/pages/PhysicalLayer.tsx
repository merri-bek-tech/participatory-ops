import { Box, Grid, Heading } from "@chakra-ui/react";
import SiteCard from "../components/SiteCard";

export default function PhysicalLayer() {
  return (
    <Box>
      <Heading mb={4}>Physical Layer</Heading>
      <Grid templateColumns="repeat(2, 1fr)" gap={4}>
        <SiteCard />
      </Grid>
    </Box>
  );
}
