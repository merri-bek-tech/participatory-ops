import { useEffect, useState } from "react"
import { Scheme } from "../../shared/types"
import WithSchemes from "./WithSchemes"
import { SchemeContext } from "../context"
import { Center } from "@chakra-ui/react"

function WithSchemeFromSchemes({ schemes, hostname, children }: { schemes: Scheme[]; hostname: string; children: React.ReactNode }) {
  const [scheme, setScheme] = useState<Scheme | null>(null)

  useEffect(() => {
    setScheme(schemes.find((scheme) => scheme.hostnames.includes(hostname)) || null)
  }, [schemes, hostname])

  if (!scheme) {
    return <Center>Could not select a scheme.</Center>
  }
  return <SchemeContext.Provider value={scheme}>{children}</SchemeContext.Provider>
}

export default function WithScheme({ children }: { children: React.ReactNode }) {
  return (
    <WithSchemes>
      {(schemes: Scheme[]) => (
        <WithSchemeFromSchemes schemes={schemes} hostname={window.location.hostname}>
          {children}
        </WithSchemeFromSchemes>
      )}
    </WithSchemes>
  )
}
