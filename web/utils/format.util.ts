import dayjs from "dayjs";

export const formatDate = (date: dayjs.Dayjs | string, format = "YYYY-MM-DD") => {
  return dayjs(date).format(format);
};
