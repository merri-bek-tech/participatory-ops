import { Heading, Box, Text } from "@chakra-ui/react";

export default function RackPhysicalDimension({ name }: { name: string }) {
  return (
    <Box bgColor="gray.300" p={2}>
      <Heading size="xs" textTransform="uppercase">
        {name}
      </Heading>
      <Text pt="2" fontSize="sm">
        not set up
      </Text>
    </Box>
  );
}
