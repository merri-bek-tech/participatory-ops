import { Box, Stack, Image, Center, Text, useColorModeValue } from "@chakra-ui/react"
import Switch1Ru from "/rack-components/switch-1ru.png"
import Switch2Ru from "/rack-components/switch-2ru.png"
import Server1Ru from "/rack-components/server-1ru.png"
import Server2Ru from "/rack-components/server-2ru.png"
import Server3Ru from "/rack-components/server-3ru.png"
import Raid2Ru from "/rack-components/raid-2ru.png"
import Raid3Ru from "/rack-components/raid-3ru.png"
import Spacer1Ru from "/rack-components/spacer-1ru.png"

const ruHeightOriginal = 55
const widthOriginal = 585

const width = 400
const ruHeight = Math.floor((ruHeightOriginal / widthOriginal) * width)
const ruGap = 4

type ComponentImageNames = "switch" | "server" | "raid" | "spacer"

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
  raid: {
    1: Raid2Ru,
    2: Raid2Ru,
    3: Raid3Ru,
  },
  spacer: {
    1: Spacer1Ru,
  },
}

function RackComponentBox({ children, ru, imageName, faded }: { children?: React.ReactNode; ru: number; imageName: ComponentImageNames; faded?: boolean }) {
  const sizes = images[imageName]
  const image = sizes && (sizes[ru] || sizes[1])
  const gapHeight = ruGap * Math.max(ru - 1, 0)
  const height = ruHeight * ru + gapHeight

  return (
    <Box height={height} pos="relative" opacity={faded ? 0.3 : 1.0}>
      {image && <Image src={image} pos="absolute" top={0} left={0} width="100%" height="100%" opacity={1.0} zIndex={0} alt="" />}

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

function Raid({ name, ru }: { name: string; ru: number }) {
  return (
    <RackComponentBox ru={ru} imageName="raid">
      <BackgroundedStrip>{name}</BackgroundedStrip>
    </RackComponentBox>
  )
}

function Spacer({ ru }: { ru: number }) {
  if (ru == 1) return <RackComponentBox ru={ru} imageName="spacer" faded={true} />
}

function RuLabels({ ru }: { ru: number }) {
  return (
    <Stack direction="column" height="100%" gap={`${ruGap}px`} pr={1}>
      {Array.from({ length: ru }).map((_, i) => (
        <Box key={i} height={ruHeight} display="flex" flexDirection="column" justifyContent="center" alignItems="flex-end">
          <Text fontSize="small" fontWeight="bold" lineHeight="100%">
            {ru - i}
          </Text>
        </Box>
      ))}
    </Stack>
  )
}

export default function PhysicalLayout() {
  return (
    <Stack direction="row" maxW={width} w="100%" gap={1} backgroundColor={useColorModeValue("gray.800", "black")} padding={3} borderRadius={5}>
      <RuLabels ru={18} />
      <Stack direction="column" w="100%" gap={`${ruGap}px`}>
        <NetworkSwitch name="Switch 1" ru={1} />
        <NetworkSwitch name="Switch 2" ru={2} />
        <Server name="Compute 1" ru={2} />
        <Spacer ru={1} />
        <Server name="Compute 2" ru={1} />
        <Server name="Compute 2" ru={1} />
        <Spacer ru={1} />
        <Spacer ru={1} />
        <Server name="Compute 2" ru={3} />
        <Raid name="RAID 1" ru={3} />
        <Raid name="RAID 2" ru={2} />
      </Stack>
    </Stack>
  )
}
