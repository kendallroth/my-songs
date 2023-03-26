import { type Pagination } from "~/api/types";

import { DEFAULT_PAGE_SIZE, getPagination, getPaginationArgs } from "../utils";

describe("getPagination", () => {
  describe("parses pagination with defaults", () => {
    const cases: {
      expected?: Partial<Pagination>;
      label: string;
      value: Partial<Pagination> | undefined;
    }[] = [
      { label: "undefined", value: undefined },
      { label: "empty", value: {} },
      { expected: { page: 2 }, label: "page only", value: { page: 2 } },
      { label: "invalid page", value: { page: -1 } },
      { label: "pageSize only", value: { perPage: DEFAULT_PAGE_SIZE } },
      { expected: { perPage: 1 }, label: "invalid pageSize", value: { perPage: -1 } },
      { label: "regular", value: { page: 1, perPage: DEFAULT_PAGE_SIZE } },
    ];

    test.each(cases)("parses incomplete pagination correctly", (input) => {
      const output = getPagination(input.value);

      const expected: Pagination = {
        page: 1,
        perPage: DEFAULT_PAGE_SIZE,
        ...(input.expected ?? {}),
      };

      expect(output).toStrictEqual(expected);
    });
  });
});

describe("getPaginationArgs", () => {
  test("returns proper arguments", () => {
    const pagination: Pagination = { page: 2, perPage: 20 };
    const output = getPaginationArgs(pagination);

    const expected = [pagination.page, pagination.perPage];
    expect(output).toStrictEqual(expected);
  });
});
