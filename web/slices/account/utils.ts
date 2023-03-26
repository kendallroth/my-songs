import { type Admin, type Record } from "pocketbase";

import { type Account } from "./types";

export const mapRecordToAccount = (record: Record | Admin): Account => ({
  created: record.created,
  email: record.email,
  id: record.id,
  name: record.name ?? null,
  username: record.username,
  verified: record.verified ?? false,
});
