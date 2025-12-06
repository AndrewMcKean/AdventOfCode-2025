const lines = ["3-5", "10-14", "12-18", "16-20"];

let validRanges = [];
let numValidIDs = 0;

const generateRanges = () => {
  let ranges = [];
  for (const line of lines) {
    const [start, end] = line.split("-");
    ranges.push({ start: Number(start), end: Number(end) });
  }

  ranges.sort((a, b) => a.start - b.start);
  return ranges;
};

const ranges = generateRanges();
let current = ranges[0];

// I cheated here by looking up how to merge arrays as I couldnt get it right for the larger input
for (let i = 1; i < ranges.length; i++) {
  const nextRange = ranges[i];
  if (nextRange.start <= current.end + 1) {
    current.end = Math.max(current.end, nextRange.end);
  } else {
    validRanges.push(current);
    current = nextRange;
  }
}

validRanges.push(current);
for (const r of validRanges) {
  numValidIDs += r.end - r.start + 1;
}

console.log(numValidIDs);
