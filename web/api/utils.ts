import { type Pagination } from "./types";

export const DEFAULT_PAGE_SIZE = 10;

/** Get pagination from potentially empty/invalid pagination values */
export const getPagination = (pagination: Partial<Pagination> | undefined): Pagination => {
  return {
    page: Math.max(pagination?.page ?? 1, 1),
    perPage: Math.max(pagination?.perPage ?? DEFAULT_PAGE_SIZE, 1),
  };
};

/** Get pocketbase pagination `getList` arguments  */
export const getPaginationArgs = (
  pagination?: Partial<Pagination>,
): [page: number, perPage: number] => {
  const safePagination = getPagination(pagination);
  return [safePagination.page, safePagination.perPage];
};
