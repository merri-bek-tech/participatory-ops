export type PhysicalDimensionStatus = "planned" | "active" | "warning" | "error";

export interface PhysicalDimensionSummary {
  name: string;
  status: PhysicalDimensionStatus;
}

export interface PhysicalDimensionSummaries {
  power: PhysicalDimensionSummary;
  compute: PhysicalDimensionSummary;
  storage: PhysicalDimensionSummary;
  dataLink: PhysicalDimensionSummary;
}

export interface RackSummary {
  name: string;
  id: string;
  dimensions: PhysicalDimensionSummaries;
}

export interface SiteSummary {
  name: string;
  id: string;
  racks: RackSummary[];
}

export interface PhysicalLayerSummary {
  sites: SiteSummary[];
}
