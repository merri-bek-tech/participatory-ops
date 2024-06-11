import { Grid, GridItem } from "@chakra-ui/react"
import { useLoaderData } from "react-router-dom"
import { ApplicationSummary } from "../types"

export default function ApplicationList() {
  const apps: ApplicationSummary[] = [
    {
      id: "1",
      name: "Application 1",
    },
    {
      id: "2",
      name: "Application 2",
    },
  ]

  return (
    <Grid templateColumns={["repeat(1, 1fr)", "repeat(2m 1fr)", "repeat(3, 1fr)"]} gap={4}>
      {apps.map((app: ApplicationSummary) => (
        <GridItem width="100%" key={app.id}>
          Hello world
        </GridItem>
      ))}
    </Grid>
  )
}
