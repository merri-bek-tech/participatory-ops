import { Heading, Text, HStack } from "@chakra-ui/react"
import { AppDimensionSummary } from "../types"
import { StatusBox } from "../../shared"

export default function AppDimension({ name, status }: AppDimensionSummary) {
  return (
    <StatusBox status={status}>
      <HStack justify="space-between">
        <Heading size="xs" textTransform="uppercase">
          {name}
        </Heading>
        <Text fontSize="sm">{status}</Text>
      </HStack>
    </StatusBox>
  )
}
