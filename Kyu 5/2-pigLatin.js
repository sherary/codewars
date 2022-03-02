/**
 * 
 * Move the first letter of each word to the end of it, 
 * then add "ay" to the end of the word. 
 * Leave punctuation marks untouched.
 * 
 * Examples
 * pigIt('Pig latin is cool'); // igPay atinlay siay oolcay
 * pigIt('Hello world !');     // elloHay orldway !
 */

function pigIt(str) {
    //Code here
    return str.replace(/\s+/g, " ").split(' ').map(x => 
        x.match(/[.,\/#!$%\^&\*;?:{}=\-_`~()]/g) ? x : x.substring(1).concat(x[0], 'ay')
    ).join(' ')
}

//alternative
// function pigIt(str){
//     return str.replace(/(\w)(\w*)(\s|$)/g, "\$2\$1ay\$3")
// }
console.log(pigIt('Pig latin is cool'))//'igPay atinlay siay oolcay'
console.log(pigIt('This is my string'))//'hisTay siay ymay tringsay'
console.log(pigIt('Pig latin is cool'))// igPay atinlay siay oolcay
console.log(pigIt('Hello world  !')) // elloHay orldway !