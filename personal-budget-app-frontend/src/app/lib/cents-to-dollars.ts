
export default function convertToDollars(cents: number): string {
  return `$${(cents / 100).toFixed(2)}`;
}