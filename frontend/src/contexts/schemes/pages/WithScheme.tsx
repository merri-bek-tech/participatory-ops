import { useEffect, useState } from "react"
import { Scheme } from "../../shared/types"
import { FullPageLoading } from "../../shared"
import { Api } from "../../shared"

const api = new Api()

export default function WithScheme({ children }: { children: React.ReactNode }) {
  const [schemes, setSchemes] = useState<Scheme[]>([])

  useEffect(() => {
    api.schemesIndex().then((data: Scheme[]) => {
      setSchemes(data)
      console.log("got schemes: ", data)
    })
  }, [])

  if (schemes.length === 0) {
    return <FullPageLoading />
  }

  return children
}
