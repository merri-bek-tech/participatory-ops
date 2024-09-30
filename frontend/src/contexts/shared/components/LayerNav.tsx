import { Box, Select } from "@chakra-ui/react"
import { useNavigate } from "react-router-dom"

type LayerID = "application" | "capability" | "virtual" | "physical"

type LayerNavProps = {
  activeLayer: LayerID
}

interface LayerLinkDetails {
  id: LayerID
  name: string
  enabled: boolean
}

export default function LayerNav({ activeLayer }: LayerNavProps) {
  const layers: LayerLinkDetails[] = [
    { id: "application", name: "Application Layer", enabled: true },
    { id: "capability", name: "Capability Layer", enabled: false },
    { id: "virtual", name: "Virtual Layer", enabled: false },
    { id: "physical", name: "Physical Layer", enabled: true },
  ]

  const navigate = useNavigate()
  const selectLayer = (layer: string) => {
    navigate(`/${layer}`)
  }

  return (
    <Box maxW={300} marginBottom="1.5em">
      <Select variant="filled" value={activeLayer} onChange={(event) => selectLayer(event.target.value)}>
        {layers.map((layer) => (
          <option key={layer.id} value={layer.id} disabled={!layer.enabled}>
            {layer.name}
          </option>
        ))}
      </Select>
    </Box>
  )
}
