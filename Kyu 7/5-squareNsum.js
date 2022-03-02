/**
 * Complete the square sum function so that it squares each number passed into it
 * and then sums the results together.
 * 
 * For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9.
*/

function squareSum(numbers) {
    // attempt 1:
    // for (let i = 0; i < numbers.length; i++) {
    //     numbers[i] = numbers[i] ** 2
    // }
    // return numbers.reduce((x, y) => x + y, 0)

    // attempt 2:
    // return numbers.reduce(function(sum, n) {
    //     return (n * n) + sum
    // }, 0)

    // attempt 3:
    return numbers.reduce((sum, num) => sum + (num * num), 0)
}

console.log(squareSum([1,2]))//5 
console.log(squareSum([0, 3, 4, 5])) //50
console.log(squareSum([]))//0