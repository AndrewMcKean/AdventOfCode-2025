let total = 0;

const getHighestStringInRow = (row: string, startPos: number) => {
  let highestValue = 0;
  let highestValueIndex = 0;
  for (let i = startPos; i < row.length; i++) {
    if (Number(row[i]) > highestValue) {
      highestValue = Number(row[i]);
      highestValueIndex = i;
    }
  }

  return { value: highestValue.toString(), index: highestValueIndex };
};

for (const row of input) {
  let totalString = "";
  let currHighestIndex = -1;

  for (let timesToSearch = 12; timesToSearch > 0; timesToSearch--) {
    const subString = row.slice(0, row.length + 1 - timesToSearch);
    const { index, value } = getHighestStringInRow(
      subString,
      currHighestIndex + 1
    );
    totalString += value;
    currHighestIndex = index;
  }
  total += Number(totalString);
}

console.log(total);
