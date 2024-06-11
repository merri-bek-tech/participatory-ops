import { Card, CardHeader, Heading, Stack, HStack, Button } from "@chakra-ui/react"

export default function AppCard() {
  return (
    <Card>
      <CardHeader p={1}>
        <Heading size="md" pl={2}>
          <HStack justify="space-between">
            <span>Wikipedia</span>
            <Button colorScheme="cyan">xx</Button>
          </HStack>
        </Heading>
      </CardHeader>

      <Stack spacing={1}></Stack>
    </Card>
  )
}
