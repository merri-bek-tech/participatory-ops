import { Stack } from "@chakra-ui/react"

export default function BaseLayout({ children }: { children: React.ReactNode }) {
  return (
    <Stack h="100vh" w="100vw" gap={0}>
      {children}
    </Stack>
  )
}
