import { useLoaderData } from "react-router-dom"
import { RackDetails } from "../types"
import { Box, Card, CardBody, Heading, Image, Stack, Text } from "@chakra-ui/react"

export default function RackDetails() {
  const { rack } = useLoaderData() as { rack: RackDetails }

  console.log(rack)

  return (
    <Box>
      <Card direction={{ base: "column", sm: "row" }} overflow="hidden">
        <Image objectFit="cover" maxW={{ base: "100%", sm: "200px" }} src="https://upload.wikimedia.org/wikipedia/commons/c/c3/My_Opera_Server.jpg" alt={`A photo of the ${rack.name} server rack`} />
        <Stack>
          <CardBody>
            <Heading size="lg">{rack.name}</Heading>

            <Text py="2">A server rack that helps power Merri-bek tech's web applications</Text>
          </CardBody>
        </Stack>
      </Card>
    </Box>
  )
}
