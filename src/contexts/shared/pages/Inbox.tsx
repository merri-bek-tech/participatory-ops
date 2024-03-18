import { Box, Card, CardBody, CardHeader, Grid, GridItem, Heading, SimpleGrid, Table, TableCaption, TableContainer, Tbody, Td, Text, Tfoot, Th, Thead, Tr } from "@chakra-ui/react"

function UnknownComponent() {
  return (
    <Card>
      <CardBody>
        <Heading as="h3" fontSize="lg">
          Unknown Component
        </Heading>
        <Box mt={4}>
          <TableContainer>
            <Table __css={{ "table-layout": "fixed", width: "full" }} variant="simple">
              <Tbody>
                <Tr>
                  <Th pl={0} w="50px">
                    UUID
                  </Th>
                  <Td p={0}>
                    <Text overflow="hidden">{"8431928b-a906-40de-bae2-ab30dfe5e2e3".slice(0, 20) + "..."}</Text>
                  </Td>
                </Tr>
              </Tbody>
            </Table>
          </TableContainer>
        </Box>
      </CardBody>
    </Card>
  )
}

export default function Inbox() {
  return (
    <Box>
      <Heading>Inbox</Heading>
      <Grid templateColumns={["repeat(1, 1fr)", "repeat(1, 1fr)", "repeat(2, 1fr)"]} gap={6}>
        {[1, 2, 3, 4, 5, 6, 7, 8, 9].map((_, i) => (
          <GridItem>
            <UnknownComponent />
          </GridItem>
        ))}
      </Grid>
    </Box>
  )
}
