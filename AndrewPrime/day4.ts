const input = ["@"];

const canRemovePaper = (xPos: number, yPos: number, rowLength: number) => {
  let numPaperFound = 0;

  if (input[yPos][xPos] !== "@") {
    return false;
  }
  // top left
  if (xPos - 1 >= 0 && yPos - 1 >= 0) {
    numPaperFound += +(input[yPos - 1][xPos - 1] === "@");
  }
  // top mid
  if (yPos - 1 >= 0) {
    numPaperFound += +(input[yPos - 1][xPos] === "@");
  }
  // top right
  if (xPos + 1 < rowLength && yPos - 1 >= 0) {
    numPaperFound += +(input[yPos - 1][xPos + 1] === "@");
  }
  // left
  if (xPos - 1 >= 0) {
    numPaperFound += +(input[yPos][xPos - 1] === "@");
  }
  // right
  if (xPos + 1 < rowLength) {
    numPaperFound += +(input[yPos][xPos + 1] === "@");
  }
  // bottom left
  if (xPos - 1 >= 0 && yPos + 1 < input.length) {
    numPaperFound += +(input[yPos + 1][xPos - 1] === "@");
  }
  // bottom mid
  if (yPos + 1 < input.length) {
    numPaperFound += +(input[yPos + 1][xPos] === "@");
  }
  // bottom right
  if (xPos + 1 < rowLength && yPos + 1 < input.length) {
    numPaperFound += +(input[yPos + 1][xPos + 1] === "@");
  }

  return numPaperFound < 4;
};

const testGrid = (gridToTest: string[]) => {
  let nextGrid = gridToTest;
  let validPositionsCount = 0;
  for (let yPos = 0; yPos < gridToTest.length; yPos++) {
    const line = gridToTest[yPos];
    for (let xPos = 0; xPos < line.length; xPos++) {
      const canRemove = canRemovePaper(xPos, yPos, line.length);
      if (canRemove) {
        validPositionsCount += 1;
        const newChars = nextGrid[yPos].split("");
        newChars[xPos] = ".";
        nextGrid[yPos] = newChars.join("");
      }
    }
  }
  return { nextGrid, validPositionsCount };
};

let gridToTest = input;
let validPositions = 0;
let looping = true;
while (looping) {
  const { nextGrid, validPositionsCount } = testGrid(gridToTest);
  if (validPositionsCount === 0) {
    looping = false;
  } else {
    gridToTest = nextGrid;
    validPositions += validPositionsCount;
  }
}

console.log(validPositions);
