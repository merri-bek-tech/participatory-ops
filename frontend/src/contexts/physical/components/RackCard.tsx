import { Card, CardHeader, Heading, Stack, Image, Text, HStack, useColorModeValue, Button } from "@chakra-ui/react"
import RackPhysicalDimension from "./RackPhysicalDimension"
import { RackSummary } from "../types"
import Rack from "../../../assets/rack.png"
import { useNavigate } from "react-router-dom"

export default function RackCard({ name, id, detailUrl, dimensions }: RackSummary) {
  const navigate = useNavigate()

  return (
    <Card>
      <CardHeader p={1}>
        <Heading size="md" pl={2}>
          <HStack justify="space-between">
            <span>{name}</span>
            <Button colorScheme="cyan" isDisabled={!detailUrl} onClick={() => detailUrl && navigate(detailUrl)}>
              <Image src={Rack} alt="rack" height="30px" />
            </Button>
          </HStack>
        </Heading>
      </CardHeader>

      <Stack spacing={1}>
        <RackPhysicalDimension {...dimensions.power} />
        <RackPhysicalDimension {...dimensions.compute} />
        <RackPhysicalDimension {...dimensions.storage} />
        <RackPhysicalDimension {...dimensions.dataLink} />
      </Stack>
    </Card>
  )
}
