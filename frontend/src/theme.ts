import { theme as chakraTheme, extendTheme } from "@chakra-ui/react";
import type { StyleFunctionProps } from "@chakra-ui/styled-system";
import { mode } from "@chakra-ui/theme-tools";

const colors = {
  ...chakraTheme.colors,
  gray: {
    50: "#F7FAFC",
    100: "#dcd8d9",
    200: "#bfbfbf",
    300: "#a6a6a6",
    400: "#8c8c8c",
    500: "#737373",
    600: "#595959",
    700: "#404040",
    800: "#282626",
    850: "#1b1919",
    900: "#150a0d",
  },
  blue: {
    ...chakraTheme.colors.blue,
    250: "#8bc7f3",
  },
};

const createSingleGradient = (color: string) => `linear-gradient(to right, ${color} 1px, transparent 1px),linear-gradient(to bottom, ${color} 1px, transparent 1px)`;

const createGradient = (light: string, dark: string) => mode(createSingleGradient(light), createSingleGradient(dark));

export const theme = extendTheme({
  config: {
    initialColorMode: "light",
    useSystemColorMode: false,
  },
  colors: colors,
  styles: {
    global: (props: StyleFunctionProps) => ({
      body: {
        bg: mode("blue.300", "gray.800")(props),
        bgImage: createGradient(colors.blue[250], colors.gray[850])(props),
        bgSize: "40px 40px",
      },
    }),
  },
});
