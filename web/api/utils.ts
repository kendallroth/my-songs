import { type PaginationInput, type PaginationResponse } from "./types";

export const DEFAULT_PAGE_SIZE = 10;

/** Get pagination from potentially empty/invalid pagination values */
export const getPaginationInput = (
  pagination: Partial<PaginationInput> | undefined,
): PaginationInput => {
  return {
    page: Math.max(pagination?.page ?? 1, 1),
    perPage: Math.max(pagination?.perPage ?? DEFAULT_PAGE_SIZE, 1),
  };
};

/** Get Pocketbase pagination `getList` arguments  */
export const getPaginationInputArgs = (
  pagination?: Partial<PaginationInput>,
): [page: number, perPage: number] => {
  const safePagination = getPaginationInput(pagination);
  return [safePagination.page, safePagination.perPage];
};

/** Get pagination from Pocketbase response  */
export const getPaginationResponse = (
  input: Partial<PaginationResponse> = {},
): PaginationResponse => ({
  page: input.page ?? 1,
  perPage: input.perPage ?? 10,
  totalItems: input.totalItems ?? 0,
  totalPages: input.totalPages ?? 0,
});
