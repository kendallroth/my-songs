import { useQuery } from "@tanstack/vue-query";

import { pocketbase } from "~/api";
import { type PaginationInput } from "~/api/types";
import { getPaginationInput, getPaginationInputArgs } from "~/api/utils";
import { type Song } from "~/slices/song/types";

export const songKeys = {
  all: (pagination?: PaginationInput) => ["songs", pagination],
} as const;

export const useSongsQuery = ({ pagination }: { pagination?: PaginationInput }) => {
  return useQuery({
    keepPreviousData: true,
    queryKey: songKeys.all(pagination),
    queryFn: async () => {
      return pocketbase.collection("songs").getList<Song>(...getPaginationInputArgs(pagination), {
        ...getPaginationInput(pagination),
        sort: "title",
      });
    },
  });
};
