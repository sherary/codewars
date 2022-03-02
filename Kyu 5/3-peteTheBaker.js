/**
 * 
 * Pete likes to bake some cakes. He has some recipes and ingredients. 
 * Unfortunately he is not good in maths. 
 * Can you help him to find out, how many cakes he could bake considering his recipes?
 * 
 * Write a function cakes(), which takes the recipe (object) 
 * and the available ingredients (also an object) 
 * and returns the maximum number of cakes Pete can bake (integer). 
 * For simplicity there are no units for the amounts (e.g. 1 lb of flour or 200 g of sugar are simply 1 or 200). 
 * Ingredients that are not present in the objects, can be considered as 0.
 * 
 * Examples:
 * must return 2
 * cakes({flour: 500, sugar: 200, eggs: 1}, {flour: 1200, sugar: 1200, eggs: 5, milk: 200}); 
 * must return 0
 * cakes({apples: 3, flour: 300, sugar: 150, milk: 100, oil: 100}, {sugar: 500, flour: 2000, milk: 2000}); 
 */

// function cakes(recipe, available) {
//     // TODO: insert code
//     let num = []
//     if (Object.keys(recipe).length > Object.keys(available).length) return 0
//     available = Object.keys(available).sort().filter(key => Object.keys(recipe).includes(key)).reduce((obj, key) => {
//         obj[key] = available[key]
//         return obj
//     }, {})
//     recipe = Object.keys(recipe).sort().filter(key => Object.keys(recipe).includes(key)).reduce((obj, key) => {
//         obj[key] = recipe[key]
//         return obj
//     }, {})
//     let availableMats = Object.values(available)
//     let recipeMats = Object.values(recipe)

//     for (let i = 0; i < availableMats.length; i++) {
//         num.push(Math.floor(availableMats[i] / recipeMats[i]))
//     }

//     return num.sort((a, b) => a - b)[0]
// }

//alternative
function cakes(recipe, available) {
    return Object.keys(recipe).reduce(function(val, ingredient) {
      return Math.min(Math.floor(available[ingredient] / recipe[ingredient] || 0), val)
    }, Infinity)  
}
console.log(cakes({apples: 3, flour: 300, sugar: 150, milk: 100, oil: 100}, {sugar: 500, flour: 2000, milk: 2000})) // 0
console.log(cakes({flour: 500, sugar: 200, eggs: 1}, {flour: 1200, sugar: 1200, eggs: 5, milk: 200})) // 2
console.log(cakes({cream: 49, sugar: 59, flour: 87}, {pears: 7000, nuts: 6000, eggs: 1400, cream: 6400, butter: 4600, crumbles: 5600, milk: 5700, oil: 7500, apples: 7200, sugar: 5100, chocolate: 9900, flour: 8000, cocoa:1500 })) // 86
console.log(cakes({crumbles: 96, cocoa: 22, sugar: 74}, { oil: 6600, butter: 6300, cocoa: 4200, milk: 5100, cream: 9500, crumbles: 8400, apples: 8900, chocolate: 4400, flour: 9400, eggs: 8500, pears: 200, nuts: 4900, sugar: 7900 })) // 87