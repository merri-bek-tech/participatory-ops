import { Card, CardHeader, Heading, Stack, HStack, Button } from "@chakra-ui/react"
import AppDimension from "./AppDimension"
import { AppSummary } from "../types"

export default function AppCard({ id, name }: AppSummary) {
  return (
    <Card>
      <CardHeader p={1}>
        <Heading size="md" pl={2}>
          <HStack justify="space-between">
            <span>{name}</span>
            <Button colorScheme="cyan">xx</Button>
          </HStack>
        </Heading>
      </CardHeader>

      <Stack spacing={1}>
        <AppDimension name="availability" status="planned" />
      </Stack>
    </Card>
  )
}
