import PocketBase from "pocketbase";

// TODO: Figure out better way to access this (from runtime config?)
const apiUrl = import.meta.env.VITE_API_URL;
if (!apiUrl || typeof apiUrl !== "string") {
  throw Error("API_URL is invalid!");
}

const client = new PocketBase(apiUrl);

export default client;
