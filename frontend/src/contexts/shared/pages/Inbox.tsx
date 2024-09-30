import { Box, Grid, GridItem, Heading } from "@chakra-ui/react"
import ComponentStatusCard from "../../physical/components/ComponentStatusCard"
import { ComponentStatus } from "../../physical/types"
import Api from "../api"
import { useContext, useEffect, useState } from "react"
import { useInterval } from "usehooks-ts"
import { SchemeContext } from "../../schemes"

export default function Inbox() {
  const [components, setComponents] = useState<ComponentStatus[]>([])
  const scheme = useContext(SchemeContext)

  const pollApi = () => {
    const api = new Api().forScheme(scheme.id)

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
            <ComponentStatusCard {...component} />
          </GridItem>
        ))}
      </Grid>
    </Box>
  )
}
