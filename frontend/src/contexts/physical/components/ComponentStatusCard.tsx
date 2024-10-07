import {
  Accordion,
  AccordionButton,
  AccordionIcon,
  AccordionItem,
  AccordionPanel,
  Box,
  Card,
  CardBody,
  CardHeader,
  Heading,
  HStack,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Tr,
  useColorModeValue,
  VStack,
} from "@chakra-ui/react"
import { ComponentStatus, ComponentDetails } from "../types"
import ComponentIcon from "./ComponentIcon"

function DetailsTable({ details }: { details: ComponentDetails }) {
  return (
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
  )
}

export default function ComponentStatusCard({ status, uuid, details }: ComponentStatus) {
  const colors = {
    online: "green." + useColorModeValue("300", "600"),
    offline: "gray." + useColorModeValue("200", "500"),
  }

  const hostName = details?.hostName || "Unknown Component"

  return (
    <Accordion allowToggle>
      <AccordionItem border="none">
        <Card>
          <CardHeader bgColor={colors[status]} borderRadius={"inherit"}>
            <HStack>
              <Box flexGrow={1}>
                <HStack>
                  {details && <ComponentIcon sysVendor={details.sysVendor} />}
                  <VStack alignItems="flex-start" justifyContent="flex-start" gap={0} flexGrow="1">
                    <Heading as="h3" fontSize="lg">
                      {hostName}
                    </Heading>
                    <Text fontSize="sm">{status}</Text>
                  </VStack>
                </HStack>
              </Box>
              {details && (
                <Box>
                  <AccordionButton>
                    <AccordionIcon />
                  </AccordionButton>
                </Box>
              )}
            </HStack>
          </CardHeader>

          {details && (
            <AccordionPanel>
              <CardBody>
                <DetailsTable details={details} />
              </CardBody>
            </AccordionPanel>
          )}
        </Card>
      </AccordionItem>
    </Accordion>
  )

  // return (
  //   <Card>
  //     <CardHeader bgColor={colors[status]} borderTopRadius={"inherit"}>
  //       <HStack>
  //         <Box flexGrow={1}>
  //           <HStack>
  //             {details && <ComponentIcon sysVendor={details.sysVendor} />}
  //             <VStack alignItems="flex-start" justifyContent="flex-start" gap={0} flexGrow="1">
  //               <Heading as="h3" fontSize="lg">
  //                 {hostName}
  //               </Heading>
  //               <Text fontSize="sm">{status}</Text>
  //             </VStack>
  //           </HStack>
  //         </Box>
  //         <Box></Box>
  //       </HStack>
  //     </CardHeader>

  //     <CardBody>
  //       {details && (
  //         <Accordion allowToggle>
  //           <AccordionItem border="none">
  //             <VStack alignItems="stretch">
  //               <Box>
  //                 <AccordionButton>
  //                   <AccordionIcon />
  //                 </AccordionButton>
  //               </Box>
  //               <Box>
  //                 <AccordionPanel>
  //                   <DetailsTable details={details} />
  //                 </AccordionPanel>
  //               </Box>
  //             </VStack>
  //           </AccordionItem>
  //         </Accordion>
  //       )}
  //     </CardBody>
  //   </Card>
  // )
}
