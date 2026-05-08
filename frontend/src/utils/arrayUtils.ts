export function mergeUniqueSortedStringArrays(original: string[], added: string[]): string[] {
  return [...new Set([...original, ...added])].sort((a, b) =>
    a.toLowerCase() > b.toLowerCase() ? 1 : -1,
  )
}
