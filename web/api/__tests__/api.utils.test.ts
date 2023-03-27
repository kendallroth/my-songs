import { type PaginationInput } from "~/api/types";

import { DEFAULT_PAGE_SIZE, getPaginationInput, getPaginationInputArgs } from "../utils";

describe("getPaginationInput", () => {
  describe("parses pagination with defaults", () => {
    const cases: {
      expected?: Partial<PaginationInput>;
      label: string;
      value: Partial<PaginationInput> | undefined;
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
      const output = getPaginationInput(input.value);

      const expected: PaginationInput = {
        page: 1,
        perPage: DEFAULT_PAGE_SIZE,
        ...(input.expected ?? {}),
      };

      expect(output).toStrictEqual(expected);
    });
  });
});

describe("getPaginationInputArgs", () => {
  test("returns proper arguments", () => {
    const pagination: PaginationInput = { page: 2, perPage: 20 };
    const output = getPaginationInputArgs(pagination);

    const expected = [pagination.page, pagination.perPage];
    expect(output).toStrictEqual(expected);
  });
});
