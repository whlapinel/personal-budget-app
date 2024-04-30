import { TimeInfo } from '@/app/lib/data/definitions';

export default function getTimeInfo(month: number, year: number): (TimeInfo) {
  const today = new Date();
  const viewedMonth = month;
  const viewedYear = year;
  const viewedDate = new Date(year, month -1 ); // JS months are 0-indexed
  const nextMonth = month === 12 ? 1 : month + 1;
  const prevMonth = month === 1 ? 12 : month - 1;
  const nextYear = month === 12 ? year + 1 : year;
  const prevYear = month === 1 ? year - 1 : year;
  const monthString = viewedDate.toLocaleString('default', { month: 'long' });
  return { today, viewedMonth, viewedYear, viewedDate, nextMonth, prevMonth, nextYear, prevYear, monthString}
}