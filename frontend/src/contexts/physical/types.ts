import { GeneralStatus } from "../shared/types"

export interface PhysicalDimensionSummary {
  name: string
  status: GeneralStatus
}

export interface PhysicalDimensionSummaries {
  power: PhysicalDimensionSummary
  compute: PhysicalDimensionSummary
  storage: PhysicalDimensionSummary
  dataLink: PhysicalDimensionSummary
}

export interface RackSummary {
  name: string
  id: string
  detailUrl?: string
  dimensions: PhysicalDimensionSummaries
}

export interface SiteSummary {
  id: string
  name: string
  racks: RackSummary[]
}

export interface RackDetails {
  id: string
  name: string
  dimensions: PhysicalDimensionSummaries
}

export interface PhysicalLayerSummary {
  sites: SiteSummary[]
  racks: Record<string, RackDetails>
}
