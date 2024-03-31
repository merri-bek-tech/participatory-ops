import { useState } from "react"
import { Scheme } from "../types"
import { FullPageLoading } from "../../shared"

export default function WithScheme({ children }: { children: React.ReactNode }) {
  const [schemes, setSchemes] = useState<Scheme[]>([])

  if (schemes.length === 0) {
    return <FullPageLoading />
  }

  return <div>with scheme</div>
}
