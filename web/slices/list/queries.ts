import { useMutation, useQuery, useQueryClient } from "@tanstack/vue-query";
import dayjs from "dayjs";

import { pocketbase } from "~/api";
import { type PaginationInput } from "~/api/types";
import { getPaginationInput, getPaginationInputArgs } from "~/api/utils";

import { type List } from "./types";

export const listKeys = {
  all: ["lists"],
  allPaginated: (pagination?: PaginationInput) => [...listKeys.all, pagination],
  single: ["list"],
  byId: (listId: string) => [...listKeys.single, listId],
} as const;

export const useGetListByIdQuery = ({ listId }: { listId: string }) => {
  return useQuery({
    keepPreviousData: true,
    queryKey: listKeys.byId(listId),
    queryFn: async () => {
      return pocketbase.collection("lists").getOne<List>(listId, {
        expand: "group,songs",
      });
    },
    select: (result) => getResultArray(result, "tags"),
  });
};

export const useGetListsQuery = ({ pagination }: { pagination?: PaginationInput }) => {
  return useQuery({
    keepPreviousData: true,
    queryKey: listKeys.allPaginated(pagination),
    queryFn: async () => {
      return pocketbase.collection("lists").getList<List>(...getPaginationInputArgs(pagination), {
        ...getPaginationInput(pagination),
        sort: "-date",
        expand: "group,songs",
      });
    },
    select: (result) => getListResultArray(result, "tags"),
  });
};

export const useStarListMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async ({ listId, starred }: { listId: string; starred: boolean }) => {
      return pocketbase.collection("lists").update<List>(listId, {
        starredAt: starred ? dayjs().toISOString() : null,
      });
    },
    onSuccess(data, variables) {
      queryClient.invalidateQueries({ queryKey: listKeys.all });
      // Update selected list to avoid an extra/unnecessary query
      queryClient.setQueryData(listKeys.byId(variables.listId), (oldData?: List) => {
        if (!oldData) return undefined;
        return { ...oldData, starredAt: data.starredAt };
      });
    },
  });
};
