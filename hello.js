// Simple Hello World program in JavaScript
console.log("Hello, World!");
console.log("This is a simple JavaScript program.");

// Simple calculation
const result = 2 + 2;
console.log(`2 + 2 = ${result}`);

// Simple array operation
const numbers = [1, 2, 3, 4, 5];
const sum = numbers.reduce((acc, num) => acc + num, 0);
console.log(`Sum of [${numbers.join(', ')}] = ${sum}`);

// Simple function
function greet(name) {
    return `Hello, ${name}!`;
}

console.log(greet("World"));