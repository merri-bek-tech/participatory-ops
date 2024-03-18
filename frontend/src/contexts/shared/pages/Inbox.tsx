import { Box, Grid, GridItem, Heading } from "@chakra-ui/react"
import UnknownComponentCard, { UnknownComponent } from "../components/UnknownComponentCard"

export default function Inbox() {
  const data: UnknownComponent[] = [
    { status: "online", uuid: "8431928b-a906-40de-bae2-ab30dfe5e2e1" },
    { status: "offline", uuid: "8431928b-a906-40de-bae2-ab30dfe5e2e2" },
    { status: "online", uuid: "8431928b-a906-40de-bae2-ab30dfe5e2e3" },
    { status: "offline", uuid: "8431928b-a906-40de-bae2-ab30dfe5e2e4" },
    { status: "online", uuid: "8431928b-a906-40de-bae2-ab30dfe5e2e5" },
  ]

  return (
    <Box>
      <Heading>Inbox</Heading>
      <Grid templateColumns={["repeat(1, 1fr)", "repeat(1, 1fr)", "repeat(2, 1fr)"]} gap={4} mt={2}>
        {data.map((component) => (
          <GridItem key={component.uuid}>
            <UnknownComponentCard {...component} />
          </GridItem>
        ))}
      </Grid>
    </Box>
  )
}