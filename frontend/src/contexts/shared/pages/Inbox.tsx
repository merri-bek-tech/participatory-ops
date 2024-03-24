import { Box, Grid, GridItem, Heading } from "@chakra-ui/react"
import UnknownComponentCard from "../components/UnknownComponentCard"
import { UnknownComponent } from "../types"
import Api from "../api"
import { useEffect, useState } from "react"
import { useInterval } from "usehooks-ts"

const api = new Api()

export default function Inbox() {
  const [components, setComponents] = useState<UnknownComponent[]>([])

  const pollApi = () => {
    api.inbox().then((data) => {
      setComponents(data)
      console.log(data)
    })
  }

  useInterval(pollApi, 2000)
  useEffect(pollApi, [])

  return (
    <Box>
      <Heading>Inbox</Heading>
      <Grid templateColumns={["repeat(1, 1fr)", "repeat(1, 1fr)", "repeat(2, 1fr)"]} gap={4} mt={2}>
        {components.map((component) => (
          <GridItem key={component.uuid}>
            <UnknownComponentCard {...component} />
          </GridItem>
        ))}
      </Grid>
    </Box>
  )
}
