import { Card, CardBody, CardHeader, Heading, Table, TableContainer, Tbody, Td, Text, Th, Tr, useColorModeValue } from "@chakra-ui/react"
import { ComponentStatus } from "../types"

export default function ComponentStatusCard({ status, uuid, details }: ComponentStatus) {
  const colors = {
    online: "green." + useColorModeValue("300", "600"),
    offline: "gray." + useColorModeValue("200", "500"),
  }

  const hostName = details?.hostName || "Unknown Component"

  return (
    <Card>
      <CardHeader bgColor={colors[status]} borderTopRadius={"inherit"}>
        <Heading as="h3" fontSize="lg">
          {hostName}
        </Heading>
        <Text fontSize="sm">{status}</Text>
      </CardHeader>
      <CardBody>
        {details && (
          <TableContainer>
            <Table variant="simple">
              <Tbody>
                {Object.entries(details).map(([key, value]) => (
                  <Tr key={key}>
                    <Th p={0} w="1px" whiteSpace="nowrap" pr={2}>
                      {key}
                    </Th>
                    <Td p={0} maxWidth="100px">
                      <Text overflow="hidden" whiteSpace="nowrap" textOverflow="ellipsis">
                        {value}
                      </Text>
                    </Td>
                  </Tr>
                ))}
              </Tbody>
            </Table>
          </TableContainer>
        )}
      </CardBody>
    </Card>
  )
}
