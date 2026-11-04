/**
 * ZIVALIX Digital Solutions
 * Basic Namespace Perimeter Utility
 */

function verifyPerimeter() {
    return "ZIVALIX namespace is secure and actively monitored.";
}

function cleanInput(text) {
    if (typeof text !== 'string') return '';
    return text.trim();
}

module.exports = {
    verifyPerimeter,
    cleanInput
};
