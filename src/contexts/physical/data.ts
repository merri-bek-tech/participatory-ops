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
            power: { name: "Power", status: "active" },
            compute: { name: "Compute", status: "active" },
            storage: { name: "Storage", status: "active" },
            dataLink: { name: "Data link", status: "active" },
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
            power: { name: "Power", status: "active" },
            compute: { name: "Compute", status: "error" },
            storage: { name: "Storage", status: "warning" },
            dataLink: { name: "Data link", status: "active" },
          },
        },
        {
          name: "Murmbal",
          id: "murmbal",
          dimensions: {
            power: { name: "Power", status: "planned" },
            compute: { name: "Compute", status: "planned" },
            storage: { name: "Storage", status: "planned" },
            dataLink: { name: "Data link", status: "planned" },
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
            power: { name: "Power", status: "planned" },
            compute: { name: "Compute", status: "planned" },
            storage: { name: "Storage", status: "planned" },
            dataLink: { name: "Data link", status: "planned" },
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
            power: { name: "Power", status: "planned" },
            compute: { name: "Compute", status: "planned" },
            storage: { name: "Storage", status: "planned" },
            dataLink: { name: "Data link", status: "planned" },
          },
        },
        {
          name: "Wangnarra",
          id: "wangnarra",
          dimensions: {
            power: { name: "Power", status: "planned" },
            compute: { name: "Compute", status: "planned" },
            storage: { name: "Storage", status: "planned" },
            dataLink: { name: "Data link", status: "planned" },
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
        power: { name: "Power", status: "active" },
        compute: { name: "Compute", status: "active" },
        storage: { name: "Storage", status: "active" },
        dataLink: { name: "Data link", status: "active" },
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
