import { useLoaderData } from "react-router-dom"
import { RackDetails } from "../types"
import { Box, Card, CardBody, Heading, Image, Stack, Text } from "@chakra-ui/react"
import RackPhysicalDimension from "../components/RackPhysicalDimension"
import RackLayout from "../components/RackLayout"

export default function RackDetails() {
  const { rack } = useLoaderData() as { rack: RackDetails }
  const { dimensions } = rack

  console.log(rack)

  return (
    <Box>
      <Card direction={{ base: "column", sm: "row" }} overflow="hidden">
        <Image objectFit="cover" maxW={{ base: "100%", sm: "200px" }} src="https://upload.wikimedia.org/wikipedia/commons/c/c3/My_Opera_Server.jpg" alt={`A photo of the ${rack.name} server rack`} />
        <Stack>
          <CardBody>
            <Heading size="lg">{rack.name}</Heading>

            <Text py="2">A server rack that helps power Merri-bek tech's web applications</Text>

            <Stack spacing={1}>
              <RackPhysicalDimension {...dimensions.power} />
              <RackPhysicalDimension {...dimensions.compute} />
              <RackPhysicalDimension {...dimensions.storage} />
              <RackPhysicalDimension {...dimensions.dataLink} />
            </Stack>
          </CardBody>
        </Stack>
      </Card>
      <Stack mt={4} direction="row" justifyContent="center">
        <RackLayout />
      </Stack>
    </Box>
  )
}
