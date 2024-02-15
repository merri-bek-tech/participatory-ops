import { Box, BoxProps, useColorModeValue } from "@chakra-ui/react";
import { GeneralStatus } from "../types";

type StatusColor = "green" | "gray" | "orange" | "red";

function ColourBox({ dashed, children, colour }: { dashed: boolean; colour: StatusColor; children: React.ReactNode }) {
  let boxProps: BoxProps = {
    p: 2,
    borderWidth: 3,
    borderColor: "transparent",
  };

  if (dashed) {
    boxProps.borderColor = useColorModeValue("gray.500", "gray.600");
    boxProps.borderWidth = 3;
    boxProps.borderStyle = "dashed";
  }

  let bgHue = useColorModeValue("300", "600");
  if (colour === "gray") bgHue = useColorModeValue("200", "500");

  return (
    <Box {...boxProps} bgColor={`${colour}.${bgHue}`}>
      {children}
    </Box>
  );
}

function colourForStatus(status: GeneralStatus): StatusColor {
  const statusColorMap: Record<GeneralStatus, StatusColor> = {
    active: "green",
    planned: "gray",
    warning: "orange",
    error: "red",
  };
  return statusColorMap[status];
}

export default function StatusBox({ status, children }: { status: GeneralStatus; children: React.ReactNode }) {
  return (
    <ColourBox dashed={status == "planned"} colour={colourForStatus(status)}>
      {children}
    </ColourBox>
  );
}
