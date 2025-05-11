// prettier.config.js
module.exports = {
    // Add trailing commas wherever possible without introducing any syntax changes.
    trailingComma: 'es5', // Possible values: "none", "es5", "all"

    // Print spaces between brackets in object literals.
    bracketSpacing: true, // Possible values: true, false

    // Only include parentheses when necessary for arrow functions.
    arrowParens: 'avoid', // Possible values: "always", "avoid"

    // End lines with the system's line endings.
    endOfLine: 'auto', // Possible values: "auto", "lf", "crlf", "cr"

    // Wrap prose at word boundaries if it exceeds the print width.
    proseWrap: 'preserve', // Possible values: "preserve", "always", "never"

    // Require or disallow whitespace between HTML and XML tags.
    htmlWhitespaceSensitivity: 'css', // Possible values: "css", "strict", "true"

    // Use single quotes instead of double quotes.
    singleQuote: true, // Possible values: true, false

    // Set the maximum line length.
    printWidth: 100, // Any positive integer

    // Set the number of spaces per indentation level.
    tabWidth: 2, // Any positive integer

    // Add a semicolon at the end of statements.
    semi: true, // Possible values: true, false

    // Indent script and style tags in Vue files by one line.
    vueIndentScriptAndStyle: true, // Possible values: true, false
};

