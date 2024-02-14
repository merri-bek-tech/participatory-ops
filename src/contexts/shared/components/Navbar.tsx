import { Box, Flex, useColorModeValue, Stack } from "@chakra-ui/react";
import ColorModeSwitcher from "./ColorModeSwitcher";

export default function Navbar() {
  return (
    <Box bg={useColorModeValue("gray.100", "gray.900")} px={4} borderBottomColor={useColorModeValue("gray.600", "black")} borderBottomWidth={2} borderBottomStyle="solid" boxShadow="md">
      <Flex alignItems={"center"} justifyContent={"space-between"}>
        <Stack direction="row" alignItems="center">
          <Box>Merri-bek Tech</Box>
        </Stack>

        <Flex alignItems={"center"}>
          <Stack direction={"row"} spacing={7}>
            <ColorModeSwitcher />
          </Stack>
        </Flex>
      </Flex>
    </Box>
  );
}
