import { useQuery } from "@tanstack/vue-query";

import { pocketbase } from "~/api";

export const dashboardKeys = {
  stats: ["stats"],
} as const;

export const useStatsQuery = () => {
  return useQuery({
    keepPreviousData: true,
    queryKey: dashboardKeys.stats,
    queryFn: async () => {
      const stats = await pocketbase.send("/api/stats", {});

      return {
        groups: stats.groups ?? 0,
        lists: stats.lists ?? 0,
        songs: stats.songs ?? 0,
      };
    },
  });
};
