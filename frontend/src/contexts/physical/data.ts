import { PhysicalLayerSummary } from "./types"

const layer: PhysicalLayerSummary = {
  sites: [
    {
      name: "Radish House",
      id: "radish-house",
      racks: [
        {
          name: "Murnong",
          id: "murnong",
          detailUrl: "/physical/rack/murnong",
          dimensions: {
            power: { name: "power", status: "active" },
            compute: { name: "compute", status: "active" },
            storage: { name: "storage", status: "active" },
            dataLink: { name: "dataLink", status: "active" },
          },
        },
      ],
    },
    {
      name: "Brunswick Library",
      id: "brunswick-library",
      racks: [
        {
          name: "Lilly Pilly",
          id: "lilly-pilly",
          dimensions: {
            power: { name: "power", status: "active" },
            compute: { name: "compute", status: "error" },
            storage: { name: "storage", status: "warning" },
            dataLink: { name: "dataLink", status: "active" },
          },
        },
        {
          name: "Murmbal",
          id: "murmbal",
          dimensions: {
            power: { name: "power", status: "planned" },
            compute: { name: "compute", status: "planned" },
            storage: { name: "storage", status: "planned" },
            dataLink: { name: "dataLink", status: "planned" },
          },
        },
      ],
    },
    {
      name: "CERES",
      id: "ceres",
      racks: [
        {
          name: "Kingfisher",
          id: "kingfisher",
          dimensions: {
            power: { name: "power", status: "planned" },
            compute: { name: "compute", status: "planned" },
            storage: { name: "storage", status: "planned" },
            dataLink: { name: "dataLink", status: "planned" },
          },
        },
      ],
    },
    {
      name: "Glenroy Community Hub",
      id: "glenroy-hub",
      racks: [
        {
          name: "Woolip",
          id: "woolip",
          dimensions: {
            power: { name: "power", status: "planned" },
            compute: { name: "compute", status: "planned" },
            storage: { name: "storage", status: "planned" },
            dataLink: { name: "dataLink", status: "planned" },
          },
        },
        {
          name: "Wangnarra",
          id: "wangnarra",
          dimensions: {
            power: { name: "power", status: "planned" },
            compute: { name: "compute", status: "planned" },
            storage: { name: "storage", status: "planned" },
            dataLink: { name: "dataLink", status: "planned" },
          },
        },
      ],
    },
  ],
  racks: {
    murnong: {
      id: "murnong",
      name: "Murnong",
      dimensions: {
        power: { name: "power", status: "active" },
        compute: { name: "compute", status: "active" },
        storage: { name: "storage", status: "active" },
        dataLink: { name: "dataLink", status: "active" },
      },
    },
  },
}

export async function siteListLoader({ params }: any) {
  return { sites: layer.sites }
}

export async function rackDetailsLoader({ params }: any) {
  return { rack: layer.racks[params.id] }
}
