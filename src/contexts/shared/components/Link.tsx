import { Link as ReactRouterLink } from "react-router-dom";
import { Link as ChakraLink, LinkProps } from "@chakra-ui/react";

export default function Link(props: LinkProps) {
  return (
    <ChakraLink as={ReactRouterLink} {...props} to={props.href}>
      {props.children}
    </ChakraLink>
  );
}
