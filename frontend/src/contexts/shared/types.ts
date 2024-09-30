export interface Scheme {
  id: string
  name: string
  hostnames: string[]
}

export type GeneralStatus = "planned" | "active" | "warning" | "error"
