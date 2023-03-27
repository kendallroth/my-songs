import { useQuery } from "@tanstack/vue-query";

import { pocketbase } from "~/api";
import { type PaginationInput } from "~/api/types";
import { getPaginationInput, getPaginationInputArgs } from "~/api/utils";

import { type List } from "./types";

export const listKeys = {
  all: (pagination?: PaginationInput) => ["lists", pagination],
} as const;

export const useListsQuery = ({ pagination }: { pagination?: PaginationInput }) => {
  return useQuery({
    keepPreviousData: true,
    queryKey: listKeys.all(pagination),
    queryFn: async () => {
      return pocketbase.collection("lists").getList<List>(...getPaginationInputArgs(pagination), {
        ...getPaginationInput(pagination),
        sort: "-date",
        expand: "songs",
      });
    },
  });
};
