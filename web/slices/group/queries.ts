import { useQuery } from "@tanstack/vue-query";

import { pocketbase } from "~/api";
import { type PaginationInput } from "~/api/types";
import { getPaginationInput, getPaginationInputArgs } from "~/api/utils";

import { type Group } from "./types";

export const groupKeys = {
  all: ["groups"],
  allPaginated: (pagination?: PaginationInput) => [...groupKeys.all, pagination],
} as const;

export const useGroupsQuery = ({ pagination }: { pagination?: PaginationInput }) => {
  return useQuery({
    keepPreviousData: true,
    queryKey: groupKeys.allPaginated(pagination),
    queryFn: async () => {
      return pocketbase.collection("groups").getList<Group>(...getPaginationInputArgs(pagination), {
        ...getPaginationInput(pagination),
        sort: "name",
        expand: "lists(group)",
      });
    },
  });
};
