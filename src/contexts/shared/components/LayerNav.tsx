import { Box, Text, Heading, VStack, useColorModeValue } from "@chakra-ui/react";
import Link from "./Link";

type LayerID = "application" | "capability" | "virtual" | "physical";

type LayerNavProps = {
  activeLayer: LayerID;
};

interface LayerLinkDetails {
  id: LayerID;
  name: string;
  enabled: boolean;
}

function LayerLink({ id, name, enabled, active }: LayerLinkDetails & { active: boolean }) {
  if (active) {
    return (
      <Heading borderBottom="2px solid black" lineHeight="100%">
        {name}
      </Heading>
    );
  } else {
    if (enabled) {
      return (
        <Link href={`/${id}`} borderBottomWidth={2} borderBottomColor={useColorModeValue("gray.900", "gray.100")}>
          {name}
        </Link>
      );
    } else {
      return (
        <Text color={"gray.600"} borderBottomWidth={2} borderBottomColor="gray.600">
          {name}
        </Text>
      );
    }
  }
}

export default function LayerNav({ activeLayer }: LayerNavProps) {
  const layers: LayerLinkDetails[] = [
    { id: "application", name: "Application Layer", enabled: false },
    { id: "capability", name: "Capability Layer", enabled: false },
    { id: "virtual", name: "Virtual Layer", enabled: true },
    { id: "physical", name: "Physical Layer", enabled: true },
  ];

  return (
    <VStack align="flex-start" mb={4}>
      {layers.map((details) => {
        return <LayerLink {...details} active={details.id == activeLayer} key={details.id} />;
      })}
    </VStack>
  );
}
