import { type Account } from "~/slices/account/types";

export interface Song {
  id: string;
  created: string;
  updated: string;
  account: string;
  accountObj?: Account;
  title: string;
  description?: string;
  // TODO: Provide enum typing???
  voicing: string;
  tags?: string[];
  rating: number;
  composer: string;
  arranged: boolean;
  starredAt?: string;
}
