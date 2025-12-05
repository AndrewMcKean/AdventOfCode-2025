const test = ["3-5", "10-14", "16-20", "12-18"];
const listOfRanges: { start: number; end: number }[] = [];

let totalUniqueIDs = 0;

for (const range of test) {
  const [start, end] = range.split("-");
  const numStart = Number(start);
  const numEnd = Number(end);
  listOfRanges.push({ start: numStart, end: numEnd });
}

// remove subsets
for (const rangeEntry of listOfRanges) {
  for (const [index, sussEntry] of listOfRanges.entries()) {
    if (
      rangeEntry.start <= sussEntry.start &&
      rangeEntry.end >= sussEntry.end
    ) {
      listOfRanges.splice(index, 1);
    }
  }
}

// amend ovelaps
for (const [index, rangeEntry] of listOfRanges.entries()) {
  for (const sussEntry of listOfRanges) {
    if (
      rangeEntry.start <= sussEntry.start &&
      rangeEntry.end <= sussEntry.end
    ) {
      listOfRanges[index] = { start: rangeEntry.start, end: sussEntry.end };
    }

    if (rangeEntry.end > sussEntry.end && rangeEntry.start >= sussEntry.start) {
      listOfRanges[index] = { start: sussEntry.end, end: rangeEntry.end };
    }
  }
}

for (const val of listOfRanges) {
  totalUniqueIDs += val.end + 1 - val.start;
}

console.log(listOfRanges);
