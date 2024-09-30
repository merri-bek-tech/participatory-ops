import { Image } from "@chakra-ui/react"
import RaspberryPi from "../../../assets/node-components/raspberry-pi-256.png"
import Laptop from "../../../assets/node-components/laptop-256.png"

function getImage(sysVendor: string) {
  switch (sysVendor) {
    case "Raspberry Pi":
      return RaspberryPi
    default:
      return Laptop
  }
}

export default function ComponentIcon({ sysVendor }: { sysVendor: string }) {
  const image = getImage(sysVendor)
  return <Image src={image} width="80px" height="80px" />
}
