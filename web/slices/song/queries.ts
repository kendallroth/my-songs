import { useQuery } from "@tanstack/vue-query";

import { pocketbase } from "~/api";
import { type Pagination } from "~/api/types";
import { getPagination, getPaginationArgs } from "~/api/utils";
import { type Song } from "~/slices/song/types";

export const songKeys = {
  all: (pagination?: Pagination) => ["songs", pagination],
} as const;

export const useSongsQuery = ({ pagination }: { pagination?: Pagination }) => {
  return useQuery({
    keepPreviousData: true,
    queryKey: songKeys.all(pagination),
    queryFn: async () => {
      return pocketbase.collection("songs").getList<Song>(...getPaginationArgs(pagination), {
        ...getPagination(pagination),
        sort: "title",
        // expand: "account",
      });
    },
  });
};
