import { Grid, GridItem } from "@chakra-ui/react"
import SiteCard from "../components/SiteCard"
import { SiteSummary } from "../types"
import { useLoaderData } from "react-router-dom"

export default function SiteList() {
  const { sites } = useLoaderData() as { sites: SiteSummary[] }

  return (
    <Grid templateColumns={["repeat(1, 1fr)", "repeat(2m 1fr)", "repeat(3, 1fr)"]} gap={4}>
      {sites.map((site: SiteSummary) => (
        <GridItem width="100%" key={site.id} colSpan={Math.min(site.racks.length, 3)}>
          <SiteCard {...site} key={site.id} />
        </GridItem>
      ))}
    </Grid>
  )
}
