import { type Account } from "~/slices/account/types";
import { type Group } from "~/slices/group/types";
import { type Song } from "~/slices/song/types";

export interface List {
  id: string;
  created: string;
  updated: string;
  account: string;
  expand?: {
    account?: Account;
    group?: Group;
    songs?: Song[];
  };
  group?: string;
  title: string;
  description?: string;
  date: string;
  director?: string;
  tags?: string[];
  rating: number;
  starredAt?: string;
  songs: string[];
}
