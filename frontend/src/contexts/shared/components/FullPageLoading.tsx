import { Center, Spinner } from "@chakra-ui/react"

export default function FullPageLoading() {
  return (
    <Center h="100vh" w="100vw">
      <Spinner size="lg" />
    </Center>
  )
}
