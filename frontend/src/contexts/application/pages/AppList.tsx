import { Grid, GridItem } from "@chakra-ui/react"
import { useLoaderData } from "react-router-dom"
import { AppSummary } from "../types"
import AppCard from "../components/AppCard"

export default function AppList() {
  const apps: AppSummary[] = [
    {
      id: "1",
      name: "Wikipedia",
    },
    {
      id: "2",
      name: "Docuwiki",
    },
    {
      id: "3",
      name: "Matrix",
    },
  ]

  return (
    <Grid templateColumns={["repeat(1, 1fr)", "repeat(2m 1fr)", "repeat(3, 1fr)"]} gap={4}>
      {apps.map((app: AppSummary) => (
        <GridItem width="100%" key={app.id}>
          <AppCard {...app} />
        </GridItem>
      ))}
    </Grid>
  )
}
