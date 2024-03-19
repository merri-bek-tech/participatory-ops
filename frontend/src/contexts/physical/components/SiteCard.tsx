import { Card, CardBody, CardHeader, HStack, Heading, Image, useColorModeValue, Grid, GridItem } from "@chakra-ui/react"
import RackCard from "./RackCard"
import { SiteSummary } from "../types"
import Site from "../../../assets/site.png"

export default function SiteCard({ name, racks }: SiteSummary) {
  const multiRack = racks.length > 1

  return (
    <Card bgColor={useColorModeValue("gray.100", "gray.900")}>
      <CardHeader>
        <Heading size="md">
          <HStack justify="space-between">
            <span>{name}</span>
            <Image src={Site} alt="rack" height="30px" filter={useColorModeValue("", "invert(100%)")} />
          </HStack>
        </Heading>
      </CardHeader>

      <CardBody pt={0}>
        {multiRack ? (
          <Grid templateColumns={["repeat(1, 1fr)", `repeat(${Math.min(2, racks.length)} 1fr)`, `repeat(${Math.min(3, racks.length)}, 1fr)`]} gap={4}>
            {racks.map((rack) => (
              <GridItem key={rack.id}>
                <RackCard {...rack} />
              </GridItem>
            ))}
          </Grid>
        ) : (
          racks.map((rack) => <RackCard {...rack} key={rack.id} />)
        )}
      </CardBody>
    </Card>
  )
}
