import { type ListResult } from "pocketbase";

/**
 * Convert an input value to an array (or empty array if invalid)
 *
 * Useful when input data might be a non-array JSON value.
 */
export const getArray = <T>(input: T): T[] => {
  if (!input) return [];
  if (!Array.isArray(input)) return [];

  return input;
};

type ResultTypeKeys<T> = keyof T | (keyof T)[];

/**
 * Convert a key(s) on an entity to an array (or empty array if invalid)
 *
 * Useful when input data might contain a non-array JSON value.
 */
export const getResultArray = <T>(input: T, keys: ResultTypeKeys<T>): T => {
  if (!input) return input;
  const keyList = Array.isArray(keys) ? keys : [keys];

  const newResult = { ...input };
  keyList.forEach((key) => {
    // TODO: Remove once figuring out how to limit passed keys to keys representing values...
    // @ts-ignore
    newResult[key] = getArray(input[key]);
  });
  return newResult;
};

/**
 * Convert a key(s) on a list of entities to an array (or empty array if invalid)
 *
 * Useful when input data might contain a non-array JSON value.
 */
export const getListResultArray = <T>(
  input: ListResult<T>,
  keys: ResultTypeKeys<T>,
): ListResult<T> => {
  if (!input) return input;

  return {
    ...input,
    items: input.items.map((item) => getResultArray(item, keys)),
  };
};
