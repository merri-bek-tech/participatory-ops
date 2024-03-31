import { createContext } from "react"
import { Scheme } from "../shared/types"

export const SchemeContext = createContext<Scheme>({} as Scheme)
