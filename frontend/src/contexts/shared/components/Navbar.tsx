import { Box, Flex, useColorModeValue, Stack } from "@chakra-ui/react"
import ColorModeSwitcher from "./ColorModeSwitcher"
import Link from "./Link"
import { SchemeContext } from "../../schemes"
import { useContext } from "react"

export default function Navbar() {
  const scheme = useContext(SchemeContext)

  return (
    <Box bg={useColorModeValue("gray.100", "gray.900")} px={4} borderBottomColor={useColorModeValue("gray.600", "black")} borderBottomWidth={2} borderBottomStyle="solid" boxShadow="md">
      <Flex alignItems={"center"} justifyContent={"space-between"}>
        <Stack direction="row" alignItems="center">
          <Link href="/">{scheme.name}</Link>
        </Stack>

        <Flex alignItems={"center"} gap={4}>
          <Link href="/">Layers</Link>
          <Link href="/inbox">Inbox</Link>
          <Stack direction={"row"} spacing={7}>
            <ColorModeSwitcher />
          </Stack>
        </Flex>
      </Flex>
    </Box>
  )
}
