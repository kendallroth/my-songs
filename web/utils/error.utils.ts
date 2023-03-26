import { type PocketbaseError } from "~/api/types";

// Default error that will be used when no error message is found
const DEFAULT_ERROR_MESSAGE = "An unknown error occurred";

/**
 * Determine whether an error includes a specific error status code number
 *
 * @param   error           - Error object/string
 * @param   targetErrorCode - Target error status code
 */
export const hasStatusCode = (error: unknown, targetErrorCode: number): boolean => {
  if (!isPocketbaseError(error)) return false;

  return error.status === targetErrorCode;
};

/**
 * Get an error message from an error
 *
 * @param   error          - Error object/string
 * @param   defaultMessage - Default error message
 */
export const getError = (error: unknown, defaultMessage?: string): string => {
  if (!error) return "";

  const errorMessage = getErrorMessage(error);
  if (!errorMessage) {
    return defaultMessage ?? DEFAULT_ERROR_MESSAGE;
  }

  return errorMessage;
};

/**
 * Get an error message from an error
 *
 * @param   error - Error object/string
 */
const getErrorMessage = (error: unknown): string | null => {
  if (!error) return null;

  // A bare error code may be provided instead of an error object
  if (typeof error === "string") return error;

  if (isPocketbaseError(error)) {
    return getErrorMessage(error.response);
  }

  // Errors are often provided as an object, but the message 'key' may vary
  if (typeof error === "object") {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const errorObject = error as any;
    const message = errorObject.message || errorObject.error;

    return getErrorMessage(message);
  }

  return null;
};

/** Whether an error is a Pocketbase error */
export const isPocketbaseError = (error: unknown): error is PocketbaseError => {
  if (typeof error !== "object" || error === null) return false;

  const keys = ["originalError", "response", "status", "url"];
  const errorKeys = Object.keys(error);
  return keys.every((key) => errorKeys.includes(key));
};
