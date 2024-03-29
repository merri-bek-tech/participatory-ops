import { Card, CardBody, CardHeader, Heading, Table, TableContainer, Tbody, Td, Text, Th, Tr, useColorModeValue } from "@chakra-ui/react"
import { UnknownComponent } from "../types"

export default function UnknownComponentCard({ status, uuid }: UnknownComponent) {
  const colors = {
    online: "green." + useColorModeValue("300", "600"),
    offline: "gray." + useColorModeValue("200", "500"),
  }

  return (
    <Card>
      <CardHeader bgColor={colors[status]} borderTopRadius={"inherit"}>
        <Heading as="h3" fontSize="lg">
          Unknown Component
        </Heading>
        <Text fontSize="sm">{status}</Text>
      </CardHeader>
      <CardBody>
        <TableContainer>
          <Table __css={{ tableLayout: "fixed", width: "full" }} variant="simple">
            <Tbody>
              <Tr>
                <Th p={0} w="50px">
                  UUID
                </Th>
                <Td p={0}>
                  <Text overflow="hidden">{uuid.slice(0, 20) + "..."}</Text>
                </Td>
              </Tr>
            </Tbody>
          </Table>
        </TableContainer>
      </CardBody>
    </Card>
  )
}
