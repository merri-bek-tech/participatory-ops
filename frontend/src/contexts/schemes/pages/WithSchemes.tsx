import { useEffect, useState } from "react"
import { Scheme } from "../../shared/types"
import { FullPageLoading } from "../../shared"
import { Api } from "../../shared"

type ChildrenWithSchemes = (schemes: Scheme[]) => React.ReactNode

export default function WithSchemes({ children }: { children: ChildrenWithSchemes }) {
  const api = new Api()
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

  return children(schemes)
}
