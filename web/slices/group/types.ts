import { type Account } from "~/slices/account/types";
import { type List } from "~/slices/list/types";

export interface Group {
  id: string;
  created: string;
  updated: string;
  account: string;
  expand?: {
    account?: Account;
    "lists(group)"?: List[];
  };
  name: string;
  description?: string;
  link?: string;
  // TODO: Provide enum typing???
  voicing: string;
  rating: number;
  starredAt?: string;
}
