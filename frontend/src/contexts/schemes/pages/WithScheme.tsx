import { useEffect, useState } from "react"
import { Scheme } from "../../shared/types"
import WithSchemes from "./WithSchemes"
import { SchemeContext } from "../context"
import { Center } from "@chakra-ui/react"

function WithSchemeFromSchemes({ schemes, children }: { schemes: Scheme[]; children: React.ReactNode }) {
  const [scheme, setScheme] = useState<Scheme | null>(null)

  useEffect(() => {
    if (schemes.length > 0) {
      setScheme(schemes[0])
    }
  }, [schemes])

  if (!scheme) {
    return <Center>Could not select a scheme.</Center>
  }
  return <SchemeContext.Provider value={scheme}>{children}</SchemeContext.Provider>
}

export default function WithScheme({ children }: { children: React.ReactNode }) {
  return <WithSchemes>{(schemes: Scheme[]) => <WithSchemeFromSchemes schemes={schemes}>{children}</WithSchemeFromSchemes>}</WithSchemes>
}
