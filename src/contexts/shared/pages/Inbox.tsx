import { Box, Grid, GridItem, Heading } from "@chakra-ui/react"
import UnknownComponentCard from "../components/UnknownComponentCard"

export default function Inbox() {
  return (
    <Box>
      <Heading>Inbox</Heading>
      <Grid templateColumns={["repeat(1, 1fr)", "repeat(1, 1fr)", "repeat(2, 1fr)"]} gap={4} mt={2}>
        {[1, 2, 3, 4, 5, 6, 7, 8, 9].map((_, i) => (
          <GridItem>
            <UnknownComponentCard status={i % 2 == 0 ? "online" : "offline"} uuid="8431928b-a906-40de-bae2-ab30dfe5e2e3" />
          </GridItem>
        ))}
      </Grid>
    </Box>
  )
}
