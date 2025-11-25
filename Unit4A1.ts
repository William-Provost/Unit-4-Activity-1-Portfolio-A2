/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-24
 * @fileoverview This program prints all the even numbers between two odd values.
 */

// variables
let startValue: number = 0;
let endValue: number = 0;

// get starting value
const startInput = prompt("Please enter a starting value (it must be an odd number): ");
startValue = parseInt(startInput || "0");

// get ending value
const endInput = prompt("Please enter an ending value (it must be an odd number): ");
endValue = parseInt(endInput || "0");

// print even numbers
let output = "";

// loop through values
for (let counter = startValue + 1; counter < endValue; counter++) {
  if (counter % 2 === 0) {
    // add number and a space (plain text, not template literal)
    output = output + counter + " ";
  }
}

// print final result (plain text)
console.log(output);

console.log("Done.");
