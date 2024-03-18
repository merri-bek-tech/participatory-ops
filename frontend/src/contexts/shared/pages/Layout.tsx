import { Outlet } from "react-router-dom";
import { Stack, Container, Box } from "@chakra-ui/react";
import Navbar from "../components/Navbar";

export default function Layout() {
  return (
    <Stack h="100vh" w="100vw" gap={0}>
      <Navbar />
      <Box overflow="scroll">
        <Container padding={4} maxWidth="1000px">
          <Outlet />
        </Container>
      </Box>
    </Stack>
  );
}
