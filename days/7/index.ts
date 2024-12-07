const fs = require("fs");

const part2 = true;

function parseInput(path: string): string[] {
    const fileContents = fs.readFileSync(path, "utf8");
    const result = fileContents.split("\n");

    return result;
}

function main() {
    const inputPath = "./my_input_file.txt";
    const lines = parseInput(inputPath);
    let sum = 0

    lines.forEach((line) => {
        const [result, equation] = parseLine(line);

        sum += checkIfEquationCanBeValid(result, equation)
    });

    console.log(sum)
}

function parseLine(line: string): [number, number[]] {
    var [result, equation] = line.split(":");
    equation = equation.trim();

    const equationArray = equation.split(" ").map((value) => parseInt(value));

    return [parseInt(result), equationArray];
}

function checkIfEquationCanBeValid(
    result: number,
    equation: number[]
): number {
    const operators = ["+", "*"];

    if (part2) operators.push("|")

    const operatorsAmount = equation.length - 1;
    const possibleCombinations = combinations(operators, operatorsAmount);

    for (const combination of possibleCombinations) {
        const equationArray = new Array();

        for (var i = 0; i < equation.length; i++) {
            equationArray.push(equation[i]);
            equationArray.push(combination[i]);
        }

        if (calculateEquation(equationArray) === result) {
            return result;
        }
    }

    return 0;
}

// code from a random stackoverflow page
const combinations = (arr: string[], length: number) => {
    const combination = (arr: string[], length: number): string[] => {
        if (length === 1) return arr;
        const result = combination(arr, length - 1).flatMap((val: string) =>
            arr.map((char) => val + char)
        );
        return result;
    };

    return combination(arr, length);
};

function calculateEquation(equation: (string | number)[]): number {
    let result: number = typeof  equation[0] == "number" ? equation[0] : 0

    equation.forEach((value, i) => {
        if (typeof value == "string") {
            switch (value) {
                case "+":
                    result += equation[i + 1] as number;
                    break;
                case "*":
                    result *= equation[i + 1] as number;
                    break;
                case "|":
                    result = parseInt(`${result}${equation[i+1]}`)
            }
        }
    })
    return result
}

main();
