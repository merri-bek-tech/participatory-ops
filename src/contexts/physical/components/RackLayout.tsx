import { Box, Stack, Image, Center } from "@chakra-ui/react"
import Switch1Ru from "/rack-components/switch-1ru.png"
import Switch2Ru from "/rack-components/switch-2ru.png"
import Server1Ru from "/rack-components/server-1ru.png"
import Server2Ru from "/rack-components/server-2ru.png"
import Server3Ru from "/rack-components/server-3ru.png"

const ruHeightOriginal = 55
const widthOriginal = 585

const width = 400
const ruHeight = (ruHeightOriginal / widthOriginal) * width

type ComponentImageNames = "switch" | "server"

const images: Record<ComponentImageNames, Record<number, string>> = {
  switch: {
    1: Switch1Ru,
    2: Switch2Ru,
  },
  server: {
    1: Server1Ru,
    2: Server2Ru,
    3: Server3Ru,
  },
}

function RackComponentBox({ children, ru, imageName }: { children: React.ReactNode; ru: number; imageName: ComponentImageNames }) {
  const image = images[imageName] && images[imageName][ru]

  return (
    <Box height={ruHeight * ru} pos="relative">
      {image && <Image src={image} pos="absolute" top={0} left={0} width="100%" height="100%" opacity={0.7} zIndex={0} alt="" />}

      <Box zIndex={1} pos="absolute" top={0} left={0} width="100%" height="100%">
        {children}
      </Box>
    </Box>
  )
}

function BackgroundedStrip({ children }: { children: React.ReactNode }) {
  return (
    <Stack direction="column" justifyContent="center" height="100%">
      <Stack direction="row" justifyContent="center" backgroundColor="whiteAlpha.700" color="black" width="100%" padding={1}>
        <Box fontSize="small" lineHeight="100%">
          {children}
        </Box>
      </Stack>
    </Stack>
  )
}

function NetworkSwitch({ name, ru }: { name: string; ru: number }) {
  return (
    <RackComponentBox ru={ru} imageName="switch">
      <BackgroundedStrip>{name}</BackgroundedStrip>
    </RackComponentBox>
  )
}

function Server({ name, ru }: { name: string; ru: number }) {
  return (
    <RackComponentBox ru={ru} imageName="server">
      <BackgroundedStrip>{name}</BackgroundedStrip>
    </RackComponentBox>
  )
}

export default function PhysicalLayout() {
  return (
    <Stack maxW={width} w="100%" gap={1} backgroundColor="gray.800" padding={1}>
      <NetworkSwitch name="Switch 1" ru={1} />
      <NetworkSwitch name="Switch 2" ru={2} />
      <Server name="Compute 1" ru={2} />
      <Server name="Compute 2" ru={1} />
      <Server name="Compute 2" ru={1} />
      <Server name="Compute 2" ru={3} />
    </Stack>
  )
}
