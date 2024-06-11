import { GeneralStatus } from "../shared/types"

export interface AppDimensionSummary {
  name: string
  status: GeneralStatus
}

export interface AppSummary {
  id: string
  name: string
}
